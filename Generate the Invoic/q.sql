WITH invoice_summary AS (SELECT
                             ps.invoice_id,
                             p.product_id,
                             ps.quantity,
                             SUM(ps.quantity * p.price) AS price,
                             SUM(ps.quantity * p.price) OVER(PARTITION BY invoice_id) AS invoice_sum
                         FROM
                             Products AS p
                                 JOIN Purchases AS ps USING (product_id)
                         GROUP BY
                             invoice_id,
                             product_id
                         ORDER BY
                             invoice_sum DESC, invoice_id ASC
)

SELECT product_id, quantity, price FROM invoice_summary WHERE invoice_id = (
    SELECT invoice_id FROM invoice_summary LIMIT 1
);
