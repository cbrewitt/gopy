# gopy

Facilities easy calling of python from golang. Embeds the python code directly into the go app executable. Supports sending numpy arrays directly between go and python.

Often I use golang as the primary server langauge/app but have cases where I need to delegate to python for certain aspects like processing data and fitting ML models. In such cases, I may not want the overhead of running a separate python service and communicating between that and my go app. I can use this project instead to simply embedd my required python code directly into my go app.

For example, I can have the following structure:

```
pythonsrc/
   main.py
main.go
```

Where:

main.py
```python
# pip install gopyadapter
from gopyadapter.core import execute

def do_something():
    return {"a":1}


```


main.go
```go
package main

import "github.com/jptrs93/gopy/gopy"

//go:embed pythonsrc/*
var pythonSrc embed.FS

type DoSomethingResult struct {
	A int `msgpack:"a"`
}

func main() {
	gopy.InitDefaultPool(pythonSrc, "/path-to-python-env/python", "main.py", 1)
	
	
	res, err := gopy.CallDefault[DoSomethingResult]("do_something")
	
	// or res := gopy.MustCallDefault[DoSomethingResult]("do_something")
}
```
