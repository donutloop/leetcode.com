SELECT project_id,
       ROUND(SUM(e.experience_years)/COUNT(p.employee_id), 2) AS average_years
FROM Project AS p
         INNER JOIN Employee AS e ON p.employee_id = e.employee_id
GROUP BY project_id;