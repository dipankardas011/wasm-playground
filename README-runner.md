# For multiple uses of wasm components

Refer:
- https://component-model.bytecodealliance.org/creating-and-consuming/composing.html
- https://component-model.bytecodealliance.org/language-support/rust.html#components-in-rust
- https://component-model.bytecodealliance.org/design/wit.html#built-in-types

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
wasm-tools compose target/wasm32-wasi/release/calculator.wasm -d ../adder/target/wasm32-wasi/release/adder.wasm -o composed-calc.wasm
```

## runner
```bash
cd runner
cargo component build --release
wasm-tools compose target/wasm32-wasi/release/runner.wasm -d ../calculator/composed-calc.wasm -o composed.wasm
wasmtime run --wasm component-model composed.wasm
```

> You need composing the WASMs of the project which you are importing
yes, runner.wasm needs its imports to either be provided by the host, a composed component, or some combination of those.  In this example, the docs:calculator/calculate@0.1.1 import is provided by calculator.wasm since wasmtime run doesn't know how to provide it.  wasmtime run only knows how to provide wasi-cli-related imports.