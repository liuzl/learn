## 一、下载nginx官方镜像和php-fpm镜像
```sh
sudo docker pull nginx
sudo docker pull bitnami/php-fpm
```
## 二、运行php-fpm
```sh
sudo docker run -d --name myphpfpm -p 9000:9000 -v /home/zliu/work/finance_kirin0.1/web/public:/usr/share/nginx/html bitnami/php-fpm
```
## 三、运行ngnix
```sh
sudo docker run -d --name myngnix -p 8080:80 -v /home/zliu/work/finance_kirin0.1/web/public:/usr/share/nginx/html nginx
```
## 四、修改ngnix配置
```sh
sudo docker exec -it myngnix /bin/bash
sudo docker cp ./ngnix.conf mynginx:/etc/nginx/conf.d/default.conf
```
