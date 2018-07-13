## 分支开发
* Git创建Develop分支的命令：
```sh
git checkout -b develop master
```

* 将Develop分支发布到Master分支的命令：
```sh
# 切换到Master分支
git checkout master

# 对Develop分支进行合并
git merge --no-ff develop
```

## 参考资料
* [Git分支管理策略](http://www.ruanyifeng.com/blog/2012/07/git.html)
