package kata

SELECT c.name as Customers FROM Customers c left join Orders o on o.CustomerId = c.Id WHERE o.Id IS NULL;
