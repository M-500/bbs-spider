-- "article:12": {"read_cnt":1,"like_cnt":0 ...}
-- 目标哈希表的名字 "article:12"
local key = KEYS[1]
-- 目标要操作的哈希表的字段 "read_cnt"
local tagKey = ARGV[1]

local delta = tonumber(ARGV[2])

local exists = redis.call("EXISTS",key)

if exists == 1 then
    redis.call("HINCRBY",key,tagKey,delta)
    -- 操作+1 成功 返回1
    return 1
else
    -- 如果Key 不存在，直接返回0 不做任何赋值操作，赋值交给后续业务去做
    return 0
end