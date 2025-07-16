package utils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"eth-check/aggregator"
	"eth-check/weth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getEthBalance(client *ethclient.Client, address common.Address) (*big.Float, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Получение "сырого" баланса кошелька (в формате 1 000 000 000 000 000 000 для 1 ETH токена)
	rawBalance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения баланса ETH: %w", err)
	}
	// Делим баланс на квинтиллион (1e18) для получения целого количества ETH
	balance := new(big.Float).Quo(new(big.Float).SetInt(rawBalance), big.NewFloat(1e18))

	return balance, nil
}

func getWethBalance(client *ethclient.Client, address, contractAddressWeth common.Address) (*big.Float, error) {
	contract, err := weth.NewWeth(contractAddressWeth, client)
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации контракта WETH: %w", err)
	}

	// Получение "сырого" баланса Weth (в формате 1 000 000 000 000 000 000 для 1 WETH токена)
	rawBalance, err := contract.BalanceOf(nil, address)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения баланса WETH: %w", err)
	}
	// Делим баланс на квинтиллион (1e18) для получения целого количества WETH
	balance := new(big.Float).Quo(new(big.Float).SetInt(rawBalance), big.NewFloat(1e18))

	return balance, nil
}

func getTotalBalance(client *ethclient.Client, address, contractAddressWeth common.Address) (*big.Float, error) {
	ethBalance, err := getEthBalance(client, address)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения баланса ETH: %w", err)
	}
	wethBalance, err := getWethBalance(client, address, contractAddressWeth)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения баланса WETH: %w", err)
	}

	commonBalance := new(big.Float).Add(ethBalance, wethBalance)

	return commonBalance, nil
}

func getLastEthPrice(contractAddressEth string, client *ethclient.Client) (*big.Float, error) {
	// Инициализация контракта Chainlink Price Feed
	aggregator, err := aggregator.NewAggregator(common.HexToAddress(contractAddressEth), client)
	if err != nil {
		return nil, fmt.Errorf("ошибка при инициализации контракта Chainlink %w", err)
	}
	// Получаем последние данные о цене токена (ETH/USD пара) из оракула Chainlink
	roundData, err := aggregator.LatestRoundData(nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении актуальной цены ETH с Chainlink %w", err)
	}

	// Делим цену 1 ETH на 100 000 000 (1e8) т.к в roundData.Answer ≈ 200000000000
	price := new(big.Float).SetInt(roundData.Answer)
	priceUsd := new(big.Float).Quo(price, big.NewFloat(1e8))

	return priceUsd, nil
}

func GetUsdBalance(client *ethclient.Client, address common.Address, contractAddressEth string, contractAddressWeth common.Address) (*big.Float, error) {
	balance, err := getTotalBalance(client, address, contractAddressWeth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении баланса %w", err)
	}
	lastEthPrice, err := getLastEthPrice(contractAddressEth, client)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении актуальной цены ETH/USD %w", err)
	}

	// Вычисляем суммарный баланс в USD (WETH + ETH * price ETH/USD)
	ethToUsd := new(big.Float).Mul(balance, lastEthPrice)

	return ethToUsd, nil
}
