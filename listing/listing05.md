Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Ответ:
```
error
```

Интерфейс равен `nil` тогда, когда оба поля равны `nil`. `test()` возвращает интерфейс, в котором данные будут `nil`, но в первом поле её структуры тип будет определен *customError и не будет равен `nil`.