((
     SELECT
         b.seat_id
     FROM
         Cinema AS a
             INNER JOIN
         Cinema AS b
         ON b.seat_id + 1 = a.seat_id
     WHERE
             a.free = 1
       AND b.free = 1)
 UNION
 (
     SELECT
         b.seat_id
     FROM
         Cinema AS a
             INNER JOIN
         Cinema AS b
         ON b.seat_id - 1 = a.seat_id
     WHERE
             a.free = 1
       AND b.free = 1))
    ORDER BY
        seat_id