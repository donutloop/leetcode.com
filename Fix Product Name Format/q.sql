SELECT
    TRIM(LOWER(product_name)) AS product_name,
    DATE_FORMAT(sale_date, '%Y-%m') AS sale_date,
    COUNT(*) AS total
FROM
    Sales
GROUP BY
    1, 2
ORDER BY
    product_name ASC,
    sale_date ASC;