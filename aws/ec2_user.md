## 允许密码登录
1. 修改`/etc/ssh/sshd_config`
* `PasswordAuthentication no`->`PasswordAuthentication yes`
  
2. 重启sshd
* `sudo /etc/init.d/ssh restart`

## 添加新用户
1. 添加用户
* `sudo adduser newuser`