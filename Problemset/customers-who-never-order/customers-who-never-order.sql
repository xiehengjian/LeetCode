
-- @Title: 从不订购的客户 (Customers Who Never Order)
-- @Author: 846188037@qq.com
-- @Date: 2020-10-15 08:27:55
-- @Runtime: 384 ms
-- @Memory: 0 B

# Write your MySQL query statement below
select customers.name as `Customers`
from customers 
where customers.id not in 
(
    select customerid from orders
);
