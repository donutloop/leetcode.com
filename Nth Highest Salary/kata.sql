package kata


CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE o INT;
  SET o := N-1;   
  RETURN (
      SELECT CASE WHEN Salary IS NULL
        THEN NULL
        ELSE Salary
        END AS Salary FROM (select distinct Salary from Employee) AS E ORDER BY Salary DESC LIMIT 1 OFFSET o
  );
END
