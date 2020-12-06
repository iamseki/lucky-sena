CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o lucky .
./lucky --b=7 --e=2,16,19,31,43,60,20,27,35,39,50,59,5,10,29,34,41

go run main.go --b=10 --e=1,2,3,4