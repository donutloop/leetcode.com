
WITH cte AS (SELECT
    id,
    case when month = "Jan" then revenue end as Jan_Revenue,
    case when month = "Feb" then revenue end as Feb_Revenue,
    case when month = "Mar" then revenue end as Mar_Revenue,
    case when month = "Apr" then revenue end as Apr_Revenue,
    case when month = "May" then revenue end as May_Revenue,
    case when month = "Jun" then revenue end as Jun_Revenue,
    case when month = "Jul" then revenue end as Jul_Revenue,
    case when month = "Aug" then revenue end as Aug_Revenue,
    case when month = "Sep" then revenue end as Sep_Revenue,
    case when month = "Oct" then revenue end as Oct_Revenue,
    case when month = "Nov" then revenue end as Nov_Revenue,
    case when month = "Dec" then revenue end as Dec_Revenue
FROM Department)


SELECT id,
    MAX(Jan_Revenue) AS Jan_Revenue,
    MAX(Feb_Revenue) AS Feb_Revenue,
    MAX(Mar_Revenue) AS Mar_Revenue,
    MAX(Apr_Revenue) AS Apr_Revenue,
    MAX(May_Revenue) AS May_Revenue,
    MAX(Jun_Revenue) AS Jun_Revenue,
    MAX(Jul_Revenue) AS Jul_Revenue,
    MAX(Aug_Revenue) AS Aug_Revenue,
    MAX(Sep_Revenue) AS Sep_Revenue,
    MAX(Oct_Revenue) AS Oct_Revenue,
    MAX(Nov_Revenue) AS Nov_Revenue,
    MAX(Dec_Revenue) AS Dec_Revenue
FROM cte GROUP BY id;