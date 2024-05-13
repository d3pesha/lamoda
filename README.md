# Lamoda

Сервис предназначен для резервации товаров на складах разными пользователями.

Для запуска приложения требуется наличие git, docker и makefile. По умолчанию backend работает на 8000 порту, а postgres - на 5432 порту.

Запуск сервиса
git clone https://github.com/d3pesha/lamoda

cd lamoda

make up или docker-compose up

Поднимается postgres контейнер. 
Поднимается контейнер с миграциями, запускает и отключается. 
Поднимается backend контейнер.
Сервис готов к работе

POSTMAN коллекция с описанием методов и тестовыми запросами:
https://documenter.getpostman.com/view/30827725/2sA3JNbg32

Также коллекция лежит в папке postman.
