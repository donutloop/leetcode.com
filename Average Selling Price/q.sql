SELECT p.product_id,
       ROUND(Sum(p.price * us.units) / Sum(us.units), 2) AS average_price
FROM   unitssold AS us
           INNER JOIN prices AS p
                      ON p.product_id = us.product_id
WHERE  us.purchase_date >= p.start_date
  AND us.purchase_date <= p.end_date
GROUP  BY us.product_id 