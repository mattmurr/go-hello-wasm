build:
	GOOS=js GOARCH=wasm go build -o build/hello.wasm
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./build/

run:
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`./build/`)))'

clean:
	rm -rf ./build/
