docker run --rm -ti \
    --network=host \
    -e http_proxy=${http_proxy} \
    -e https_proxy=${https_proxy} \
    -v $(pwd):/root \
    -w /root \
    higress-registry.cn-hangzhou.cr.aliyuncs.com/plugins/wasm-go-builder:go1.19-tinygo0.25.0-oras1.0.0 \
    go mod tidy && tinygo build -o ./main.wasm -scheduler=none -target=wasi ./main.go 