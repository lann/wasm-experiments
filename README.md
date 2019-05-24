WASM experiments

```sh
# Build WASM payload.
rustup target add wasm32-unknown-unknown
(cd payload && cargo build)

# Run payload in server.
(cd server && go run .)

# Run payload in browser.
(cd browser && ./serve.py)
```
