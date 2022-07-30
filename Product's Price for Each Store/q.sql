SELECT
    product_id,
    MAX(store1) AS store1,
    MAX(store2) AS store2,
    MAX(store3) AS store3
FROM
    (
        SELECT
            product_id,
            price AS store1,
            NULL AS store2,
            NULL AS store3
        FROM
            Products
        WHERE
                store = "store1"
        UNION ALL
        SELECT
            product_id,
            NULL AS store1,
            price AS store2,
            NULL AS store3
        FROM
            Products
        WHERE
                store = "store2"
        UNION ALL
        SELECT
            product_id,
            NULL AS store1,
            NULL AS store2,
            price AS store3
        FROM
            Products
        WHERE
                store = "store3"
    ) AS transpose
GROUP BY
    product_id
