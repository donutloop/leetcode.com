SELECT
    left_operand,
    operator,
    right_operand,
    (CASE
         WHEN operator = ">" AND v1.value > v2.value THEN "true"
         WHEN operator = ">" AND v1.value <= v2.value THEN "false"
         WHEN operator = "=" AND v1.value = v2.value THEN "true"
         WHEN operator = "=" AND v1.value != v2.value THEN "false"
         WHEN operator = "<" AND v1.value < v2.value THEN "true"
         WHEN operator = "<" AND v1.value >= v2.value THEN "false"
        END) AS value
FROM Expressions AS e
    INNER JOIN Variables AS v1 ON v1.name = e.left_operand
    INNER JOIN Variables AS v2 ON v2.name = e.right_operand;