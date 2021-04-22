
-- @Title: 超过经理收入的员工 (Employees Earning More Than Their Managers)
-- @Author: 846188037@qq.com
-- @Date: 2020-10-15 08:21:44
-- @Runtime: 283 ms
-- @Memory: 0 B

# Write your MySQL query statement below
select a.Name as `Employee` from Employee as a,Employee as b
where a.ManagerId=b.Id AND a.Salary > b.Salary
