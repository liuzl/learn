## 获取本机IP

```sh
# 方法1
ip addr | awk '/^[0-9]+: / {}; /inet.*global/ {print gensub(/(.*)\/(.*)/, "\\1", "g", $2)}'

# 方法2
ip -o -4 addr show up primary scope global | awk '{print $4}' | awk -F'/' '{print $1}'
```
