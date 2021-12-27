package kata

select Name as Employee from Employee as e where Salary > (select Salary from Employee as ie where ie.Id = e.ManagerId)
