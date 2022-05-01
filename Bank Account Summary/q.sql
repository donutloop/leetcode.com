SELECT
    Ux.user_id,
    Ux.user_name,
    (credit + SUM(COALESCE(Tx.amount, 0))) AS credit,
    (
        CASE WHEN (credit + SUM(COALESCE(Tx.amount, 0))) < 0 THEN 'Yes' ELSE 'No' END
        ) AS credit_limit_breached
FROM
    Users AS Ux
        INNER JOIN (
        (
            SELECT
                user_id,
                - amount AS amount
            FROM
                Users AS u
                    LEFT JOIN Transactions AS t ON u.user_id = t.paid_by
        )
        UNION ALL
        (
            SELECT
                user_id,
                amount AS amount
            FROM
                Users AS u
                    LEFT JOIN Transactions AS t ON u.user_id = t.paid_to
        )
    ) AS Tx ON Tx.user_id = Ux.user_id
GROUP By
    Ux.user_id;