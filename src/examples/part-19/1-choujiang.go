package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

type Prize struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
}

func initPrizes(str string) ([]Prize, error) {
	var prizes []Prize

	err := json.Unmarshal([]byte(str), &prizes)
	log.Infof("prizes: %v", prizes)
	if err != nil {
		log.Println("invalid json prizes")
		return nil, err
	}
	if len(prizes) == 0 {
		log.Println("prizes is empty")
		return nil, fmt.Errorf("prizes is empty")
	}
	var sum float64 = 0
	for index, v := range prizes {
		if v.Name == "" {
			return nil, fmt.Errorf("name is empty")
		}
		if v.Probability == 0 {
			return nil, fmt.Errorf("probability is empty")
		}
		v.Id = index
		sum = sum + v.Probability
	}
	if sum != 100 {
		return nil, fmt.Errorf("total probability is not equal to 100")
	}

	return prizes, nil

}

func main() {
	// 假设有以下三项奖项
	// 1. 50% 摇中一台苹果电脑
	// 2. 40% 摇中一台华为电脑
	// 3. 10% 摇中 5W现金
	// 注意：在此示例中需要保证所有 probability 之和等于 100。
	// 以下是json格式的字符串表示，可以直接使用
	prizesStr := "[{\"name\":\"苹果电脑\",\"probability\":50.0},{\"name\":\"华为电脑\",\"probability\":40.0},{\"name\":\"5W现金\",\"probability\":10.0}]"
	prizes, err := initPrizes(prizesStr)
	if err != nil {
		log.Fatal(err)
	}

	// 生成一个0到99之间的随机数
	randomNumber := rand.Intn(100)

	// 根据配置的百分比确定摇到的奖项
	sum := 0
	for _, p := range prizes {
		sum += int(p.Probability)
		if randomNumber < sum {
			fmt.Printf("摇到的奖项是：%s,号码:%d\n", p.Name, randomNumber)
			break
		}
	}
}
