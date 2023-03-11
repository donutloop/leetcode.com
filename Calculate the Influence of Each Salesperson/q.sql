SELECT salesperson_id,
       name,
       IFNULL(SUM(price), 0) AS total
FROM Salesperson AS s
         LEFT JOIN Customer USING(salesperson_id)
         LEFT JOIN Sales USING(customer_id)
GROUP BY s.salesperson_id;