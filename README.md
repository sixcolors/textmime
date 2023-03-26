## About
A test of mime.TypeByExtension() on a variety of platforms and go versions.

## Results

### CI
#### Linux
Passes on all versions of go.

#### Windows
Passes on all versions of go.

### macOS
Passes on all versions of go.

### Local

#### macOS Ventura 13.2.1 (22D68)
Fails on go 1.20.2 with the following error:
```
=== RUN   TestAccepts
    main_test.go:13: mimetype is text/plain; charset=utf-8, want application/json
--- FAIL: TestAccepts (0.00s)
FAIL
FAIL    github.com/sixcolors/textmime   0.329s
FAIL
```