WITH summary AS (
    SELECT
        MIN(activity_date) AS activity_date,
        user_id
    FROM
        Traffic
    WHERE
            activity = "login"
    GROUP BY
        user_id
)

SELECT
    activity_date AS login_date,
    COUNT(DISTINCT user_id) AS user_count
FROM
    summary
WHERE
    (
                activity_date >= (DATE_SUB("2019-06-30", INTERVAL 90 DAY))
            AND activity_date <= "2019-06-30"
        )
GROUP BY
    activity_date;
