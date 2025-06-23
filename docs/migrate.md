## Создание миграции
```
migrate create -ext sql -dir migrations/ <name>
```

## Применение миграции
```
migrate -path migrations/ -database "postgresql://daniluk_admin:baby@localhost:5436/backend?sslmode=disable" -verbose up
```