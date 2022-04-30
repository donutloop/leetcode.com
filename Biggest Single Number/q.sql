SELECT (SELECT num
        FROM   mynumbers
        GROUP  BY num
        HAVING Count(num) = 1
        ORDER  BY num DESC
           LIMIT  1) AS NUM; 