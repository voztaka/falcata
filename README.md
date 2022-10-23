# falcata
## Overview
When you copy and paste strings from other sites, it would have duplicated space or new line. Falcata will cut this innecessary fields.

## Installation

```bash
go get github.com/voztaka/falcata
```

## Sample Code

```go
package main

import(
	"fmt"
	"github.com/voztaka/falcata"
)

var copy = `The copy string
  has
new lines`

func main() {
	falcata.CutFields(&c)
	
	fmt.Println(c) // The copy string has new lines
}
```
