
-- @Title: 变更性别 (Swap Salary)
-- @Author: 846188037@qq.com
-- @Date: 2020-10-14 17:16:45
-- @Runtime: 183 ms
-- @Memory: 0 B

# Write your MySQL query statement below
update salary 
set sex=case sex when 'm' then 'f' else 'm' end;
