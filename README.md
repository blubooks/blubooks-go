

# migrateion tool

copy .env, change db_Host to localhost

go run main.go -e .env.migrate -m down


go run main.go -u 




https://sqlines.com/oracle-to-mariadb/connect_by_prior
-- select id, title from (
WITH RECURSIVE cte_connect_by AS (
     SELECT 1 AS level, s.* 
       FROM sections s WHERE section_id is null 
     UNION ALL
     SELECT level + 1 AS level, s.* 
       FROM cte_connect_by r INNER JOIN sections  s ON  r.id = s.section_id 
  )
  SELECT id, title, level
  FROM cte_connect_by
  ORDER BY level, sort
-- ) t



https://stackoverflow.com/questions/71434494/golang-create-nested-from-linear-array-bassed-on-id-and-parent-id