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
    -- 自增不成功
    return 0
end