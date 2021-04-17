compile:
	echo "Compiling for every OS and Platform"
	go build -o bin/api api/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 api/main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 api/main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 api/main.go

run:
	go run api/main.go