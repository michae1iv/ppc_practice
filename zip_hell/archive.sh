#!/bin/bash

zip flag.zip flag.txt

max=1000
for i in `seq 1 $max`
do
    zip "flag_$i.zip" "flag.zip"
    mv "flag_$i.zip" "flag.zip"
done

echo "Архивирование завершено!"