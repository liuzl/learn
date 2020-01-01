# 设计文档

## 数据结构说明

数据结构主要在源文件[types.go](https://github.com/liuzl/fmr/blob/master/types.go)中定义。

### FMR结构

```go
// FMR stands for Funtional Meaning Representation
type FMR struct {
	Fn   string `json:"fn,omitempty"`
	Args []*Arg `json:"args,omitempty"`
}
```

其中有两个字段：`Fn`表示函数名称，`Args`表示函数的参数。

### FMR函数参数

```go
// Arg is the type of argument for functions
type Arg struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
```

其中`Type`有如下几种类型：
* `index`：索引类型，`$1`表示第一个子节点的FMR，`$2`表示第二个子节点的FMR，以此类推。其中`$0`表示被此节点匹配到的自然语言字符串；
* `context`：上下文类型，目前用`$@`表示，后续计划更改为`@`；
* `string`：字符串类型
* `float`：浮点数类型
* `int`：整数类型
* `func`：FMR类型
