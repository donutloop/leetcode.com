
WITH cte AS (
    (
        SELECT
            p.player_id,
            p.player_name,
            1 AS win
        FROM
            Players AS p
                INNER JOIN Championships AS c ON p.player_id = c.Wimbledon
    )
    UNION ALL
    (
        SELECT
            p.player_id,
            p.player_name,
            1 AS win
        FROM
            Players AS p
                INNER JOIN Championships AS c ON p.player_id = c.Fr_open
    )
    UNION ALL
    (
        SELECT
            p.player_id,
            p.player_name,
            1 AS win
        FROM
            Players AS p
                INNER JOIN Championships AS c ON p.player_id = c.US_open
    )
    UNION ALL
    (
        SELECT
            p.player_id,
            p.player_name,
            1 AS win
        FROM
            Players AS p
                INNER JOIN Championships AS c ON p.player_id = c.Au_open
    )
)
SELECT
    player_id,
    player_name,
    COUNT(win) AS grand_slams_count
FROM
    cte
GROUP BY
    player_id;

