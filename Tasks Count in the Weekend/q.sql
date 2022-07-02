SELECT Sum(CASE
               WHEN Weekday(submit_date) <= 4 THEN 1
               ELSE 0
    END) AS weekend_cnt,
       Sum(CASE
               WHEN Weekday(submit_date) > 4 THEN 1
               ELSE 0
           END) AS working_cnt
FROM   tasks;