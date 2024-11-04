
# Python Automation
В данной практике речь пойдет о ppc - мощном инструменте в руках пентестера, позволяющему автоматизировать рутинные действия.
Python Automation часто используется для решения задач в CTF, пентесте и тестировании функционала приложений, с помощью скриптов можно автоматизировать множество рутинных дейтсвий, что позволяет сосредоточиться на выполнении более сложных задач. Например, PPC может быть использован для написания скриптов для брутфорса паролей, фаззинга, сбора информации с сайтов и приложений, и т.д. Кроме того, PPC также может использоваться для написания скриптов, которые могут просканировать диапазон IP-адресов для поиска уязвимых узлов.


Все материалы и докер-контейнеры оставлены в данном репозитории, помимо этого написан скрипт автоматической установки docker, настройки и остановки контейнеров (запусткать с sudo).

Если вы делаете данное задание самостоятельно на своей машине, то за место <host> подставляете localhost
- Установка docker на машину:
```bash
sudo sh docker_install.sh
```
- Запуск всех контейнеров:
```bash
sudo sh build_dockers.sh
```
- Остановка всех контейнеров:
```bash
sudo sh stop_dockers.sh
```

## button_clicker:
Регестрируемся, и видим что нам нужно нажать кнопку 1000 раз, при этом после нажатия появляется таймер в 30 секунд. Если обновить страницу или найти скрипт js, в котором прописан таймер, то можно понять, что таймер работает на стороне клиента, следовательно его можно обойти если отправлять POST запросы вручную.

Также необходимо заметить, что при нажатии кнопки обновляется jwt-токен, в котором указано сколько нажатий было произведено. Таким образом после каждого нажатия необходимо обновлять установленный jwt-токен и передавать обновленный токен в каждом следующем запросе.

Готовый скрипт находится в button_clicker/solve/solve.py

Написав скрипт, получим флаг: flag{So_t1red_of_suchthing3}

## hidden_flag:
Флаг спрятан на одной из страниц, чтобы его найти необходимо просмотреть каждую страницу, однако на всех страниц, кроме одной спрятан пустой флаг, чтобы найти нужный флаг можно воспользоваться регулярками

Готовый скрипт находится в hidden_flag/solve/solve.py

После написания скрипта получаем флаг: flag{H3reIamYou__do1ngGreat}

## posts:
Флаг появляется после авторизации под конкретным пользователем, чтобы его получить, нужно собрать список авторов всех постов и забрутфорсить пароли к каждому пользователю.

Готовый скрипт находится в posts/solve/solve.py

flag{y0u_Ar3_REEEally_ppC_master_Congrats}

## zip_hell:
Флаг находится в файле flag.txt, однако данный файл заархивирован бесчисленное множество раз.
Нужно рекурсивно разархивировать каждый архив в архиве, лучше всего использовать стандартные команды операционной системы, модуль os

Готовый скрипт находится в zip_hell/solve.py

flag{Wher3--1S-The_EEEEEEnd}
## Полезные ссылки:
[Регулярные выражения](https://habr.com/ru/articles/349860/)

[Руководство по библиотеке requests](https://pythonru.com/biblioteki/kratkoe-rukovodstvo-po-biblioteke-python-requests)

[Чтение файлов в Python](https://www.w3schools.com/python/python_file_open.asp)

[Справочник по модулю os](https://docs-python.ru/standart-library/modul-os-python/)



