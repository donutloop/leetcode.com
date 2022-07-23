SELECT p.product_name,
       s.year,
       s.price
FROM Product AS p
         INNER JOIN Sales AS s
                    ON s.product_id = p.product_id
GROUP BY s.sale_id;