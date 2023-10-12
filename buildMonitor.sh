rm osxmonitor tuxmonitor winmonitor.exe
export GO111MODULE=off
GOOS=darwin GOARCH=amd64 go build -o osxmonitor monitorhttp.go
GOOS=linux GOARCH=amd64 go build -o tuxmonitor monitorhttp.go
GOOS=windows GOARCH=amd64 go build -o winmonitor.exe monitorhttp.go

