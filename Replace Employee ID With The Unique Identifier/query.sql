SELECT eu.unique_id, e.name FROM Employees AS e LEFT JOIN EmployeeUNI AS eu ON eu.id = e.id;
