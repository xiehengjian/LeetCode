
-- @Title: 大的国家 (Big Countries)
-- @Author: 846188037@qq.com
-- @Date: 2020-09-24 09:22:45
-- @Runtime: 206 ms
-- @Memory: 0 B

# Write your MySQL query statement below
SELECT name,population,area FROM World WHERE area > 3000000 OR population > 25000000
