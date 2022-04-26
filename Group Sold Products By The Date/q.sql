SELECT
    sell_date,
    COUNT(DISTINCT product) AS num_sold,
    (
        SELECT
            GROUP_CONCAT(DISTINCT product
      ORDER BY
         product SEPARATOR ',')
        FROM
            Activities AS a1
        WHERE
                a1.sell_date = a2.sell_date
    )
                            AS products
FROM
    Activities AS a2
GROUP BY
    sell_date;