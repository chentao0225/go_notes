## 1.初始包管理
关于包管理的总结：
    - 一个文件夹可以称为一个包。
    - 在文件夹(包)中可以创建多个文件。
    - 在同一个包下的每个为文件中必须制定 包名称且相同

重点：关于包的分类
    - main包，如果是main包，则必须写一个main函数，此函数就是项目的入口(main主函数)。编译生成的就是一个可执行的文件。
    - 非main包，用来将代码进行分类，分别放在不同的包和文件中。

## 2.输出
在终端将想要展示的数据显示出来，例如：欢迎登录、请输入用户名等。。。
- 内置函数
    - print
    - println
- fmt包(推荐)
    - fmt.Print
    - fmt.Println
    - fmt.Printf
        - %s，占位符 "文本"
        - %d，占位符 整数
        - %f，占位符 小数(浮点数)
        - %.2f 占位符 小数(保留小数点后2位，四舍五入)
扩展：进程里有stdin/stdout/stderr。

## 3.注释
- 单行注释，//
- 多行注释，/* */

## 4.初始数据类型
- 整型，整数
- 字符串，文本
- 布尔型，真假

## 5.变量
可以理解为昵称
    - 声明+赋值
    ```go
        var str string = "soleil"
    ```
    - 先声明后赋值
    ```go
        var str string
        str = "soleil"
    ```
### 5.1声明变量的意义
  - 编写代码省事
  - 存储结果，方便之后使用
  - 存储用户输入的值
    ```go
        var name string
        fmt.Scanf("%s", &name)
        if name == "soleil" {
            fmt.Println("用户名输入正确")
        } else {
            fmt.Println("用户名输入错误")
        }
    ```
### 5.2变量名要求
  - 变量名必须只包含：字母、数字、下划线
  - 数字不能开头
  - 不能使用go语言内置的关键字
  - 建议
    - 变量名见名知意
    - 驼峰式命名

### 5.3变量简写
  - 声明 + 赋值
    ```go
        var name string = "soleil"
        var name = "soleil"
        name := "soleil"
    ```
  - 先声明再复制
    ```go
        var name string
        var age int
        
        var a,b,c string
        name = "soleil"
        age = 22
        a = "A"
        b = "B"
        c = "C"

    ```

因式分解，例如：声明5个变量，分别有字符串、整形

```go
    var (
        name = "soleil"
        age = 22
        gender string
        length int
        yz bool
    )
```

### 5.4作用域
如果我们定义了大括号，那么在大括号中定义的变量。
 - 不能被上级使用。
 - 可以在同级使用。
 - 可以在子级使用。

全局变量和局部变量
 - 全局变量，未写在函数中的变量称为全局变量；不可以使用v1:xx方式简化；可以基于因式分解方式声明多个变量；项目中寻找变量时最后一环。
 - 局部变量，编写在{}里面的变量；可以使用任意方式简化；可以基于因式分解方式声明多个变量；

### 5.5赋值及内存相关
**注意事项**
> 使用int、string、bool这三种数据类型时，如果遇到变量的赋值则会拷贝一份。【值类型】

## 6.常量
不可被修改的变量。

### 6.1因式分解
```go
    const (
        v1 = 1
        v2 = 2
        v3 = 3
    )
```
### 6.2全局
```go
    package main
    const Data = 999
    const (
        pi = 3.1415926
        gender = "男"
    )
    func main(){
        //...
    }
```
### 6.3 iota
可有可无，当做一个在声明常量时的一个计数器
```go
    const (
        v1 = iota + 1
        v2 
        v3
        _
        v5
    )
```

## 7.输入
让用户输入数据，完成项目交互。
- fmt.Scan
    > 当使用Scan时，会提示用户输入  
    > 用户输入完成之后，会得到两个值：count，用户输入了几个值；err，用户输入错误则是错误信息，否则是 nil  
    > fmt.Scan 要求输入两个值，必须输入两个，否则他会一直等待
- fmt.Scanln
    > fmt.Scanln 等待回车
- fmt.Scanf

```go
    var name string
    fmt.Scanf("%s",&name)
    fmt.Println(name)
```

## 8.条件语句

### 8.1最基本
```go
    if 条件 {
        成立后，此代码块执行
    }else{
        不成立，此代码块执行
    }
```
### 8.2多条件判断
```go
    if 条件A{
        ...
    }else if 条件B{
        ...
    }else{
        ...
    }
```