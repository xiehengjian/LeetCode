
-- @Title: 组合两个表 (Combine Two Tables)
-- @Author: 846188037@qq.com
-- @Date: 2020-10-15 00:00:01
-- @Runtime: 550 ms
-- @Memory: 0 B

# Write your MySQL query statement below
select FirstName, LastName,City,State 
from Person left join Address
on Person.PersonId=Address.PersonId;
