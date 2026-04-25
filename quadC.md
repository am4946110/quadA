# QuadC Function Documentation

## Description

`QuadC` is a Go function that prints a rectangle of a specified width and height using ASCII characters. The rectangle is drawn with specific corner and edge characters:

- **'A'** - Top-left and top-right corners
- **'B'** - Edges (sides and edges of bottom row)
- **'C'** - Bottom-left and bottom-right corners
- **' '** (space) - Interior of the rectangle

## Function Signature

```go
func QuadC(x, y int)
```

### Parameters

- `x` (int): The width of the rectangle
- `y` (int): The height of the rectangle

## Behavior

- If `x` or `y` are **not positive** (≤ 0), the function prints **nothing**
- If `x` or `y` equal **1**, the function handles edge cases appropriately

## Examples

### Example 1: QuadC(5, 3)

```
ABBBA
B   B
CBBBC
```
- Width: 5, Height: 3
- Top row has 'A' corners with 'B's filling the middle
- Middle row has 'B' on sides with spaces in between
- Bottom row has 'C' corners with 'B's filling the middle

### Example 2: QuadC(5, 1)

```
ABBBA
```
- Width: 5, Height: 1
- Single row rectangle with 'A' at corners

### Example 3: QuadC(1, 1)

```
A
```
- Width: 1, Height: 1
- Single character 'A' representing a 1×1 rectangle

### Example 4: QuadC(1, 5)

```
A
B
B
B
C
```

- Width: 1, Height: 5
- First row: 'A' (top)
- Middle rows: 'B' (edges)
- Last row: 'C' (bottom)

## Implementation Details

1. **Input Validation**: Returns early if x or y are not positive
2. **Special Case Handling**: Handles width = 1 separately
3. **Row Construction**:
   - **First row**: 'A' + (x-2) × 'B' + 'A'
   - **Middle rows**: 'B' + (x-2) × ' ' + 'B'
   - **Last row**: 'C' + (x-2) × 'B' + 'C'

## Usage

### Program Test 1

```go
package main
import "piscine"

func main() {
  piscine.QuadC(5, 3)
}
```

### Program Test 2

```go
package main
import "piscine"

func main() {
  piscine.QuadC(5, 1)
}
```

### Program Test 3

```go

package main
import "piscine"

func main() {
  piscine.QuadC(1, 1)
}
```

### Program Test 4

```go
package main
import "piscine"

func main() {
  piscine.QuadC(1, 5)
}
```

## Notes

- The function uses `fmt.Print` and `fmt.Println` for output
- Each row is printed on a new line
- No spacing is added between characters within a row
- Invalid dimensions (x ≤ 0 or y ≤ 0) result in no output