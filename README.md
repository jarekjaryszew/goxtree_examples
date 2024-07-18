# goXtree example

Proof of concept app build with goXtree (https://github.com/jarekjaryszew/goxtree).

## Usage
Please clone this repo next to goXtree.
```
your_dir
|- goxtree
|- goxtree_examples
```

Go to example folder and build it:
```sh
cd goxtree_examples/simple
GOOS=js GOARCH=wasm go build -o main.wasm
```

Serve the content of the directory using http. For example:
```sh
python3 -m http.server 8000
```
