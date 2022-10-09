#!/bin/bash/

fuser -k 4222/tcp &> /dev/null
fuser -k 3333/tcp &> /dev/null

sleep 1

echo "Nats-streaming"
nats-streaming-server &> /dev/null &

sleep 1

./Service &

sleep 2

echo "http://localhost:3333/"

echo
echo "Можно начать публикацию в канал"
cat ;


echo "Публикуются Json"


sleep 1

./PublisherServ -ms 100 -j ./json &> /dev/null

sleep 1

echo "Готово!"
