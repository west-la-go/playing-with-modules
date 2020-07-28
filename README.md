# playing-with-modules
Just a simple example for moving over to Go Modules from Dep


1. Check out the `Gopkg.toml` and `main.go` to see which dependencies are currently used.
    - note that the the `main.go` only shows `github.com/piquette/finance-go/quote`, but the `Gopkg.lock` shows that and `github.com/shopspring/decimal`
    - this decimal package is a dependancy of the quote package
2. Run `dep init`
    - this will create a `go.mod` file with a list of dependancies found in your code and in the `Gopkg.toml`
3. Run `go build` to create the executable, and generate the `go.sum` which acts like the `Gopkg.lock` and has all of the versions used
