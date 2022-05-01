SELECT dept_name,
       COUNT(s.dept_id) AS student_number
FROM Department AS d
         LEFT JOIN Student AS s
                   ON s.dept_id = d.dept_id
GROUP BY dept_name
ORDER BY student_number DESC,
         dept_name ASC;