

const hosts = {
    prod: 'https://weapp.shujingyun.com/ldxf',
    test: "https://www.keyike.ltd",
    dev: 'http://127.0.0.1:8181',
    // dev: 'https://weapp.shujingyun.com/ldxf',
}


// 定义变量(是否是测试环境还是测试其他环境)
// let env = process.env.WEB_ENV || "prod";
let env = process.env.WEB_ENV || "dev";



module.exports = {
    host: hosts[env],//
    version: "yangs",
    env: env
}
