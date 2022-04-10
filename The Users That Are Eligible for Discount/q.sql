CREATE PROCEDURE getUserIDs(startDate DATE, endDate DATE, minAmount INT)
BEGIN
SELECT DISTINCT(user_id)
FROM Purchases
WHERE time_stamp <= endDate AND time_stamp >= startDate AND amount >= minAmount
ORDER BY user_id ASC;
END