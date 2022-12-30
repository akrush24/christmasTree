# christmasTree

## Only run with go
```bash
git clone git@github.com:akrush24/christmasTree.git
go run ./christmasTree/christmasTree.go "UserName"
```

## Build and run
```bash
go build -ldflags="-s" ./christmasTree.go
./christmasTree
```

## Build for other arch
```bash
go tool dist list
env GOOS=linux GOARCH=amd64 go build -o christmasTree.linux.amd64 -ldflags="-s" ./christmasTree.go
env GOOS=darwin GOARCH=arm64 go build -o christmasTree.darwin.arm64 -ldflags="-s" ./christmasTree.go
env GOOS=darwin GOARCH=amd64 go build -o christmasTree.darwin.amd64 -ldflags="-s" ./christmasTree.go
env GOOS=windows GOARCH=amd64 go build -o christmasTree.windows.amd64 -ldflags="-s" ./christmasTree.go
```

## Result
![image](https://user-images.githubusercontent.com/3369193/210077025-4843727e-097f-49db-ad45-ab990fb7a8bc.png)
