# Home-Service
[![Build](https://github.com/NRKA/backend-bootcamp-assignment-2024/actions/workflows/build.yml/badge.svg)](https://github.com/NRKA/backend-bootcamp-assignment-2024/actions/workflows/build.yml)

## Установка и конфигурация
+ Склонировать репозиторий:
  ```
  git clone https://github.com/NRKA/home-service.git
  ```
+ Запустить *docker compose* из корневой директории
  ```make
  docker-compose up -d
  ```
+ Поднятие *миграции*
  ```make
  make migration-up
  ```
## Использование
### Сервис поддерживает 8 эндпоинтов:
+ `GET    /dummyLogin -- (no Auth)`
+ `POST   /login -- (no Auth)`
+ `POST   /register -- (no Auth)`
+ `GET    /house/{id} -- (Auth only)`
+ `POST    /house/{id}/subscribe -- (Auth only)`
+ `POST   /flat/create -- (Auth only)`
+ `POST   /house/create -- (Moderations only)`
+ `POST   /flat/update -- (Moderations only)`


## CURL Запросы
### dummyLogin
```
curl -X GET localhost:8080/dummyLogin?user_type=client -i
curl -X GET localhost:8080/dummyLogin?user_type=moderator -i
```
### login
```
curl -X POST localhost:8080/login -d '{"id":id, "password":"password"}' -i
```
### register
```
curl -X POST localhost:8080/register -d '{"email": "test@gmail.com","password": "Секретная строка","user_type": "moderator"}' -i
```
### houseID
```
curl -X GET localhost:8080/house/id -H "Authorization: Bearer token" -i
```
### houseID subscribe
```
curl -X POST localhost:8080/house/id/subscribe -d '{"email":"test@gmail.com"}' -H "Authorization: Bearer token" -i
```
### flatCreate
```
curl -X POST localhost:8080/flat/create -d '{"number":123,"house_id":123,"price":123213,"rooms":4}' -H "Authorization: Bearer token" -i
```
### houseCreate
```
curl -X POST localhost:8080/house/create -d '{"address":"Лесная улица, 7, Москва, 125196","year": 2000,"developer":"Мэрия города"}' -H "Authorization: Bearer token" -i
```
### flatUpdate
```
curl -X POST localhost:8080/flat/update -d '{"id":123456,"status":"approved"}' -H "Authorization: Bearer token" -i
```