# Go Functions

## Function Arguments and Return Values

- Functions can declare usable variables in the return statement

```go
func doubleValues(myVal1, myVal2 float64) (val1 float64, val2 float64) {
  val1 = myVal1 * 2
  val2 = myVal2 * 2

  // return val1, val2 // not necessary to specify values if already specified in return types
  return // OK - will return (val1, val2)
}
```

## Panic Function

```go
panic("panic message")
```
