变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
每一个源文件都可以包含且只包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
```go
package trans //定义包名

import "math" //引入math函数

var Pi float64 //定义Pi

func init() { //制作init包
   Pi = 4 * math.Atan(1)  //利用math函数中，因为arctan1等于Π/4，所以Π等于4×arctan1
}
```go

```go
package main

import (
   "fmt"
   "./trans" //引入Pi的init包
)

var twoPi = 2 * trans.Pi //为twoPi赋值

func main() {
   fmt.Printf("2*Pi = %g\n", twoPi)  //输出2Π的值
}