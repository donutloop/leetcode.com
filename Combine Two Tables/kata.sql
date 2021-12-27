package kata

# Write your MySQL query statement below
SELECT p.FirstName, p.LastName, City, State FROM Person p LEFT JOIN Address a ON a.PersonId = p.PersonId; 
