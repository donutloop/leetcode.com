SELECT ROUND((SELECT COUNT(*) FROM Delivery AS d2 WHERE d2.order_date = d2.customer_pref_delivery_date)/COUNT(*)*100, 2) AS immediate_percentage FROM Delivery AS d1;
