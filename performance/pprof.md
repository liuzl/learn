## [Go的pprof使用](https://www.cnblogs.com/yjf512/archive/2012/12/27/2835331.html)

### 程序中引入`runtime/pprof`包
```go
package main

import (
  "flag"
  "runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }
}

```

### 运行程序生成profile文件
```sh
$ ./main -cpuprofile profile.prof
```

### 使用`go tool pprof`的`web`命令生成可用浏览器查看的svg文件
```
$ go tool pprof ./main profile.prof
```

### 使用浏览器查看svg文件
