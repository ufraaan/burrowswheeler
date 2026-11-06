### burrows–wheeler transform

go package and cli for byte-based [burrows–wheeler transform](https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform) and its inverse.

```go
last, primary := bwt.Transform([]byte("banana"))
original, err := bwt.InverseTransform(last, primary)
```

cli reads bytes from stdin. it writes the transformed bytes to stdout and the primary index to stderr:

```sh
printf banana | go run ./cmd/bwt
printf annbaa | go run ./cmd/bwt -d -primary 3
```
