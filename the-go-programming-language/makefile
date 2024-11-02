GO=$(HOME)/Documents/Apps/Golang/current/bin/go
ANDROID_CC_ROOT=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/linux-x86_64/bin

run: build/ir.ArdeshirV.Introducing-Go-linux-amd64
	rm -Rf ir.ArdeshirV.Introducing-Go* >> /dev/null
	build/ir.ArdeshirV.Introducing-Go-linux-amd64

build/ir.ArdeshirV.Introducing-Go-linux-amd64: main.go
	GOOS=linux GOARCH=amd64 $(GO) build -o build/ir.ArdeshirV.Introducing-Go-linux-amd64 .

all: main.go
	[ ! -d "build" ] && mkdir build || echo >> /dev/null
	GOOS=linux GOARCH=amd64 go build -o build/ir.ArdeshirV.Introducing-Go-linux-amd64 .
	GOOS=linux GOARCH=386 go build -o build/ir.ArdeshirV.Introducing-Go-linux-386 .
	GOOS=linux GOARCH=arm go build -o build/ir.ArdeshirV.Introducing-Go-linux-arm .
	GOOS=linux GOARCH=arm64 go build -o build/ir.ArdeshirV.Introducing-Go-linux-arm64 .

	GOOS=windows GOARCH=amd64 go build -o build/ir.ArdeshirV.Introducing-Go-windows-amd64 .
	GOOS=windows GOARCH=386 go build -o build/ir.ArdeshirV.Introducing-Go-windows-386 .
	GOOS=windows GOARCH=arm go build -o build/ir.ArdeshirV.Introducing-Go-windows-arm .
	GOOS=windows GOARCH=arm64 go build -o build/ir.ArdeshirV.Introducing-Go-windows-arm64 .

	GOOS=freebsd GOARCH=386 go build -o build/ir.ArdeshirV.Introducing-Go-freebsd-386 .
	GOOS=freebsd GOARCH=amd64 go build -o build/ir.ArdeshirV.Introducing-Go-freebsd-amd64 .

	GOOS=darwin GOARCH=amd64 go build -o build/ir.ArdeshirV.Introducing-Go-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o build/ir.ArdeshirV.Introducing-Go-darwin-arm64 .

	CC=x86_64-linux-android21-clang CGO_ENABLED=1 GOOS=android GOARCH=amd64 go build -o build/ir.ArdeshirV.Introducing-Go-android-amd64 .
	CC=i686-linux-android21-clang CGO_ENABLED=1 GOOS=android GOARCH=386 go build -o build/ir.ArdeshirV.Introducing-Go-android-386 .
	CC=armv7a-linux-androideabi21-clang CGO_ENABLED=1 GOOS=android GOARCH=arm go build -o build/ir.ArdeshirV.Introducing-Go-android-arm .
	CC=aarch64-linux-android21-clang  CGO_ENABLED=1 GOOS=android GOARCH=arm64 go build -o build/ir.ArdeshirV.Introducing-Go-android-arm64 .

clean:
	rm -Rf build/* ir.ArdeshirV.Introducing-Go*
