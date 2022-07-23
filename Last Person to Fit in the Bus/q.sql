WITH cte AS
(
   SELECT
      person_name,
      SUM(weight) OVER (
   ORDER BY
      turn ROWS UNBOUNDED PRECEDING) AS running_total
   FROM
      Queue
)

SELECT
   person_name
FROM
   cte
WHERE
   running_total <= 1000 ORDER BY running_total DESC LIMIT 1;
