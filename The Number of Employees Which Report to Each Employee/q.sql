SELECT a.reports_to as employee_id, b.name, COUNT(a.reports_to) AS reports_count, ROUND(AVG(a.age), 0) AS average_age
FROM Employees AS a INNER JOIN Employees AS b ON a.reports_to = b.employee_id
WHERE a.employee_id IN (SELECT DISTINCT(employee_id) FROM Employees WHERE reports_to IS NOT NULL)
GROUP BY a.reports_to ORDER BY a.reports_to;
