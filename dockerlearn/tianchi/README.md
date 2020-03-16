# 步骤

1. 构建镜像

```sh
sudo docker build -t registry.cn-shenzhen.aliyuncs.com/zliu/tianchi:1.0 /home/zliu/tianchi
```

2. 执行

```sh
sudo nvidia-docker run -v /data:/tcdata registry.cn-shenzhen.aliyuncs.com/zliu/tianchi:1.0 sh run.sh
```
