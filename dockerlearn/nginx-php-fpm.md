## 一、下载nginx官方镜像和php-fpm镜像
```sh
sudo docker pull nginx
sudo docker pull bitnami/php-fpm
```
## 二、运行php-fpm
```sh
sudo docker run -d --name myphpfpm -p 9000:9000 -v /home/zliu/work/finance_kirin0.1/web/public:/usr/share/nginx/html bitnami/php-fpm
```
