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

* `df[df.Number.str.match('^628953604|^628953605|.*894.*')]`
