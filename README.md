# interpretor
A simple interpretor using goolgle translate.

# example use

```golang
package main

import (
    "fmt"
    "github.com/fyh/interpretor"
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

# install 
```
go get github.com/fyh/interpretor
```

# documentation

see [http://godoc.org/github.com/fyh/interpretor](http://godoc.org/github.com/fyh/interpretor)
