/**
 * @Author: Hongker
 * @Description:
 * @File:  weight
 * @Version: 1.0.0
 * @Date: 2020/6/14 15:08
 */

package balance

import (
	"github.com/ebar-go/ego/utils/array"
)

const(
	InvalidIndex = -1
)


// WeightBalance 权重轮训算法
type WeightBalance struct {
	count         int // 节点数
	maxWeight     int // 最大权值
	lastNodeIndex int // 上一次的节点
	currentWeight int // 当前权值
	gcd           int // 公约数
	weights []int
}


// Init
func (loader *WeightBalance) Reload(weights []int) {
	arr := array.Int(weights)
	loader.lastNodeIndex = InvalidIndex
	loader.currentWeight = 0
	loader.count = arr.Length()


	if loader.count > 1 {
		// 计算最大公约数
		loader.gcd = NGcd(weights, loader.count)
	} else {
		loader.gcd = 1
	}
	loader.maxWeight = arr.Max()
	loader.weights = weights

}

// RandomIndex
func (loader *WeightBalance) RandomIndex() int {
	if loader.count == 0 {
		return InvalidIndex
	}

	// 使用权重轮询调度算法计算当前节点
	for {
		// 计算权值
		loader.lastNodeIndex = (loader.lastNodeIndex + 1) % loader.count
		if loader.lastNodeIndex == 0 {
			loader.currentWeight = loader.currentWeight - loader.gcd
			if loader.currentWeight <= 0 {
				loader.currentWeight = loader.maxWeight
				if loader.currentWeight == 0 {
					return InvalidIndex
				}
			}
		}

		// 判断权值，如果当前节点的权重大于当前权值，则返回该节点
		if loader.weights[loader.lastNodeIndex] >= loader.currentWeight {
			return loader.lastNodeIndex
		}
	}
}


func Gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}else {
		return Gcd(b, a%b)
	}


}

func NGcd(numbers []int, n int) int {
	if n == 1 {
		return numbers[n-1]
	}

	return Gcd(numbers[n-1], NGcd(numbers, n-1))
}
