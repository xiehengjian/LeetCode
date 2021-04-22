
-- @Title: 有趣的电影 (Not Boring Movies)
-- @Author: 846188037@qq.com
-- @Date: 2020-09-24 09:28:59
-- @Runtime: 169 ms
-- @Memory: 0 B

# Write your MySQL query statement below
SELECT * FROM cinema WHERE description != 'boring' AND mod(id,2) = 1
ORDER BY rating DESC
