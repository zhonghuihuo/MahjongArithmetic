/*================================================================
*   Copyright (C) 2017 刘虎. All rights reserved.
*
*   File MahjongLaiZiHuCheck.go
*   Creator ：刘虎
*   Data ：2017/05/21
*   Summary ：麻将胡牌算法
*
================================================================*/

/*
mahjongMatrix = {
	[0] = {[0]=0,[1]=0,[2]=0,[3]=0,[4]=0,[5]=0,[6]=0,[7]=0,[8]=0,[9]=0,[10]=0,[11] = 0}, --万
	[1] = {[0]=0,[1]=0,[2]=0,[3]=0,[4]=0,[5]=0,[6]=0,[7]=0,[8]=0,[9]=0,[10]=0,[11] = 0}, --筒
	[2] = {[0]=0,[1]=0,[2]=0,[3]=0,[4]=0,[5]=0,[6]=0,[7]=0,[8]=0,[9]=0,[10]=0,[11] = 0}, --条
}
*/

package main

import (
	"fmt"
	"strconv"
)

type MahjongMatrix [3][12]int

var (
	mahjongType      = []string{"万", "筒", "条"}
	mahjongValueList = []int{0x0101, 0x0201, 0x0301, 0x0401, 0x0501, 0x0601, 0x0701, 0x0801, 0x0901,
		0x1101, 0x1201, 0x1301, 0x1401, 0x1501, 0x1601, 0x1701, 0x1801, 0x1901,
		0x2101, 0x2201, 0x2301, 0x2401, 0x2501, 0x2601, 0x2701, 0x2801, 0x2901}
	allLaiZiCardList = []int{0x0301, 0x0302, 0x0303}
	checkCount       = 0
)

func main() {

	cardList := []int{0x0101, 0x0102, 0x0201, 0x0202, 0x0301, 0x0302, 0x0401, 0x0402, 0x0501, 0x0502, 0x0601, 0x0602, 0x0901, 0x0902}
	laiZiCardList := getAndRmoveLaiZiCard(cardList)
	isHu := checkLaiZiHu(cardList, len(laiZiCardList))
	fmt.Println("isHu : ", isHu)
	fmt.Println("check hu count is ", checkCount)
}

// 赖子胡牌检测(遍历)
func checkLaiZiHu(cardList []int, laiZiCount int) bool {
	for _, mahjongValue := range mahjongValueList {
		tempCardList := append(cardList, mahjongValue)
		if laiZiCount <= 1 {
			checkCount++
			mahjongMatrix := getMahjongMatrixWithCardList(tempCardList)
			printCardsInfoByMahjongMatrix(mahjongMatrix)
			isHu := checkHu(mahjongMatrix)
			if isHu {
				return isHu
			}
		} else if laiZiCount > 1 {
			isHu := checkLaiZiHu(tempCardList, laiZiCount-1)
			if isHu {
				return isHu
			}
		}
	}
	return false
}

// 去除赖子牌
func getAndRmoveLaiZiCard(cardList []int) []int {

	laiZiCardList := make([]int, 0, len(allLaiZiCardList))
	for i := 0; i < len(cardList); i++ {
		for _, laiZiValue := range allLaiZiCardList {
			if laiZiValue == cardList[i] {
				laiZiCardList = append(laiZiCardList, cardList[i])
				cardList[i] = 0
				break
			}
		}
	}
	return laiZiCardList
}

// 检测胡牌
func checkHu(mahjongMatrix MahjongMatrix) bool {

	mahjongMatrixList := getMahjongMatrixListByRemoveTwoCards(mahjongMatrix)
	for i := 0; i < len(mahjongMatrixList); i++ {
		removeThreeLinkCards(&mahjongMatrixList[i])
		removeTheSameThreeCards(&mahjongMatrixList[i])
		isHu := checkMatrixAllElemEqualZero(mahjongMatrixList[i])
		if isHu {
			return isHu
		}
	}
	return false
}

// 去除句子
func removeThreeLinkCards(mahjongMatrix *MahjongMatrix) {

	for i := 0; i < len(mahjongMatrix); i++ {
		for j := 0; j < len(mahjongMatrix[i])-2; j++ {
			if mahjongMatrix[i][j] > 0 && mahjongMatrix[i][j+1] > 0 && mahjongMatrix[i][j+2] > 0 {
				mahjongMatrix[i][j] -= 1
				mahjongMatrix[i][j+1] -= 1
				mahjongMatrix[i][j+2] -= 1
				j--
			}
		}
	}
}

// 去除克子
func removeTheSameThreeCards(mahjongMatrix *MahjongMatrix) {

	for i := 0; i < len(mahjongMatrix); i++ {
		for j := 0; j < len(mahjongMatrix[i]); j++ {
			if mahjongMatrix[i][j] >= 3 {
				mahjongMatrix[i][j] -= 3
			}
		}
	}
}

// 检测矩阵中元素是否全为0
func checkMatrixAllElemEqualZero(mahjongMatrix MahjongMatrix) bool {

	for i := 0; i < len(mahjongMatrix); i++ {
		for j := 0; j < len(mahjongMatrix[i]); j++ {
			if mahjongMatrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

// 通过去除麻将矩阵中一个将之后的麻将矩阵列表
func getMahjongMatrixListByRemoveTwoCards(mahjongMatrix MahjongMatrix) []MahjongMatrix {

	var mahjongMatrixList []MahjongMatrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 12; j++ {
			if mahjongMatrix[i][j] >= 2 {
				temp := mahjongMatrix
				fmt.Printf("temp:%p \n", &temp)
				fmt.Printf("mahjongMatrix:%p \n", &mahjongMatrix)
				temp[i][j] -= 2
				mahjongMatrixList = append(mahjongMatrixList, temp)
			}
		}
	}
	return mahjongMatrixList
}

// 通过CardList获取麻将矩阵
func getMahjongMatrixWithCardList(cardList []int) MahjongMatrix {

	var mahjongMatrix = MahjongMatrix{}
	for i := 0; i < len(cardList); i++ {
		if cardList[i] != 0 {
			cardType := cardList[i] >> 12
			cardValue := (cardList[i] >> 8) & 0xF
			mahjongMatrix[cardType][cardValue] = mahjongMatrix[cardType][cardValue] + 1
		}
	}
	return mahjongMatrix
}

// 打印麻将矩阵中牌的信息
func printCardsInfoByMahjongMatrix(mahjongMatrix MahjongMatrix) {

	cardInfo := "{"
	for i := 0; i < 3; i++ {
		for j := 0; j < 12; j++ {
			for k := 0; k < mahjongMatrix[i][j]; k++ {
				cardInfo = cardInfo + strconv.Itoa(j) + mahjongType[i] + " "
			}
		}
	}
	cardInfo = cardInfo + "}"
	fmt.Println(cardInfo)
}
