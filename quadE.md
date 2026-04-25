# QuadE Function Documentation

## Description

`QuadE` is a Go function that prints a rectangle of a specified width and height using ASCII characters. The rectangle uses a **reversed corner pattern**:

- **Top row**: `A` on left, `C` on right with `B`'s between
- **Middle rows**: `B` on left and right sides with spaces between
- **Bottom row**: `C` on left, `A` on right with `B`'s between (corners are reversed!)

This creates a visually interesting asymmetric rectangle.

## Function Signature

```go
func QuadE(x, y int)
```

### Parameters

- `x` (int): The width of the rectangle
- `y` (int): The height of the rectangle

## Behavior

- If `x` or `y` are **not positive** (≤ 0), the function prints **nothing**
- If `x` or `y` equal **1**, the function handles edge cases appropriately
- **Key Feature**: The top and bottom rows have reversed corners (A↔C)

## Examples

### Example 1: QuadE(5, 3)
```
ABBBC
B   B
CBBBA
```
- Width: 5, Height: 3
- Top row: A + BBB + C (A on left, C on right)
- Middle row: B + 3 spaces + B
- Bottom row: C + BBB + A (reversed: C on left, A on right)

### Example 2: QuadE(5, 1)
```
ABBBC
```
- Width: 5, Height: 1
- Single row with A on left, C on right
- No bottom row since height is 1

### Example 3: QuadE(1, 1)
```
A
```
- Width: 1, Height: 1
- Single character 'A' representing a 1×1 rectangle

### Example 4: QuadE(1, 5)
```
A
B
B
B
C
```
- Width: 1, Height: 5
- First row: 'A' (top-left)
- Middle rows: 'B' (edges)
- Last row: 'C' (bottom-left, reversed)

### Example 5: QuadE(6, 4)
```
ABBBC
B    B
B    B
CBBBA
```
- Width: 6, Height: 4
- Two middle rows with spaces between B's
- Notice the reversed corners on top and bottom

### Example 6: QuadE(8, 5)
```
ABBBBBC
B      B
B      B
B      B
CBBBBBA
```
- Width: 8, Height: 5
- Larger rectangle showing the reversed corner pattern clearly

## Implementation Details

1. **Input Validation**: Returns early if x or y are not positive
2. **Special Case Handling**: Handles width = 1 separately
3. **Row Construction**:
   - **First row**: 'A' + (x-2) × 'B' + 'C'
   - **Middle rows**: 'B' + (x-2) × ' ' + 'B'
   - **Last row** (if y > 1): 'C' + (x-2) × 'B' + 'A' (reversed!)

## Usage Examples

### Program Test 1
```go
package main
import "piscine"

func main() {
	piscine.QuadE(5, 3)
}
```
Output:
```
ABBBC
B   B
CBBBA
```

### Program Test 2
```go
package main
import "piscine"

func main() {
	piscine.QuadE(5, 1)
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
	piscine.QuadE(1, 1)
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
	piscine.QuadE(1, 5)
}
```
Output:
```
A
B
B
B
C
```

## Comparison with QuadD

| Feature | QuadD | QuadE |
|---------|-------|-------|
| Top-left corner | A | A |
| Top-right corner | C | C |
| Bottom-left corner | A | C |
| Bottom-right corner | C | A |
| Horizontal edges | B | B |
| Vertical edges | B | B |
| Interior | space | space |
| **Key Difference** | Same corners on top and bottom | Reversed corners on bottom |

The main difference is that QuadE has **reversed corners** on the bottom row compared to the top row, creating an interesting asymmetric visual pattern.

## Key Characteristics

- Asymmetric rectangle design with reversed corner characters
- Creates a dynamic visual effect with the corner swap
- Perfect for creative ASCII art borders
- Handles all dimension cases correctly
- No trailing spaces in output
- Clean and simple implementation

## Edge Cases

| Input | Output | Notes |
|-------|--------|-------|
| (1, 1) | A | Single character |
| (1, 5) | A / B / B / B / C | Vertical line with reversed ends |
| (5, 1) | ABBBC | Horizontal line, no bottom row |
| (0, 5) | (nothing) | Invalid: x not positive |
| (5, 0) | (nothing) | Invalid: y not positive |
| (2, 2) | AB / CA | Minimal 2×2 rectangle |

## Visual Patterns

The reversed corners create an interesting visual progression:

```
Small (3×3):    Medium (5×3):   Large (7×4):
ABB             ABBBC           ABBBBBC
B B             B   B           B      B
CBA             CBBBA           B      B
                                CBBBBBA
```

Notice how the corners swap between top and bottom!