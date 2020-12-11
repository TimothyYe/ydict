# Binary name
BINARY=ydict
# Builds the project
build:
		GO111MODULE=on go build -o ${BINARY} -ldflags "-X main.Version=${VERSION}"
		GO111MODULE=on go test -v
# Installs our project: copies binaries
install:
		GO111MODULE=on go install
release:
		# Clean
		go clean
		rm -rf *.gz
		# Build for mac
		GO111MODULE=on go build -ldflags "-s -w -X main.Version=${VERSION}"
		tar czvf ${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}
		# Build for arm
		go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -ldflags "-s -w -X main.Version=${VERSION}"
		tar czvf ${BINARY}-arm64-${VERSION}.tar.gz ./${BINARY}
		# Build for linux
		go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w -X main.Version=${VERSION}"
		tar czvf ${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
		# Build for win
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w -X main.Version=${VERSION}"
		tar czvf ${BINARY}-win64-${VERSION}.tar.gz ./${BINARY}.exe ./win-deps/bass.dll ./win-deps/mpg123.exe
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean
		rm -rf *.gz

.PHONY:  clean build
