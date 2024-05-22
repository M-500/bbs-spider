


# 生成 api 项目的指令
goctl api go -api ./api/user.api -dir ./api -style=goZero

# 生成 rpc 服务的命令
goctl rpc protoc ./rpc/xxx.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc