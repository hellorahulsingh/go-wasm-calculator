## Go to cmd/wasm directory
cd cmd/wasm
## Once we have written the code in any go file inside wasm directory
GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm


## In order to run
cd assets/
## Check go root path
go env GOROOT
## Run the below command
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

## Run the application
cd cmd/server 
## Command
go run main.go