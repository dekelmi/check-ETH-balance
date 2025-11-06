Before launching, insert your application's URL into the .env file in the INFURA_RPC_URL section
(You can create it at https://developer.metamask.io)

Launch in 2 commands1.
1. "docker build -t eth-check ."
2. "docker run -p 5090:5090 -it eth-check"
Simply enter your ETH wallet address and get your balance in USD, taking into account ETH + WETH
