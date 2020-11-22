del out
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
go tool cgo main.go
go build -buildmode=c-shared -o out\Patch.dll