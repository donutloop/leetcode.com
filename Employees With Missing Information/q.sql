SELECT employee_id
FROM (
         (SELECT e.employee_id,
                 name,
                 salary
          FROM Employees AS e
                   LEFT JOIN Salaries AS s ON s.employee_id=e.employee_id)
         UNION
         (SELECT s.employee_id,
                 name,
                 salary
          FROM Salaries AS s
                   LEFT JOIN Employees AS e ON s.employee_id=e.employee_id)) AS ec
WHERE name IS NULL
   OR salary IS NULL
ORDER BY employee_id;