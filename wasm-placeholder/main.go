package main

import (
        "github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
        "github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
        "github.com/tidwall/gjson"
)

func main() {
        wrapper.SetCtx(
                // 插件名称
                "placeholder",
                // 为解析插件配置，设置自定义函数
                // wrapper.ParseConfigBy(parseConfig),
                // 为处理请求头，设置自定义函数
                wrapper.ProcessRequestHeadersBy(onHttpRequestHeaders),
        )
}

// 自定义插件配置
type MyConfig struct {
	
}

// 在控制台插件配置中填写的YAML配置会自动转换为JSON，此处直接从JSON这个参数里解析配置即可
func parseConfig(json gjson.Result, config *MyConfig, log wrapper.Log) error {
        // 解析出配置，更新到config中
        return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config MyConfig, log wrapper.Log) types.Action {
        return types.ActionContinue
}

