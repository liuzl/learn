# wordpress installation on ubuntu

### install nginx & php
```sh
sudo apt install nginx
sudo apt install php-fpm
sudo apt install php-mysql
```
### install docker & mysql
```sh
sudo apt install docker.io
sudo docker pull mysql
# using dir on host system
sudo docker run --name first -p 3306:3306 -v /hostdir/dbfiles:/var/lib/mysql -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql
```
