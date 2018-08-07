## tips

* 按条件筛选：`df[df['value']>60000]`
* `read_csv`的时候空值保持为空字符串，使用参数`keep_default_na=False`
  ```python
  df = pd.read_csv(file, keep_default_na=False)
  ```


## References
* https://pandas.pydata.org/pandas-docs/stable/10min.html
