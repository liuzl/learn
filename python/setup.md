# pypi

## 注册用户zliu
* https://packaging.python.org/tutorials/packaging-projects/
* zliu/liang@zliu.org

## setup.py编写

### An example

#### Create the following file structure

```
package_tutorial/
  example_pkg/
    __init__.py
```

`example_pkg/__init__.py` is required to import the directory as a package, and can simply be an empty file.

#### Create the package files

```
package_tutorial/
  example_pkg/
    __init__.py
  setup.py
  LICENSE
  README.md
```

#### Create setup.py

```python
import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="example-pkg-YOUR-USERNAME-HERE", # Replace with your own username
    version="0.0.1",
    author="Example Author",
    author_email="author@example.com",
    description="A small example package",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/pypa/sampleproject",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
)
```

#### Generating distribution archives

```sh
python3 -m pip install --user --upgrade setuptools wheel
python3 setup.py sdist bdist_wheel
```

#### Uploading the distribution archives

```sh
python3 -m pip install --user --upgrade twine
python3 -m twine upload --repository-url https://test.pypi.org/legacy/ dist/*
```

## 参考资料
* [python 编写简单的setup.py](https://www.cnblogs.com/lyrichu/p/6818008.html)
* [How to avoid building C library with my python package?](https://stackoverflow.com/questions/31380578/how-to-avoid-building-c-library-with-my-python-package)
* [setup.py里的几个require](https://note.qidong.name/2018/01/python-setup-requires/)
