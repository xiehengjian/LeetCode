
-- @Title: 查找重复的电子邮箱 (Duplicate Emails)
-- @Author: 846188037@qq.com
-- @Date: 2020-10-14 17:03:29
-- @Runtime: 294 ms
-- @Memory: 0 B

# Write your MySQL query statement below
select Email from
(select Email,count(Email) as num
from Person 
group by Email) as statitic where num>1;
