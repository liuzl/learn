## tips

* 按条件筛选：`df[df['value']>60000]`
* read_csv的时候空值保持为空字符串，使用参数keep_default_na=False
  ```python
  df = pd.read_csv(file, keep_default_na=False)
  ```
  `keep_default_na`参数的作用是决定要不要保留默认应该转换的缺失值列表，将这个参数设为`False`之后同时不定义`na_values`参数，就可以在读取文件时不将任何值转换为缺失值`NaN`。

## References
* https://pandas.pydata.org/pandas-docs/stable/10min.html

## 通过正则表达式筛选数据
```python
df[df.Number.str.match('^628953604|^628953605|.*894.*')]
```

## 随机打乱数据顺序
```python
df = df.sample(frac=1).reset_index(drop=True)
```

## 读取超大csv文件
```python
for df in pd.read_csv("src.csv", header=None, chunksize=100000):
    df[1] = df[1].sample(frac=1).reset_index(drop=True)
    df.to_csv("dst.csv", header=None, index=False, mode='a')
```

## 直接读取s3文件
需要安装`s3fs`
```sh
conda install s3fs
```
```python
import pandas as pd
df = pd.read_csv("s3://izidata/gray.csv.gz", header=None)
```

## How to treat NULL as a normal string with pandas?
* https://stackoverflow.com/questions/44128033/pandas-reading-null-as-a-nan-float-instead-of-str
* https://stackoverflow.com/questions/50683765/how-to-treat-null-as-a-normal-string-with-pandas

```python
pd.read_csv(StringIO(data), na_filter=False)
```

## join数据
```python
result = pd.merge(click, df, on="key")
```

## 一列拆多列
```python
In [1]: from pandas import *

In [2]: def calculate(x):
   ...:     return x*2, x*3
   ...: 

In [3]: df = DataFrame({'a': [1,2,3], 'b': [2,3,4]})

In [4]: df
Out[4]: 
   a  b
0  1  2
1  2  3
2  3  4

In [5]: df["A1"], df["A2"] = zip(*df["a"].map(calculate))

In [6]: df
Out[6]: 
   a  b  A1  A2
0  1  2   2   3
1  2  3   4   6
2  3  4   6   9
```

## 去重
```python
df.drop_duplicates()
```

## 读取jsonl文件
```python
from pandas import json_normalize
import json

df = json_normalize([json.loads(x) for x in open('file.json')])
```
