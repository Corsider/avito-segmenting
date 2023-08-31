# avito-segmenting
## Запуск
```
docker-compose up -d --build
```
## SQL
```
DROP TABLE IF EXISTS userss;
DROP TABLE IF EXISTS segmentlist;
CREATE TABLE userss (user_id SERIAL PRIMARY KEY, segments INTEGER[]);
CREATE TABLE segmentlist (segment_id serial primary key , slug varchar(128));
```
