CREATE FUNCTION getUserIDs(startDate DATE, endDate DATE, minAmount INT) RETURNS INT
BEGIN
RETURN (
    SELECT COUNT(DISTINCT(user_id))
    FROM Purchases
    WHERE time_stamp <= endDate AND time_stamp >= startDate AND amount >= minAmount
);
END