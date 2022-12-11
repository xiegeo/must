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
- Lightweight assertions. No deps, 3 digit loc count.
- `must` is designed to stand out like `unsafe`.
  So normal error handling and none panicky helper functions should live else where.
  
### Turn errors into panics

``` go
// NoError panics on error.
//
// Use this because test coverage.
func NoError(err error) {
    if err != nil {
        panic(err)
    }
}

// Value panics on error, otherwise returns the first value.
//
// Use this instead of writing a Must version of your function.
func Value[T any](out T, err error) T {
    if err != nil {
        panic(err)
    }
    return out
}
```

Also `Value2` is available for longer function signatures.
`V` and `V2` are their short aliases since the `must.` prefix already stands out enough.

### Turn assertions into panics

Functions like `True(bool)` and friends. To assert a condition or panic.

`B1`, `B2`... are generic assertions that pass their first inputs unmodified with compile time type info preserved.
They return a function to specify what assertions are made on each field.

``` go
func ExampleB3() {
    // lets say we have an io.RuneReader
    readRune := func() (r rune, size int, err error) {
        return 'a', 1, nil
    }

    // that must return 1 byte runes and never err out
    char, _, _ := must.B3(readRune())(must.Any, 1, nil)
    fmt.Println(string(char))

    // output:
    // a
}
```

Assertions have optional debug arguments, to provide additional information when
violated. Usually, just pass in the line comment as string.

### Panic test helpers

`must.Panic` and `must.Recover`. Useful for writing tests for panicky cases.

## Compatibility

Since this library is heavily dependent on generics. **1.18** is the minimum Go version supported until a new go version has a feature too good to pass.

Tagged versions follow **sematic versioning**. Untagged master/main branch is for development.

Values used in panics are only for debugging. No guarantees are provided on how panics are constructed. This means the recovered value and stack trace could change even in bugfix versions.

## Version 1.0 milestone

- Gain enough downstream usage cases to prove API fitness and stability.

## Alternatives 

See [wiki](https://github.com/xiegeo/must/wiki/Alternatives)
