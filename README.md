## About
This is a sample app to investigate the behavior of mime.TypeByExtension() on a variety of platforms and go versions.

It is being used to investigate the following issue:
https://github.com/gofiber/fiber/issues/2383

In GoFiber/Fiber, the following code is used:
```go
c.Request().Header.Set(HeaderAccept, "text/*, application/json")
c.Accepts("json", "text") // should return "json", but on local macOS returns "text"
```

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
#### macOS Majave 10.14.6 (18G9323)
Fails on go 1.20.2 with the following error:
```
=== RUN   TestAccepts
    main_test.go:13: mimetype is text/plain; charset=utf-8, want application/json
--- FAIL: TestAccepts (0.01s)
FAIL
FAIL    github.com/sixcolors/textmime   0.016s
FAIL
```

## Investigation
The following observations were made:
- On a local macOS machine, the file extension "text" is associated with the mime type "text/plain".
- The application that registers "text" is /Applications/TextEdit.app which is a default macOS application.
- This application is not installed on the CI machines.
