# Pair Comparison — Find Max and Min

> Find the largest and smallest element in an array using the fewest comparisons possible.

---

## Problem Description

Given an array of `n` elements, find its **maximum** and **minimum** values while minimizing the total number of comparisons performed.

---

## Input

```
Input(Array[n], n)
```

| Parameter | Description           |
| --------- | --------------------- |
| `Array`   | Array of `n` integers |
| `n`       | Number of elements    |

---

## Expected Output

```
Output(Maximum, Minimum)
```

| Value     | Description                        |
| --------- | ---------------------------------- |
| `Maximum` | `max(Array[n])` — largest element  |
| `Minimum` | `min(Array[n])` — smallest element |

---

## Brute Force Solution

The naive approach iterates through the entire array once, performing **two comparisons per element** — one for the maximum and one for the minimum.

```go
func FindMaxMin(array []int, n int) (int, int) {
    if n == 1 {
        return array[0], array[0]
    }

    maximum := array[0]
    minimum := array[0]

    for i := 1; i < n; i++ {
        if maximum < array[i] {
            maximum = array[i]
        }
        if minimum > array[i] {
            minimum = array[i]
        }
    }

    return maximum, minimum
}
```

### Complexity

| Metric      | Value      |
| ----------- | ---------- |
| Time        | `O(n)`     |
| Space       | `O(1)`     |
| Comparisons | `2(n - 1)` |

---