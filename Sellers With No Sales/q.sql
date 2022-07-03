
SELECT seller_name FROM Seller AS s WHERE seller_id NOT IN (SELECT seller_id FROM Orders AS o WHERE YEAR(o.sale_date) = 2020) ORDER BY s.seller_name