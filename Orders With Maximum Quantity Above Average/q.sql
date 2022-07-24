WITH cte AS
         (SELECT order_id,
                 (SUM(quantity)/COUNT(product_id)) AS average_quantity,
                 MAX(quantity) AS max_quantity
          FROM OrdersDetails
          GROUP BY order_id)

SELECT order_id
FROM cte AS c
WHERE c.max_quantity >
      (SELECT MAX(average_quantity)
       FROM cte);