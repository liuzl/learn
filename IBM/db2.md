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

## Python连接db2

```python
import ibm_db
conn = ibm_db.connect("DATABASE=sample;HOSTNAME=localhost;PORT=50000;PROTOCOL=TCPIP;UID=db2inst1;PWD=db2inst1;", "", "")
sql = "SELECT * FROM STAFF"
stmt = ibm_db.exec_immediate(conn, sql)
result = ibm_db.fetch_both(stmt)
while result:
    print(result)
    result = ibm_db.fetch_both(stmt)
```

## db2常见命令
```sh
$ #连接数据库
$ db2 connect to sample
$ #列出数据表
$ db2 list tables
$ #列出系统表
$ db2 list tables for system
$ #列出所有表
$ db2 list tables for all 
$ #列出用户表
$ db2 list tables for user
$ #显示表结构
$ db2 describe table tablename

```
