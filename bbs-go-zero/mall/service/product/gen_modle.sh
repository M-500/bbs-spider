goctl model mysql ddl -src ./model/product.sql -dir ./model -c

# 生成API的命令
goctl api go -api ./api/product.api -dir ./api

# 生成rpc服务的命令
goctl rpc protoc ./rpc/product.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc