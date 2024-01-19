# For multiple uses of wasm components

## adder
```bash
cd adder
cargo component build --release
wasm-tools component wit ./target/wasm32-wasi/release/adder.wasm
```

## calculator
```bash
cd calculator
cargo component build --release
wasm-tools component wit ./target/wasm32-wasi/release/calculator.wasm
```

## runner
```bash
cd runner
cargo component build --release
wasmtime run --wasm component-model ./target/wasm32-wasi/release/runner.wasm
```