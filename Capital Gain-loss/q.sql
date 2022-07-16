SELECT
    s.stock_name,
    SUM(CASE WHEN s.operation="Buy" THEN -s.price ELSE s.price END) AS capital_gain_loss
FROM Stocks AS s
GROUP BY s.stock_name;