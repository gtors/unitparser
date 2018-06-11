# Unitparser
Parser of Physical Quantities and SI Units (Russian and English variants) implemented in Golang 

# Notes
* Float numbers not supported yet
* Currently supported follow measurement units: Ampere, Volt, Ohm, Farad, Hertz and Henry.

# Example

```go
package main

import "github.com/gtors/unitparser"

func main() {

    // Russian
    if q, ok := unitparser.ParseString("1 мкФ"); ok {
      println(q.Value)
      // 1
      println(q.Prefix.String())
      // μ
      println(q.Kind.String())
      // F
      println(q.String()) 
      // 1 μF
    }
    
    // English
    if q, ok := unitparser.ParseString("1 μF"); ok {
      // ... Same as above
    }
}
```
