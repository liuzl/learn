# 步骤

1. 构建镜像

```sh
sudo docker build -t registry.cn-shenzhen.aliyuncs.com/zliu/tianchi:1.0 /home/zliu/tianchi
```

2. 执行

```sh
sudo nvidia-docker run -v /data:/tcdata registry.cn-shenzhen.aliyuncs.com/zliu/tianchi:1.0 sh run.sh
```

3. 将镜像推送到Registry

```sh
sudo docker login --username=qq94048699 registry.cn-shenzhen.aliyuncs.com
sudo docker push registry.cn-shenzhen.aliyuncs.com/zliu/tianchi:1.0
```
