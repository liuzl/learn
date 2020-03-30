## 获取本机IP

```sh
# 方法1
ip addr | awk '/^[0-9]+: / {}; /inet.*global/ {print gensub(/(.*)\/(.*)/, "\\1", "g", $2)}'

# 方法2
ip -o -4 addr show up primary scope global | awk '{print $4}' | awk -F'/' '{print $1}'
```

* [完美解决Linux crontab调用脚本 ifconfig、ip命令获取IP返回为空问题](https://b.abczn.com/archives/1370)，命令使用绝对路径，比如`/sbin/ip`
