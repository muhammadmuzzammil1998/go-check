# go-check
A simple error check for Go errors

## What does it do
It simply checks if `error` provided as argument is `nil` or not. If it isn't `nil`, you can fire a function to handle it. It can be any function like `Fatalln`. You can define an anonymous function as well.

## How to use it

### `go get` it
```sh
$ go get muzzammil.xyz/go-check
```

### Use in code
```go
...

err := someFuncThatMayReturnError()

// Usage 1
check.Error(err, log.Fatalln)

// Usage 2
check.Error(err, func(v ...interface{}) {
  fmt.Println("Got error", v)
})

...
```

### But why?

I didn't want to use this everytime -
```go
err := someFuncThatMayReturnError()
if err != nil {
  log.Fatalln(err)
}
```
