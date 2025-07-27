// tag_bitmap_server.go

// 改进版本：
// - 使用 E.164 手机号格式（去除 +，转换为 uint64）作为 ID 存入 Roaring Bitmap（使用 sroar 支持 uint64）
// - BadgerDB 中以手机号字符串为 Key 存储用户信息
// - 标签查询支持多个标签交集

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dgraph-io/badger/v4"
	"github.com/gorilla/mux"
	"github.com/outcaste-io/sroar"
)

var (
	tagBitmaps = make(map[string]*sroar.Bitmap) // 标签名 → bitmap（uint64）
	tagPath    = "./bitmaps"                    // bitmap 文件目录
	db         *badger.DB
)

type User struct {
	Phone string   `json:"phone"`
	Name  string   `json:"name"`
	Tags  []string `json:"tags"`
}

func main() {
	loadBitmaps()
	initBadger()
	r := mux.NewRouter()
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/tag/add", addTagHandler).Methods("POST")
	r.HandleFunc("/query", queryHandler).Methods("GET")
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}

func initBadger() {
	opts := badger.DefaultOptions("./badger")
	db0, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	db = db0
}

func loadBitmaps() {
	os.MkdirAll(tagPath, 0755)
	files, _ := os.ReadDir(tagPath)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".bitmap") {
			tag := strings.TrimSuffix(f.Name(), ".bitmap")
			fp := filepath.Join(tagPath, f.Name())
			data, err := os.ReadFile(fp)
			if err == nil {
				bm := sroar.FromBuffer(data)
				tagBitmaps[tag] = bm
				log.Println("Loaded tag:", tag)
			}
		}
	}
}

func persistBitmap(tag string) error {
	fp := filepath.Join(tagPath, tag+".bitmap")
	data := tagBitmaps[tag].ToBuffer()
	return os.WriteFile(fp, data, 0644)
}

func e164ToUint64(phone string) (uint64, error) {
	clean := strings.TrimPrefix(phone, "+")
	return strconv.ParseUint(clean, 10, 64)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	id, err := e164ToUint64(u.Phone)
	if err != nil {
		http.Error(w, "invalid phone", 400)
		return
	}

	for _, tag := range u.Tags {
		bm, ok := tagBitmaps[tag]
		if !ok {
			bm = sroar.NewBitmap()
			tagBitmaps[tag] = bm
		}
		bm.Set(id)
		persistBitmap(tag)
	}

	val, _ := json.Marshal(u)
	db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(u.Phone), val)
	})
	w.WriteHeader(http.StatusCreated)
}

func addTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("name")
	if tag == "" {
		http.Error(w, "missing tag", 400)
		return
	}
	if _, ok := tagBitmaps[tag]; !ok {
		tagBitmaps[tag] = sroar.NewBitmap()
		persistBitmap(tag)
		fmt.Fprintln(w, "Tag created")
	} else {
		fmt.Fprintln(w, "Tag exists")
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	tags := r.URL.Query()["tag"]
	if len(tags) == 0 {
		http.Error(w, "no tag specified", 400)
		return
	}
	var res *sroar.Bitmap
	for i, tag := range tags {
		bm, ok := tagBitmaps[tag]
		if !ok {
			http.Error(w, "tag not found: "+tag, 404)
			return
		}
		if i == 0 {
			res = bm.Clone()
		} else {
			res.And(bm)
		}
	}

	result := []User{}
	for _, id := range res.ToArray() {
		phone := strconv.FormatUint(id, 10)
		db.View(func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(phone))
			if err == nil {
				val, _ := item.ValueCopy(nil)
				var u User
				json.Unmarshal(val, &u)
				result = append(result, u)
			}
			return nil
		})
	}
	json.NewEncoder(w).Encode(result)
}

