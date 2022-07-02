SELECT a.product_id,
       Round(Sum(a.price) / Sum(a.units), 2) AS average_price
FROM   (SELECT p.product_id,
               p.price * us.units AS price,
               us.units
        FROM   unitssold AS us
                   INNER JOIN prices AS p
                              ON p.product_id = us.product_id
        WHERE  us.purchase_date >= p.start_date
          AND us.purchase_date <= p.end_date
        GROUP  BY us.product_id,
                  p.price,
                  us.units,
                  us.purchase_date
       ) AS a
GROUP  BY product_id 