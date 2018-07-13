## 分支开发
* 创建develop分支
```sh
git checkout -b develop master
```

* 将develop分支发布到master分支
```sh
# 切换到Master分支
git checkout master

# 对Develop分支进行合并
git merge --no-ff develop
```

## 新建删除远程分支
* 查看本地分支
```sh
$ git branch
* master
  release

$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
nothing to commit, working directory clean
```
* 新建本地分支
```sh
$ git checkout -b develop

$ git branch
* develop
  master
  release
```
* 推送到远程分支
```sh
$ git push origin develop:develop
```
* 删除远程分支
```sh
$ git push origin --delete develop
```

## 参考资料
* [Git分支管理策略](http://www.ruanyifeng.com/blog/2012/07/git.html)
* [新建远程分支和删除](https://www.jianshu.com/p/ea1dab2de419)
