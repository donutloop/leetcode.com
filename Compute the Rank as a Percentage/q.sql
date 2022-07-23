
WITH cte
         AS (SELECT student_id,
                    department_id,
                    RANK() OVER (PARTITION BY department_id ORDER BY mark DESC) student_rank_in_the_department,
                     COUNT(*) OVER (PARTITION BY department_id) AS the_number_of_students_in_the_department
             FROM Students)

SELECT student_id,
       department_id,
       COALESCE(ROUND((student_rank_in_the_department - 1) * 100 / (the_number_of_students_in_the_department - 1), 2), 0) AS percentage
FROM cte
GROUP BY student_id;