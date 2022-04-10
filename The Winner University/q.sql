SELECT
    CASE WHEN ((SELECT COUNT(*) FROM NewYork WHERE score >= 90) < (SELECT COUNT(*) FROM California WHERE score >= 90))
             THEN "California University"
         WHEN ((SELECT COUNT(*) FROM NewYork WHERE score >= 90) = (SELECT COUNT(*) FROM California WHERE score >= 90))
             THEN "No Winner"
         ELSE "New York University"
        END AS winner
FROM DUAL;