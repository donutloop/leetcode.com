SELECT p.product_name,
       Sum(o.unit) AS unit
FROM   products AS p
           INNER JOIN orders AS o
                      ON o.product_id = p.product_id
WHERE  Month(o.order_date) = 2
GROUP  BY p.product_name,
    Year(o.order_date),
    Month(o.order_date)
HAVING unit >= 100;