# Case2

#### Old:

```go
package main

import "encoding/json"
import "fmt"
type DataTypeA struct {
	DataA int `json:"dataA"`
}
type DataTypeB struct {
	DataB string `json:"dataB"`
}

func main() {
	dataA := "{\"dataA\":1}"
	datab := "{\"dataB\":\"one\"}"
	var a DataTypeA
	_=json.Unmarshal([]byte(dataA), &a)
	var b DataTypeB
	_=json.Unmarshal([]byte(datab), &b)
	
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}
```

* New
```go
package main

import "encoding/json"
import "fmt"

type DataTypeA struct {
	DataA int `json:"dataA"`
}
type DataTypeB struct {
	DataB string `json:"dataB"`
}

func unmarshal[T any](data []byte) (t T, err error) {
	err = json.Unmarshal(data, &t)
	return
}

func main() {
	dataA := "{\"dataA\":1}"
	datab := "{\"dataB\":\"one\"}"

	a, _ := unmarshal[DataTypeA]([]byte(dataA))
	b, _ := unmarshal[DataTypeB]([]byte(datab))

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}

```

