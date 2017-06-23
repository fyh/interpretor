# interpretor
A simple interpretor using goolgle translate.

# example use

```golang
package main

import (
    "fmt"
    "fyh/interpreter"
)

func main()  {
    i, _ := interpreter.Translate("hello, world", interpreter.ZH_CN)
    fmt.Println(i)

    i, _ = interpreter.Translate("你好，世界!", interpreter.EN)
    fmt.Println(i)
}
```

output is 
```
你好，世界
Hello world!
```
