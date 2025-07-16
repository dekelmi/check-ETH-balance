Перед запуском вставьте url своего приложения в .env файл в разделе INFURA_RPC_URL
(Создать можно на https://developer.metamask.io)

Запуск в 2 команды
1. # docker build -t eth-check .
2. # docker run -p 5090:5090 -it eth-check

Далее просто вводите адрес ETH кошелька и получаете баланс в USD с учетом самого ETH + WETH