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

| Parameter | Description           |
| --------- | --------------------- |
| `Array`   | Array of `n` integers |
| `n`       | Number of elements    |
| `value`   | Target sum to find    |

---

## Expected Output

```
Output(i, j)
```

| Value | Description                                       |
| ----- | ------------------------------------------------- |
| `i`   | Index of the first element of the pair            |
| `j`   | Index of the second element of the pair (`j > i`) |

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

The nested loop means that for every new element added to the array, the number of comparisons grows **quadratically** making this approach impractical for large inputs.

---

## Optimized Solution - Hash Map

Instead of checking every pair explicitly, this approach uses a **hash map** to remember elements already seen. For each element, it computes the value it would need as a partner (`value - array[i]`) and checks whether that partner has already been visited.

```go
func SumPairs(array []int, n int, value int) (int, int) {
    numbers := make(map[int]int)

    for i := 0; i < n; i++ {
        valuePair := value - array[i]
        val, exist := numbers[valuePair]
        if exist && val != i {
            return i, val
        }
        numbers[array[i]] = i
    }

    return -1, -1
}
```

### Why is this better?

In the brute force approach, finding a pair requires scanning all remaining elements for each position, resulting in `n(n-1)/2` comparisons and `O(n²)` time.

In the hash map approach, each element is processed **exactly once**. For every `array[i]`, the algorithm computes its complement `valuePair = value - array[i]` and performs an **O(1) hash map lookup** to check if that complement was already seen. If found, the pair is returned immediately.

This reduces the problem from a double loop to a **single linear pass**:

1. Compute the complement of the current element → `1 operation`
2. Look up the complement in the hash map → `O(1)` average
3. Store the current element in the hash map → `O(1)` average

### Concrete example with n = 8

| Approach    | Operations       |
| ----------- | ---------------- |
| Brute Force | `8 × 7 / 2 = 28` |
| Hash Map    | `≤ 8`            |

In the best case the pair is found on the second element; in the worst case the entire array is traversed once.

### Complexity

| Metric | Brute Force | Hash Map |
| ------ | ----------- | -------- |
| Time   | `O(n²)`     | `O(n)`   |
| Space  | `O(1)`      | `O(n)`   |

Unlike the pair comparison optimization which only reduced the constant factor while keeping the same `O(n)` class, the hash map approach achieves a fundamentally better **time complexity class** trading a small amount of extra memory for a dramatic reduction from quadratic to linear time. For `n = 1000`, this means roughly 500,000 operations vs. 1,000.
