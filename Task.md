1.利用闭包实现一个计数器
```go
package main

import "fmt"

func count() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c := count()
	for {//无限循环，可设置计数多少
		fmt.Println(c())
	}
}
```

2.map和slice扩容的区别
(i)slice
*slice基于数组，拥有长度和容量
*添加元素时，长度小于容量会直接将元素添加进slice，等于则会分配更大的底层数组再将原元素复制到新数组再添加新元素
*slice的扩容形式为n+n
*slice的时间复杂度为O(n),平均时间复杂度为O(1)
(ii)map
*map基于哈希表
*元素数量比桶的数量超过某个阈值会进行扩容
*map扩容形式为重新计算桶的数量重新哈希再分配
*时间复杂度较高，平均时间复杂度为O(1)

3.比较字符串"hello，世界"的长度 和for range 该字符串的循环次数
```go
package main

import "fmt"

func main() {
	s := "hello，世界"
	a := len(s)
	fmt.Printf("字符串长度为%d\n", a)

	count := 0
	for range s {
		count++
	}
	fmt.Printf("for range循环次数%d", count)
}
```
输出结果：
```go
字符串长度为14
for range循环次数8
```
原因：
*len(s)返回字符串字节数，ASCⅡ字符"hello"占5字节，而","和"世界"非2字节
*for range遍历"世""界"时分别视为一个字符

4.如何进行结构体之间的比较
*如果结构体所有字段支持相等比较，可以使用==直接比较
比如机试类似题
```go
package main

import (
	"fmt"
)

type Person struct {
	xingge int
	waimao int
}

func main() {
	truesjn := Person{xingge: 100, waimao: 100}
	sjn1 := Person{xingge: 100, waimao: 100}
	sjn2 := Person{xingge: 100, waimao: 90}

	fmt.Println("sjn1为真", truesjn == sjn1)
	fmt.Println("sjn2为假", truesjn == sjn2)
}
```
*如果含有不可比较字段（切片），可以自定义函数进行比较，也可以指定某一部分比较
```go
package main

import (
	"fmt"
)

type Person struct {
	xingge int
	waimao int
}

func main() {
	truesjn := Person{xingge: 100, waimao: 100}
	sjn1 := Person{xingge: 100, waimao: 100}
	sjn2 := Person{xingge: 100, waimao: 90}

	fmt.Println("sjn1为真", truesjn.waimao == sjn1.waimao)
	fmt.Println("sjn2为假", truesjn.waimao == sjn2.waimao)
}
```
指定比较外貌

5.以下哪里x进行重新声明，为什么
```go
func main() {
    x := "hello!"
    for i := 0; i < len(x); i++ {
        x := x[i]//place 1
        if x != '!' {
            x := x + 'A' - 'a'//place 2
            fmt.Printf("%c", x) 
        }
    }
}
```
x在上述两个位置重新定义
main下面的定义x是一个局部变量而非全局变量
在for循环中不可被识别
故for循环中是重新定义一个新的变量，这个变量与上面的局部变量不是同一个变量
place 1保存识别的每一个字母
place 2将其变为大写形式