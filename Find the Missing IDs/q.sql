WITH RECURSIVE nums AS (
    SELECT 1 AS value
UNION ALL
SELECT value + 1 AS value
FROM nums
WHERE nums.value < (SELECT MAX(customer_id) FROM Customers))

SELECT value AS ids FROM nums WHERE value NOT IN(SELECT customer_id FROM Customers) ORDER BY value ASC;