# Binary name
BINARY=ydict
# Builds the project
build:
		go build -o ${BINARY}
		go test -v
# Installs our project: copies binaries
install:
		go install
release:
		# Clean
		go clean
		rm -rf *.gz
		# Build for mac
		go build
		tar czvf ydict-mac64-${VERSION}.tar.gz ./ydict
		# Build for linux
		go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
		tar czvf ydict-linux64-${VERSION}.tar.gz ./ydict
		# Build for win
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
		tar czvf ydict-win64-${VERSION}.tar.gz ./ydict.exe
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean

.PHONY:  clean build
