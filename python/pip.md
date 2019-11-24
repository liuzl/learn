> Cannot uninstall 'llvmlite'. It is a distutils installed project and thus we cannot accurately determine which files belong to it which would lead to only a partial uninstall.

```sh
pip install librosa --ignore-installed llvmlite
```
