1.前缀和后缀
HasPrefix 判断字符串 s 是否以 prefix 开头：
```go
strings.HasPrefix(s, prefix string) bool
```

HasSuffix 判断字符串 s 是否以 suffix 结尾：
```go
strings.HasSuffix(s, suffix string) bool
```

例如
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string = "This is an example of a string"
    fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
    fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
```
输出：
T/F? Does the string "This is an example of a string" have prefix Th? true

2.字符串包含关系
Contains 判断字符串 s 是否包含 substr：
```go
strings.Contains(s, substr string) bool
```

3.判断子字符串或字符在父字符串中出现的位置（索引）
Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
```go
strings.Index(s, str string) int
```

LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
```go
strings.LastIndex(s, str string) int
```
如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位：
```go
strings.IndexRune(s string, ch int) int
```
示例 4.14 index_in_string.go
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string = "Hi, I'm Marc, Hi."

    fmt.Printf("The position of \"Marc\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Marc"))

    fmt.Printf("The position of the first instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Hi"))
    fmt.Printf("The position of the last instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

    fmt.Printf("The position of \"Burger\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Burger"))
}
```
输出：
The position of "Marc" is: 8
The position of the first instance of "Hi" is: 0
The position of the last instance of "Hi" is: 14
The position of "Burger" is: -1
