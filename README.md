shazam
======
A web client to interact with shazam.com APIs written in Go.

You can find the docs at https://godoc.org/github.com/bst27/shazam.

# Example
How to fetch the Shazam charts for Berlin, Germany:
```
package main

import (
	"github.com/bst27/shazam"
)

func main() {
	charts := shazam.FetchCityCharts(shazam.Cities().GermanyBerlin())
}
```