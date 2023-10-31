# querystring

## 背景介绍
最近我们面临一个需求：需要提供一个表达式查询功能，例如：

* 1、TI:(小米)
* 2、TI:(小米 OR 华为)
* 3、mobile:((小米 OR 华为) NOT (小米*Pro OR 华为*Pro)) AND type:("手机" OR "电脑") NOT RAM:(8)

我们需要将这些表达式解析转换成Elasticsearch的查询语句。
尽管Elasticsearch本身支持查询字符串，但无法满足我们的需求。
最初，我们尝试使用正则表达式进行提取，但遇到了各种问题。
后来，我们尝试将表达式替换为Go语言AST能够识别的语法，但这种方法不够方便修改和扩展。
最终，我们发现了一个优秀的开源库 [bytedance/go-querystring-parser](https://github.com/bytedance/go-querystring-parser)
并基于该库进行了定制修改，以满足我们的需求。

## 使用方法

首先，你需要引入我们定制的查询字符串处理库：

```go
    import (
        "github.com/MichealJl/querystring"
    )
```

```go
    expr := `TI:(小米 OR 华为)`
    q, err := querystring.Parse(expr)
    if err != nil {
        // 处理错误
    }
    
    // 业务处理
```

## 自定义开发

如果你需要定制满足自己需求的查询语法，你可以修改 querystring.y 文件，定义适合你的语法规则。然后，运行以下命令来生成相应的Go代码：
```shell
goyacc -o querystring.y.go querystring.y
```
    
    