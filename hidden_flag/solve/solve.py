import requests
import re

headers = {}

for id in range(1, 999):
    url = f'http://127.0.0.1:8020/?id={id}'
    response = requests.get(url=url, headers=headers)
    match = re.search(r'flag\{\w+\}', response.text)
    if match:
        print(match[0])