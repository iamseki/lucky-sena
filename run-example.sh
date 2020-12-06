## Run with GO
go run main.go --b=10 --e=1,2,3,4

## COMPILE FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

## COMPILE FOR WINDOWS
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build .

## COMPILE CUSTOM OUTPUT
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o randomNameBin .
