# pilosa

* https://www.pilosa.com/docs/latest/getting-started/
* [白皮书](https://www.pilosa.com/pdf/PILOSA%20-%20Technical%20White%20Paper.pdf)

## 界面
* https://github.com/pilosa/console

## MISC
### 启动服务
```sh
pilosa server --handler.allowed-origins="http://localhost:8000"
```
### 启动console
```sh
git clone https://github.com/pilosa/console
cd console/assets
python -m http.server
```
