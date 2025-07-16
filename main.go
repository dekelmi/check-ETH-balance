package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"eth-check/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// Запрашивает адрес ETH кошелька и выводит суммарный баланс (ETH + WETH) в USD, используя данные Chainlink
func main() {
	fmt.Println("Укажите адресс ETH кошелька")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()

	addressStr := strings.TrimSpace(reader.Text())
	if addressStr == "" {
		log.Fatal("ETH адресс не может быть пустым")
	}
	if !common.IsHexAddress(addressStr) {
		log.Fatal("введенный ETH адрес некорректен")
	}

	// Переводим строку адреса в байты ее значений (требование go-ethereum)
	address := common.HexToAddress(addressStr)
	if address == (common.Address{}) {
		log.Fatal("адрес не может состоять только из нулей")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("не удалось загрузить данные из переменной окружения /.env файла")
	}

	infuraUrl := os.Getenv("INFURA_RPC_URL")
	if infuraUrl == "" {
		log.Fatal("INFURA_RPC_URL не найден в env, проверьте его наличие в вашем окружении")
	}
	contractAddressEth := os.Getenv("CHAINLINK_CONTRACT")
	if contractAddressEth == "" {
		log.Fatal("CHAINLINK_CONTRACT не найден в env, проверьте его наличие в вашем окружении")
	}
	strWethAddress := os.Getenv("WETH_CONTRACT")
	if strWethAddress == "" {
		log.Fatal("WETH_CONTRACT не найден в env, проверьте его наличие в вашем окружении")
	}

	if !common.IsHexAddress(strWethAddress) {
		log.Fatal("введенный адрес WETH контракта некорректен")
	}
	contractAddressWeth := common.HexToAddress(strWethAddress)

	// Подключение к узлу ETH
	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		log.Fatalf("не удалось подключиться к узлу ETH: %v", err)
	}
	defer client.Close()

	usdBalance, err := utils.GetUsdBalance(client, address, contractAddressEth, contractAddressWeth)
	if err != nil {
		log.Fatalf("ошибка при подсчете ETH в USD: %v", err)
	}

	fmt.Printf("ETH to USD -> %.2f\n", usdBalance)
}
