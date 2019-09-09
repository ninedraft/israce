# israce

`israce` is a helper package, which provides just two items: a boolean constant `Race` and an assert function `MustRace`. `Race` is true if program built with `-race` flag, and `MustRace` helper fails test, if program **not built** with `-race` flag. It useful then you want to guarantee that tests are run with enabled test detector.

## Installation 

```bash
go get github.com/ninedraft/israce
```

## Usage 

### In code

```go
if israce.Race {
    log.Println("program running with enabled race detector!")
}

```

### In tests
```go

func TestConcurrentCode(test *testing.T) {
    israce.MustRace(test)
    /*
        ...
    */
}
```