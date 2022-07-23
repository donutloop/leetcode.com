

WITH cte AS (
    (
        SELECT
            c.Wimbledon AS player_id
        FROM
            Championships AS c
    )
    UNION ALL
    (
        SELECT
            c.Fr_open AS player_id
        FROM
            Championships AS c
    )
    UNION ALL
    (
        SELECT
            c.US_open AS player_id
        FROM
            Championships AS c
    )
    UNION ALL
    (
        SELECT
            c.Au_open AS player_id
        FROM
            Championships AS c
    )
)
SELECT
    p.player_id,
    player_name,
    COUNT(c.player_id) AS grand_slams_count
FROM
    Players AS p
        INNER JOIN cte AS c ON p.player_id = c.player_id
GROUP BY
    p.player_id;


