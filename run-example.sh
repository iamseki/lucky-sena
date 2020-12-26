## Run with GO
go run main.go --b=7 --e=1,2,3,4

## Migrate xlsx to Mongo
#go run main.go --f=mega_sena.xlsx

## COMPILE FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lucky-sena cmd/main.go 

## COMPILE FOR WINDOWS
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o lucky-sena.exe cmd/main.go

## COMPILE CUSTOM OUTPUT
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o randomNameBin cmd/main.go
