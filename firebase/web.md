## Firebase部署网站
### 安装开发环境Firebase CLI
```sh
# on AWS ubuntu server
# 安装Node.js、npm和firebase
curl -sL https://deb.nodesource.com/setup_9.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo apt install npm
npm install -g firebase-tools
```
### 初始化网站
```sh
firebase login --no-localhost
firebase init
```
### 部署网站
```sh
firebase deploy
```

### 参考资料
* *[firebase托管官方文档](https://firebase.google.com/docs/hosting/?hl=zh-cn)*
* *[codelabs例子](https://codelabs.developers.google.com/codelabs/firebase-web/)*
