GO=$(HOME)/Documents/Apps/Golang/current/bin/go
ANDROID_CC_ROOT=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/linux-x86_64/bin

run: build/ir.ArdeshirV.The-Go-Programming-Language-linux-amd64
	rm -Rf ir.ArdeshirV.The-Go-Programming-Language* >> /dev/null
	build/ir.ArdeshirV.The-Go-Programming-Language-linux-amd64

build/ir.ArdeshirV.The-Go-Programming-Language-linux-amd64: main.go
	GOOS=linux GOARCH=amd64 $(GO) build -o build/ir.ArdeshirV.The-Go-Programming-Language-linux-amd64 .

all: main.go
	[ ! -d "build" ] && mkdir build || echo >> /dev/null
	GOOS=linux GOARCH=amd64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-linux-amd64 .
	GOOS=linux GOARCH=386 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-linux-386 .
	GOOS=linux GOARCH=arm go build -o build/ir.ArdeshirV.The-Go-Programming-Language-linux-arm .
	GOOS=linux GOARCH=arm64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-linux-arm64 .

	GOOS=windows GOARCH=amd64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-windows-amd64 .
	GOOS=windows GOARCH=386 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-windows-386 .
	GOOS=windows GOARCH=arm go build -o build/ir.ArdeshirV.The-Go-Programming-Language-windows-arm .
	GOOS=windows GOARCH=arm64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-windows-arm64 .

	GOOS=freebsd GOARCH=386 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-freebsd-386 .
	GOOS=freebsd GOARCH=amd64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-freebsd-amd64 .

	GOOS=darwin GOARCH=amd64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-darwin-arm64 .

	CC=x86_64-linux-android21-clang CGO_ENABLED=1 GOOS=android GOARCH=amd64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-android-amd64 .
	CC=i686-linux-android21-clang CGO_ENABLED=1 GOOS=android GOARCH=386 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-android-386 .
	CC=armv7a-linux-androideabi21-clang CGO_ENABLED=1 GOOS=android GOARCH=arm go build -o build/ir.ArdeshirV.The-Go-Programming-Language-android-arm .
	CC=aarch64-linux-android21-clang  CGO_ENABLED=1 GOOS=android GOARCH=arm64 go build -o build/ir.ArdeshirV.The-Go-Programming-Language-android-arm64 .

clean:
	rm -Rf build/* ir.ArdeshirV.The-Go-Programming-Language*
