## django rest absolute urls always return 127.0.0.1

fixed by using the following Caddyfile: https://github.com/caddyserver/examples/blob/master/django/Caddyfile
* https://stackoverflow.com/questions/32166549/how-do-i-configure-the-django-rest-framework-pagination-url
* https://stackoverflow.com/questions/19669376/django-rest-framework-absolute-urls-with-nginx-always-return-127-0-0-1

## [Caddy server reload after update CaddyFile](https://caddy.community/t/caddy-server-reload-after-update-caddyfile/2645)

```sh
# method 1
pkill -USR1 caddy

## method 2
ps -C caddy
PID TTY TIME CMD
1392 pts/0 00:00:00 caddy
kill -s USR1 1392
```

## 参考资料

* https://liuzhichao.com/2018/caddy/
