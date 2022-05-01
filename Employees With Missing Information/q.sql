(SELECT e.employee_id
 FROM Employees AS e
          LEFT JOIN Salaries AS s ON s.employee_id=e.employee_id
 WHERE salary IS NULL)
UNION
(SELECT s.employee_id
 FROM Salaries AS s
          LEFT JOIN Employees AS e ON s.employee_id=e.employee_id
 WHERE name IS NULL)
ORDER BY employee_id