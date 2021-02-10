## Run tests

go run ./... -v # To look for tests in all dirs and subdirs in this project

go run ./services -v # To run tests only in services package

go test ./tests/...  -v # To run all tests in tests subfolders

## Run with GO
go run main/cmd/generate-bet/*.go --b=7 --e=11,12,13,14,16,20,22,33,34,35,36,37,41,42,45,52,53,55,57,60 --p

## COMPILE generate-bet FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o generate-bet main/cmd/generate-bet/*.go 

## COMPILE generate-bet FOR WINDOWS
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o generate-bet.exe main/cmd/generate-bet/*.go 

## COMPILE scrapper FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scraper main/cmd/scraping/*.go

## COMPILE scrapper FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o converter main/cmd/converter/*.go

## COMPILE CUSTOM OUTPUT
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o randomNameBin cmd/cli/main.go
