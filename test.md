# Demo MarkDown

## 标题 demo
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题

## 段落 demo

我是红红不气, 我要用markdown
我的QQ是:1026191605

## 强调 demo

我是*斜体*文字
我是**粗体**文字
我是***加粗斜体***文字
我是~~删除线~~文字

## 列表 demo

### 无序列表

- 一级列表  
    - 二级列表  
        - 三级列表
            - 四级列表
- QQ:1026191605
### 有序列表
1. Name: honghongbuqi
    1. 二级有序列表
    2. 
3. 插入列表
2. QQ:1026191605

## 链接 demo

### 内嵌式链接

- 外部链接[百度](http://www.baidu.com)
- 内部链接1: 链接仓库的其他文件: [README](README.md)
- 内部链接2: 链接本文档的其他部分: [代码块 demo](test.md#代码块-demo)

### 引用式链接

- 外部链接[百度]
- 内部链接1: 链接仓库的其他文件: [README]
- 内部链接2: 链接本文档的其他部分: [代码块 demo]

## 图片 demo

- 外部图片 
![baidu]( https://www.baidu.com/img/bd_logo1.png "百度图片") 

- 内部图片
![](coder.jpg)

## 引用 demo

- 普通引用

>这是一段引用.
>hello world!

出自《c program》

- 多重引用

>>> 这是多重引用

## 代码块 demo

- 行内代码

这段代码来声明变量 `var h int = 3`


- 块式代码
```go
 import (
            "github.com/bradfitz/gomemcache/memcache"
    )

    func main() {
         mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
         mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

         it, err := mc.Get("foo")
         ...
    }
```

[百度]: http://www.baidu.com
[README]: README.md
[代码块 demo]:test.md#代码块-demo










