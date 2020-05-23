# テスト用のcurlコマンド

## GET

`curl -i -X GET http://127.0.0.1:8080/record/`

## POST

`curl -i -X POST -H "Content-Type: application/json" -d '{"userId": 1, "subId": 0, "studyTime": 1000, "dateTime": "2020-05-20 03:05:20"}' http://127.0.0.1:8080/record/`
