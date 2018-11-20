# 运维系统相关

* http://book.open-falcon.org/zh_0_2/

## 流量旁路
* [基于生产环境的仿真流量测试](https://www.lengyuewusheng.com/2017/08/13/00017_%E5%9F%BA%E4%BA%8E%E7%94%9F%E4%BA%A7%E7%8E%AF%E5%A2%83%E7%9A%84%E4%BB%BF%E7%9C%9F%E6%B5%81%E9%87%8F%E6%B5%8B%E8%AF%95/)
* [nginx旁路](https://tenfy.cn/2017/09/22/nginx-capture-multi/)
* **https://github.com/buger/goreplay**
  * [设置镜像流量的百分比](https://github.com/buger/goreplay/wiki/Rate-limiting)
    ```sh
    # replay server will not get more than 20% of requests 
    # useful for high-load environments
    sudo ./goreplay --input-raw :9301 --output-http="http://localhost:8301|20%"
    ```
## iptables
* http://www.111cn.net/sys/linux/46342.htm
