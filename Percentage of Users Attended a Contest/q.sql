SELECT
  contest_id,
  ROUND(
    COUNT(*) / (
      SELECT
        COUNT(*)
      FROM
        Users
    ) * 100,
    2
  ) AS percentage
FROM
  Register AS r
  INNER JOIN Users AS u ON u.user_id = r.user_id
GROUP BY
  contest_id
ORDER BY
  percentage DESC,
  contest_id ASC;
