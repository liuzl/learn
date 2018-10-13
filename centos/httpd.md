# Apache + PHP on CentOS

## 安装程序
```sh
sudo yum install httpd.x86_64
sudo yum install php php-devel
# 安装php的mongo扩展
sudo yum install php-pecl-mongo.x86_64
```
## 启动httpd
```sh
sudo service httpd start
```
## 创建测试脚本
```sh
cat "<?php\nphpinfo();" > /var/www/html/info.php
```
## 打开浏览器
http://your_ip/info.php
