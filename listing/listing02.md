Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```

Отложенные функции могут считывать и присваивать именованные возвращаемые значения возвращаемой функции. В test() x является именованным возвращаемым значением, поэтому при выходе из программы defer инкрементирует её и получается 2. 

В anotherTest() возвращаемое значение не является именованным.

Defer`ы работают по принципу (LIFO) «последним пришел — первым обслужен»