```
tinygo build -o main.wasm -target wasi -scheduler=none ./cmd
wasm-opt --asyncify -O wasm.wasm -o out.wasm
```

https://github.com/GoogleChromeLabs/asyncify