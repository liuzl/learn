# 制作ubuntu安装U盘
## 步骤
* 下载ubuntu server
```sh
wget http://mirrors.yun-idc.com/ubuntu-releases/18.04.1.0/ubuntu-18.04.1.0-live-server-amd64.iso
```

* 将ISO转为DMG
```sh
hdiutil convert -format UDRW -o ubuntu-18.04.1.0-live-server-amd64 ubuntu-18.04.1.0-live-server-amd64.iso
```

* 卸载U盘设备
```sh
# 查看U盘设备号
# diskutil list
diskutil unmountDisk /dev/disk3
```

* 创建可启动的U盘
```sh
sudo dd if=ubuntu-18.04.1.0-live-server-amd64.dmg of=/dev/disk3 bs=1m
```

* 弹出U盘
```sh
diskutil eject /dev/disk3
```

## 参考资料
* [Mac OS X ubuntu usb安装启动盘制作](https://www.jianshu.com/p/e2f41e217109)
* [ubuntu-18.04.1.0-live-server-amd64.iso](http://mirrors.yun-idc.com/ubuntu-releases/18.04.1.0/ubuntu-18.04.1.0-live-server-amd64.iso)
