local key = KEYS[1]
-- 对应 hincrby中的field字段
local cntKey = ARGV[1]
local delta = tonumber(ARGV[2])
local exists = redis.call("EXISTS",key)

if exists == 1 then
    redis.call("HINCRBY",key,cntKey,delta)
    --  如果自增成功，返回1
    return 1
else
    -- 自增不成功 思考一下为什么这里对于Key不存在直接返回0(提示:为了性能，同时为了缓存一致性)
    return 0
end