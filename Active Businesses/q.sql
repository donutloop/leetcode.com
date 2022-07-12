WITH EventsAVG AS (
    SELECT
        event_type,
        AVG(occurences) AS occurences_avg
    FROM
        Events
    GROUP BY
        event_type
)
SELECT
    DISTINCT business_id
FROM
    EventsAVG AS ea
        JOIN Events AS e ON e.event_type = ea.event_type
WHERE
        occurences > ea.occurences_avg
GROUP BY
    business_id
HAVING
        COUNT(e.event_type) > 1;
