# 生成 api 项目的指令
goctl api go -api ./api/*.api -dir ./api -style=goZero

goctl rpc protoc ./rpc/*.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc -style=goZero

goctl model mysql ddl -src ./model/product.sql -dir ./model -c
# 如果要生成带缓存的就用这个
goctl model mysql ddl -src ./model/product.sql -dir ./model -c --cache true