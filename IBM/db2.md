## 使用docker安装db2环境

```sh
# 拉取db2的docker镜像
docker pull ibmcom/db2express-c

# 启动docker
docker run --name DB2ExpressC -d -p 50000:50000 -e DB2INST1_PASSWORD=db2inst1 -e LICENSE=accept  ibmcom/db2express-c:latest db2start

# 进入到启动的容器
docker exec -it DB2ExpressC /bin/bash

# 切换用户到 db2inst1
su - db2inst1

# 安装默认实例
db2sampl

# 连接到新创建的数据库实例
db2 connect to sample

# 执行SQL
db2 "SELECT * FROM STAFF"


```
