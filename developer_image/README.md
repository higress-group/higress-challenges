# 镜像构建
请参赛者替换以下文件：
```
├── build
│   ├── cc_deny.wasm
│   ├── ip_deny.wasm
│   └── waf_deny.wasm
├── conf
│   ├── cc_deny.yaml
│   ├── ip_deny.yaml
│   └── waf_deny.yaml
```

之后构建镜像：
```bash
docker build -t <your image name> .
```