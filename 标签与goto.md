for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（:）结尾的单词（gofmt 会将后续代码自动移至下一行）。
示例 5.13 for6.go：
（标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母）
```go
package main

import "fmt"

func main() {

LABEL1:
    for i := 0; i <= 5; i++ {
        for j := 0; j <= 5; j++ {
            if j == 4 {
                continue LABEL1
            }
            fmt.Printf("i is: %d, and j is: %d\n", i, j)
        }
    }

}
```
continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置
如果将 continue 改为 break，则不会只退出内层循环，而是直接退出外层循环了

另外，还可以使用 goto 语句和标签配合使用来模拟循环。
示例 5.14 goto.go：
```go
package main

func main() {
    i:=0
    HERE:
        print(i)
        i++
        if i==5 {
            return
        }
        goto HERE
}
```
上面的代码会输出 01234

如果必须使用 goto，应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败