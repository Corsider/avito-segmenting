# avito-segmenting
## Стек
Go 1.21 \
PostgreSQL 15 \
Gin \
github.com/lib/pq 

## Запуск
```
docker-compose up -d --build
```
## SQL
Приложен файл PG DUMP (avito)
```
DROP TABLE IF EXISTS userss;
DROP TABLE IF EXISTS segmentlist;
CREATE TABLE userss (user_id SERIAL PRIMARY KEY, segments INTEGER[]);
CREATE TABLE segmentlist (segment_id serial primary key , slug varchar(128));
```
## Примеры
Получить сегменты пользователя с айди 3: (GET)
```
127.0.0.1:8080/getUserSegments?user_id=3
```

Добавить юзера в сегмент (POST)
```
127.0.0.1:8080/addUserToSegment
```
```json
{
    "UserID": 9,
    "SlugsToAdd": [1,2],
    "SlugsToDelete": [3]
}
```

Получить всех юзеров (GET) - добавил для удобства
```
127.0.0.1:8080/getAllUsers
```
