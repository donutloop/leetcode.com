(SELECT num
 FROM mynumbers GROUP  BY num
 HAVING Count(num) = 1
 ORDER  BY num DESC
     LIMIT 1)
UNION
(SELECT NULL)
LIMIT 1