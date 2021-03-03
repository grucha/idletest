This is a playground repo for golang + objc crosscompilation. Objc code comes from including macOS SDK bits.

On a macOS Intel the compilation works as expected but then trying to compile for M1 chip fails:
```
% GOARCH=arm64 go build
./main.go:8:2: undefined: os_specific.ShowInactiveTime
```

It fails in a way that seems like there is no code for that arch in place.
