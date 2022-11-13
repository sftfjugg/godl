Download file in Golang
==================

## Featues:
* skip TLS checking
* download batches of files concurrently
* resume incomplete downloads

## Requires
Go v1.6+

### make deps: 

安裝代码静态检查golangci-lint[https://golangci-lint.run/usage/configuration/]

`go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1`

golangci-lint 默认启用的 linter:

deadcode - 死代码检查

errcheck - 返回错误是否使用检查

gosimple - 检查代码是否可以简化

govet - 代码可疑检查，比如格式化字符串和类型不一致

ineffassign - 检查是否有未使用的代码

staticcheck - 静态分析检查

structcheck - 查找未使用的结构体字段

typecheck - 类型检查

unused - 未使用代码检查

varcheck - 未使用的全局变量和常量检查