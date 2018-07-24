# 终端命令使用ss

## ShadowsocksX-NG

添加如下代码到`.bash_profile`
```sh
export http_proxy='http://localhost:1087'
export https_proxy='http://localhost:1087'

function proxy_off(){
    unset http_proxy
    unset https_proxy
    echo -e "已关闭代理"
}

function proxy_on() {
    export no_proxy="localhost,127.0.0.1,localaddress,.localdomain.com"
    export http_proxy="http://127.0.0.1:1087"
    export https_proxy=$http_proxy
    echo -e "已开启代理"
}
```

## 参考资料
* https://tech.jandou.com/to-accelerate-the-terminal.html
