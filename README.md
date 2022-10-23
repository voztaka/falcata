# falcata
## Overview
Copying and pasting strings from other sites would have duplicated space or new lines. Falcata will cut these unnecessary fields.

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

var copiedString = ` The copy string
  has
new lines.  `

func main() {
	falcata.CutFields(&copiedString)
	fmt.Println(copiedString)
}
```

## Run
```bash
go run main.go
The copied string has new lines.
```



