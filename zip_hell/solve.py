import os

for i in range(1, 10000):
    if os.path.isfile('flag.txt'):
        print("Файл flag.txt найден!")
        break
    else:
        os.system('unzip -o "flag.zip"')