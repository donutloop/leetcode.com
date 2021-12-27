package kata

DELETE p1 FROM Person p1
INNER JOIN Person p2
WHERE
    p1.id > p2.id AND 
    p1.email = p2.email;
