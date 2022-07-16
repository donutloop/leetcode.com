SELECT a.team_name AS home_team,
       b.team_name AS away_team
FROM Teams AS a
INNER JOIN Teams AS b ON b.team_name != a.team_name;