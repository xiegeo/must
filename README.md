# must [![Go Reference](https://pkg.go.dev/badge/.svg)](https://pkg.go.dev/github.com/xiegeo/must)

When you don't need error handling

I found myself using this snipped lots.

``` go
func must[T any](out T, err error) T {
    if err != nil {
        panic(err)
    }
    return out
}
```

Occasionally a little more, so made a home for them, DRY.

## Scope

- Error handling is not always necessary. Panic on error should be
  easier than ignoring them.
- Lightweight assertions.

### Turn errors into panics

``` go
// NoError panics on error.
func NoError(err error) {
    if err != nil {
        panic(err)
    }
}

// Value panics on error, otherwise returns the first value.
func Value[T any](out T, err error) T {
    if err != nil {
        panic(err)
    }
    return out
}
```

Also Value2 is available for longer function signatures.

### Turn assertions into panics

`must.True` and friends.

Assertions have optional debug arguments, to provide additional information when
violated. Usually, just pass in the line comment as string.

### Panic test helpers

`must.Panic` and `must.Recover`. Useful for writing tests for panics cases.

## Compatibility

Since this library is heavily dependent on generics. 1.18 is the minimum Go version supported for now.

Tagged versions will follow sematic versioning. Untagged master/main branch is for development.

Values set in panics are only for debugging. No guarantees are provided on how panics are argumented.

## Planning

Generic assertions ;-)
