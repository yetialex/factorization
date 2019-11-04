prime factorization
======

Web сервис разложения натуральных чисел на простые множители

Для сборки сервиса необходимы make и docker.

Сборка осуществляется в отдельном контейнере.

## Быстрый старт

    make build
    docker-compose up

## Запуск тестов

    make test
    
## Примеры запросов

    Запрос:
        curl -H Content-Type:application/json -H Accept:application/json -X POST -d '{"number":"5555"}' http://127.0.0.1:8090/factorize
    Ответ:
        {"success":true,"message":"5555 = 5 * 11 * 101","payload":"Execution time: 20.1µs"}
    
    Запрос:
        curl -X GET http://127.0.0.1:8090/factorize/111111111111111111111111111111
    Ответ:    
        {"success":true,"message":"111111111111111111111111111111 = 3 * 7 * 11 * 13 * 31 * 37 * 41 * 211 * 241 * 271 * 2161 * 9091 * 2906161","payload":"Execution time: 2.457ms"}
        