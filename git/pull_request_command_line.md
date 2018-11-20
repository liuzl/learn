**Step 1**: From your project repository, check out a new branch and test the changes.
```sh
git checkout -b crawlerclub-master master
git pull https://github.com/crawlerclub/et.git master
```
**Step 2**: Merge the changes and update on GitHub.
```sh
git checkout master
git merge --no-ff crawlerclub-master
git push origin master
```
