# Docker

## Tutorials
* [Improving your data science workflow with Docker](https://unsupervisedpandas.com/data-science/docker-for-data-science/)
  * By [Rob Harrigan](https://unsupervisedpandas.com/), January 03, 2018

## Docker操作

### 进入docker shell
```sh
docker exec -it <container_id> bash
```

### https://www.fengbohello.top/archives/go-env-docker
```sh
docker run -it --rm -v /Users/hwang/golangdocker/go:/go --privileged golang bash
```

### 重启docker进程

```sh
systemctl restart docker.service
sudo service docker restart
```

### Run `docker` command without `sudo`

添加$USER到docker usergroup
```sh
sudo usermod -aG docker $USER
```
