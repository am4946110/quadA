# QuadD Function Documentation

## Description

`QuadD` is a Go function that prints a rectangle of a specified width and height using ASCII characters. The rectangle uses:

- **`A`** - Left corners (top-left and bottom-left)
- **`C`** - Right corners (top-right and bottom-right)  
- **`B`** - All edge cells except corners
- **` `** (space) - Interior of the rectangle

## Function Signature

```go
func QuadD(x, y int)
```

### Parameters

- `x` (int): The width of the rectangle
- `y` (int): The height of the rectangle

## Behavior

- If `x` or `y` are **not positive** (≤ 0), the function prints **nothing**
- If `x` or `y` equal **1**, the function handles edge cases appropriately

## Examples

### Example 1: QuadD(5, 3)
```
ABBBC
B   B
ABBBC
```
- Width: 5, Height: 3
- Top row: A + 3 B's + C
- Middle row: B + 3 spaces + B
- Bottom row: A + 3 B's + C

### Example 2: QuadD(5, 1)
```
ABBBC
```
- Width: 5, Height: 1
- Single row rectangle with A on left, C on right

### Example 3: QuadD(1, 1)
```
A
```
- Width: 1, Height: 1
- Single character 'A' representing a 1×1 rectangle

### Example 4: QuadD(1, 5)
```
A
B
B
B
A
```
- Width: 1, Height: 5
- First row: 'A' (top-left)
- Middle rows: 'B' (edges)
- Last row: 'A' (bottom-left)

### Example 5: QuadD(6, 4)
```
ABBBC
B    B
B    B
ABBBC
```
- Width: 6, Height: 4
- Two middle rows with spaces between B's

### Example 6: QuadD(8, 5)
```
ABBBBBC
B      B
B      B
B      B
ABBBBBC
```
- Width: 8, Height: 5
- Larger rectangle with 6 B's on top/bottom edges

## Implementation Details

1. **Input Validation**: Returns early if x or y are not positive
2. **Special Case Handling**: Handles width = 1 separately
3. **Row Construction**:
   - **First row**: 'A' + (x-2) × 'B' + 'C' (if x > 1)
   - **Middle rows**: 'B' + (x-2) × ' ' + 'B' (if x > 1)
   - **Last row**: 'A' + (x-2) × 'B' + 'C' (if x > 1 and y > 1)

## Usage Examples

### Program Test 1
```go
package main
import "piscine"

func main() {
	piscine.QuadD(5, 3)
}
```
Output:
```
ABBBC
B   B
ABBBC
```

### Program Test 2
```go
package main
import "piscine"

func main() {
	piscine.QuadD(5, 1)
}
```
Output:
```
ABBBC
```

### Program Test 3
```go
package main
import "piscine"

func main() {
	piscine.QuadD(1, 1)
}
```
Output:
```
A
```

### Program Test 4
```go
package main
import "piscine"

func main() {
	piscine.QuadD(1, 5)
}
```
Output:
```
A
B
B
B
A
```

## Comparison with QuadC

| Feature | QuadC | QuadD |
|---------|-------|-------|
| Top-left corner | A | A |
| Top-right corner | A | C |
| Bottom-left corner | C | A |
| Bottom-right corner | C | C |
| Horizontal edges | B | B |
| Vertical edges | B | B |
| Interior | space | space |

The main difference is that QuadD has different characters for the right corners (C) while QuadC uses A for all top corners and C for all bottom corners.

## Key Characteristics

- Simple and clean ASCII art rectangle
- Asymmetric corner design (A on left, C on right)
- Perfect for creating visual separators or frames
- Handles all dimension cases correctly
- No trailing spaces in output