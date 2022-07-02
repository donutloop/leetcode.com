SELECT
    name
FROM
    SalesPerson
WHERE
        name NOT IN (
        SELECT
            sp.name
        FROM
            SalesPerson AS sp
                INNER JOIN Orders AS o ON o.sales_id = sp.sales_id
                INNER JOIN Company AS c ON c.com_id = o.com_id
        WHERE
                c.name = "RED"
    );
