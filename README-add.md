# Add
to try it out
you need to

```bash
cargo install cargo-component
```

```bash
cd add
cargo component build --release
```

it attaches the wit to the implementation into a single wasm component

to use it we need some wasm runner (Sorry for any incorrect word choices)

```bash
# cd  into the submodule to get a prefabricated runner
cd component-docs/component-model/examples/example-host/
```

```bash
cargo run --release -- 1 2 <add the location to the generated wasm>add/target/wasm32-wasi/release/add.wasm
```

> it should return you with result
