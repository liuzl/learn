# IBM for banking

## DB2

* https://hub.docker.com/r/ibmcom/db2express-c/
* https://blog.csdn.net/qq_39154147/article/details/79219414

* python package ibm_db
  * https://github.com/ibmdb/python-ibmdb/issues/187

* 运行DB2
  * https://dongrenwen.github.io/2018/05/07/docker-install-db2/
  * docker run -it --network=host -p50000:50000 -e DB2INST1_PASSWORD=db2inst1-pwd -e LICENSE=accept ibmcom/db2express-c:latest bash
