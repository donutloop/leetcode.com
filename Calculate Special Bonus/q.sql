SELECT employee_id,
    CASE
        WHEN employee_id % 2 = 1 AND LEFT(name, 1) != 'M' THEN salary
    ELSE 0
END AS bonus
FROM Employees;