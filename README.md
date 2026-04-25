# quadA

## Instructions

Write a function `QuadA` that prints a valid rectangle with a given width `x` and height `y`.

The function must draw the rectangles as shown in the examples below.

- If `x` and `y` are positive numbers, print the rectangle.
- Otherwise, print nothing.

Make sure you submit all the necessary files to run the program.

---

## Expected function

```go
func QuadA(x, y int) {

}

```

Expected 1

```go
package main

import "piscine"

func main() {

    piscine.QuadA(5, 3)
}

```

Output

```go

o---o
|   |
o---o

```

Expected 2

```go

package main

import "piscine"

func main() {
    piscine.QuadA(5, 1)
}

```

Output

```go

o---o

```

Expected 3

```go

package main

import "piscine"

func main() {
    piscine.QuadA(1, 5)
}

```

Output

```go

o
|
|
|
o
