# 制作ubuntu安装U盘
## 步骤
* 下载ubuntu server
```
wget http://mirrors.yun-idc.com/ubuntu-releases/18.04.1.0/ubuntu-18.04.1.0-live-server-amd64.iso
```

* 将ISO转为DMG
```
hdiutil convert -format UDRW -o ubuntu-18.04.1.0-live-server-amd64 ubuntu-18.04.1.0-live-server-amd64.iso
```


## 参考资料
* [Mac OS X ubuntu usb安装启动盘制作](https://www.jianshu.com/p/e2f41e217109)
* [ubuntu-18.04.1.0-live-server-amd64.iso](http://mirrors.yun-idc.com/ubuntu-releases/18.04.1.0/ubuntu-18.04.1.0-live-server-amd64.iso)
