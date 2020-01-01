# 设计文档

## 数据结构说明

数据结构主要在源文件[types.go](https://github.com/liuzl/fmr/blob/master/types.go)中定义。

### FMR函数参数

```go
// Arg is the type of argument for functions
type Arg struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
```

其中`Type`有如下几种类型：
* `index`
* `context`
* `string`
