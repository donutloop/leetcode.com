select
    company_id,
    employee_id,
    employee_name,
    (CASE
         WHEN (MAX(salary) OVER(PARTITION BY company_id) >= 1000 AND MAX(salary) OVER(PARTITION BY company_id) <= 10000) THEN ROUND((salary-(salary*0.24)), 0)
         WHEN (MAX(salary) OVER(PARTITION BY company_id) < 1000) THEN ROUND(salary,0)
         WHEN (MAX(salary) OVER(PARTITION BY company_id) > 10000) THEN ROUND((salary-(salary*0.49)),0)
        END) AS salary
from Salaries;