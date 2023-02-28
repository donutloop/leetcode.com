WITH sales_matrix AS
         (SELECT item_category AS CATEGORY,
                 SUM(quantity) AS MONDAY,
                 0 AS TUESDAY,
                 0 AS WEDNESDAY,
                 0 AS THURSDAY,
                 0 AS FRIDAY,
                 0 AS SATURDAY,
                 0 AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 0
          GROUP BY item_category
          UNION SELECT item_category AS CATEGORY,
                       0 AS MONDAY,
                       SUM(quantity) AS TUESDAY,
                       0 AS WEDNESDAY,
                       0 AS THURSDAY,
                       0 AS FRIDAY,
                       0 AS SATURDAY,
                       0 AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 1
          GROUP BY item_category
          UNION ALL SELECT item_category AS CATEGORY,
                           0 AS MONDAY,
                           0 AS TUESDAY,
                           SUM(quantity) AS WEDNESDAY,
                           0 AS THURSDAY,
                           0 AS FRIDAY,
                           0 AS SATURDAY,
                           0 AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 2
          GROUP BY item_category
          UNION ALL SELECT item_category AS CATEGORY,
                           0 AS MONDAY,
                           0 AS TUESDAY,
                           0 AS WEDNESDAY,
                           SUM(quantity) AS THURSDAY,
                           0 AS FRIDAY,
                           0 AS SATURDAY,
                           0 AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 3
          GROUP BY item_category
          UNION ALL SELECT item_category AS CATEGORY,
                           0 AS MONDAY,
                           0 AS TUESDAY,
                           0 AS WEDNESDAY,
                           0 AS THURSDAY,
                           SUM(quantity) AS FRIDAY,
                           0 AS SATURDAY,
                           0 AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 4
          GROUP BY item_category
          UNION ALL SELECT item_category AS CATEGORY,
                           0 AS MONDAY,
                           0 AS TUESDAY,
                           0 AS WEDNESDAY,
                           0 AS THURSDAY,
                           0 AS FRIDAY,
                           SUM(quantity) AS SATURDAY,
                           0 AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 5
          GROUP BY item_category
          UNION ALL SELECT item_category AS CATEGORY,
                           0 AS MONDAY,
                           0 AS TUESDAY,
                           0 AS WEDNESDAY,
                           0 AS THURSDAY,
                           0 AS FRIDAY,
                           0 AS SATURDAY,
                           SUM(quantity) AS SUNDAY
          FROM Items
                   JOIN Orders USING (item_id)
          WHERE WEEKDAY(order_date) = 6
          GROUP BY item_category
          UNION ALL SELECT item_category AS CATEGORY,
                           0 AS MONDAY,
                           0 AS TUESDAY,
                           0 AS WEDNESDAY,
                           0 AS THURSDAY,
                           0 AS FRIDAY,
                           0 AS SATURDAY,
                           0 AS SUNDAY
          FROM Items AS i
                   LEFT JOIN Orders AS o ON i.item_id = o.item_id
          WHERE o.item_id IS NULL
          GROUP BY item_category)

SELECT CATEGORY,
       SUM(MONDAY) AS MONDAY,
       SUM(TUESDAY) AS TUESDAY,
       SUM(WEDNESDAY) AS WEDNESDAY,
       SUM(THURSDAY) AS THURSDAY,
       SUM(FRIDAY) AS FRIDAY,
       SUM(SATURDAY) AS SATURDAY,
       SUM(SUNDAY) AS SUNDAY
FROM sales_matrix
GROUP BY CATEGORY
ORDER BY CATEGORY;