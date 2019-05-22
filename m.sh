CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/cli

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/mac/cli

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/cli.exe 


zip -j ./build/window.zip ./build/cli.exe

rm -fr ./build/cli.exe

zip -j ./build/mac.zip ./build/mac/cli

rm -fr ./build/mac

tar -czvf ./build/linux.tar.gz -C ./build cli

rm -fr ./build/cli