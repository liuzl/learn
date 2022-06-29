# 学习资料
网络上收集整理的学习资料

## 服务高可用技术
1. [你所不知道的TIME_WAIT和CLOSE_WAIT](https://mp.weixin.qq.com/s?__biz=MzI4MjA4ODU0Ng==&mid=402415747&idx=1&sn=2458ba4fe1830eecdb8db725d3f395fa&scene=1&srcid=0219ONAdie0Wa6o3pW47CIln&key=710a5d99946419d9d2aa813f0ff66c0cc084049b289f17d551d542f09c5f327f17617efb5d5c308bfac52a8b4ca612a8&ascene=0&uin=Mjk1ODMyNTYyMg%3D%3D&devicetype=iMac+MacBookPro11%2C4+OSX+OSX+10.11.3+build(15D21)&version=11020201&pass_ticket=vFd4jXC%2F94fd4APMhl%2FH3VGBe0rwoZxqkT0m3VscvbbfVlhihe4EEqUYSH1z1Bbv)

> 要点: 讲解TIME_WAIT的文章百度一下可以发现一大堆, 但是往往缺少讲述比较详细准确的文章, 本文从协议层讲述了TIME_WAIT的原理, 并且给出了实际场景中的案例.

* [网络优化之net.ipv4.tcp_tw_recycle和tcp_tw_reuse参数](https://www.qingtingip.com/h_256514.html)

```sh
echo "1" > /proc/sys/net/ipv4/tcp_tw_reuse
echo "1" > /proc/sys/net/ipv4/tcp_tw_recycle
```

## 分布式系统实践
1. [Elasticsearch 架构以及源码概览](https://mp.weixin.qq.com/s?__biz=MzA4NjgwMDQ0OA==&mid=2652445487&idx=1&sn=fb99fac1db2ad8120e98f00165323b2a&scene=1&srcid=0711ZEP1y8SLXzpImelywT5N&key=77421cf58af4a653fadcc7594168daa7b232d787e310fbe4536f0919298fe4284caa929b0d7a9463712f20c95a3e4b9a&ascene=0&uin=Mjk1ODMyNTYyMg%3D%3D&devicetype=iMac+MacBookPro11%2C4+OSX+OSX+10.11.5+build(15F34)&version=11020201&pass_ticket=zxcWol980kEHZxuscCwv6Bo4lYDE30dPbCHaz0sBBISlEX1GvXJw875tzW6lzG8M)

> 重点介绍了ElasticSearch的分布式架构。

2. [驱动海量大数据实时多维分析，优酷为什么会选择Druid？](https://mp.weixin.qq.com/s?__biz=MzA5NzkxMzg1Nw==&mid=2653160326&idx=1&sn=9c6a91df0ff088f799eefe2ca14926ca&scene=0&key=8dcebf9e179c9f3a11295728e84286c8427ddea06d951c4e08f2efb52ad1d982d9f5f8bbe434244929ff240d0ec88b5d&ascene=0&uin=Mjk1ODMyNTYyMg%3D%3D&devicetype=iMac+MacBookPro11%2C4+OSX+OSX+10.11.5+build(15F34)&version=11020201&pass_ticket=Uax4it219TvGLrsu%2B5wo6lhIzCrLx6RPRmntpCJlCs%2FagzKHed%2B%2F8X23mVqpb6o0)

> 得益于搜索引擎的模糊查询能力, 现在采用搜索引擎作为metric数据存储的系统越来越多了。常用的就是elasticsearch，这篇文章介绍了一个新的系统Druid，作为elasticsearch的竞争对手，看看Druid为我们带来了什么。

## misc

* https://apple.stackexchange.com/questions/254380/macos-mojave-invalid-active-developer-path

## json to csv
* https://www.sanity.io/blog/exporting-your-structured-content-as-csv-using-jq-in-the-command-line

## gcc on windows
* https://blog.kedixa.top/2017/install-mingw-w64/

## -> in python
* https://stackoverflow.com/questions/14379753/what-does-mean-in-python-function-definitions

## ssh keep alive
```sh
vim ~/.ssh/config
ServerAliveInterval 60
```

* https://wpbeaches.com/how-to-make-your-ssh-terminal-shell-sessions-last-longer-on-macos/

## mac terminal start vscode
```sh
# add the following line to .bashrc
code () { VSCODE_CWD="$PWD" open -n -b "com.microsoft.VSCode" --args $* ;}
```

```sh
# start vscode under the current dir
code .
```

## mac, zsh after install nvm

> zsh compinit: insecure directories, run compaudit for list.

https://blog.csdn.net/CaptainJava/article/details/109585966

## jq

https://stackoverflow.com/questions/61492210/how-to-stringify-json-using-jq
