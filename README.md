WASM experiments

```sh
rustup target add wasm32-unknown-unknown
(cd payload && cargo build)
(cd server && go run .)
```
