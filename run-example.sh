## Run tests

go test ./... -v -c # To look for tests in all dirs and subdirs in this project verbose e coverage

go test ./services -v# To run tests only in services package

go test ./... -coverprofile cover.out  # To run all tests and take it cover.out

go tool cover -html=cover.out # To see coverage profile



## Run with GO
go run main/cmd/generate-bet/*.go --b=7 --e=11,12,13,14,16,20,22,33,34,35,36,37,41,42,45,52,53,55,57,60 --p

## COMPILE generate-bet FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o generate-bet main/cmd/generate-bet/*.go 

## COMPILE generate-bet FOR WINDOWS
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o generate-bet.exe main/cmd/generate-bet/*.go 

## COMPILE scrapper FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scraper main/cmd/scraping/*.go

## COMPILE converter FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o converter main/cmd/converter/*.go

## COMPILE populate-db FOR LINUX
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o populate-db main/cmd/populate-db/*.go

## COMPILE CUSTOM OUTPUT
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o randomNameBin cmd/cli/main.go
