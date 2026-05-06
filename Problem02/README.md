# Two pointers / Hashing

> Given an array and a target value, find two elements whose sum equals that value.

---

## Problem Description

Given an array of `n` elements and a target `value`, find the **indices** of two distinct elements that add up to `value`.

---

## Input

```
Input(Array[n], n, value)
```

| Parameter | Description                        |
| --------- | ---------------------------------- |
| `Array`   | Array of `n` integers              |
| `n`       | Number of elements                 |
| `value`   | Target sum to find                 |

---

## Expected Output

```
Output(i, j)
```

| Value | Description                                          |
| ----- | ---------------------------------------------------- |
| `i`   | Index of the first element of the pair               |
| `j`   | Index of the second element of the pair (`j > i`)   |

If no valid pair exists, returns `(-1, -1)`.

---

## Brute Force Solution

The naive approach checks **every possible pair** of elements in the array. For each element at index `i`, it iterates over all subsequent elements at index `j > i` and tests whether their sum equals the target value.

```go
func SumPairs(array []int, n int, value int) (int, int) {
    for i := 0; i < n-1; i++ {
        for j := i + 1; j < n; j++ {
            if array[i]+array[j] == value {
                return i, j
            }
        }
    }

    return -1, -1
}
```

### Complexity

| Metric      | Value        |
| ----------- | ------------ |
| Time        | `O(n²)`      |
| Space       | `O(1)`       |
| Comparisons | `n(n-1) / 2` |

The nested loop means that for every new element added to the array, the number of comparisons grows **quadratically** — making this approach impractical for large inputs.

