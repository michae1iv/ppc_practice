import requests

url = 'http://127.0.0.1:8080/submit'
cookie = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGlja3MiOjAsInN1YiI6Im1pY2hhZWxpdiJ9.bX0VkqBZghmJ3nyNgqqPAb5rNck_-FudIBk9Ygjf6KE"

for i in range(1000):
    # Подготавливаем заголовки
    headers = {}
    if cookie:
        headers['Cookie'] = f'jwt={cookie}'

    response = requests.post(url, headers=headers)

    # Проверяем, установил ли сервер новое значение cookie
    if 'Set-Cookie' in response.headers:
        # Извлекаем новое значение cookie (предполагаем, что JWT токен установлен в cookie 'jwt')
        cookies = response.headers['Set-Cookie']
        for c in cookies.split(';'):
            if c.startswith('jwt='):
                print(c.split('=')[1])
                cookie = c.split('=')[1]  #Убираем jwt=
                print(f'Обновленное значение cookie: {cookie}')

    print(f'Запрос {i + 1} завершен, код ответа: {response.status_code}')

print('Все запросы завершены.')