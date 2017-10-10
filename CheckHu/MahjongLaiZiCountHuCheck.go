/**
 * MahjongLaiZiCountHuCheck.go
 * Created by 刘虎 on 2017-10-03 14:40:10
 *
 * huliuworld@yahoo.com
 * Copyright © 2017 刘虎. All rights reserved.
 */

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
	allLaiZiCardList = []int{0x0301, 0x0302, 0x0303, 0x0304}
	checkCount       = 0
)

func main() {
	cardList := []int{0x0101, 0x0101, 0x0201, 0x0202, 0x0301, 0x0302, 0x0303, 0x0401, 0x0501, 0x0502, 0x0601, 0x0602, 0x0901, 0x0902}
	laiZiCard := getAndRmoveLaiZiCard(cardList)
	isHu := checkLaiZiHu(cardList, len(laiZiCard))
	fmt.Println("isHu ", isHu)
}

func checkLaiZiHu(cardList []int, laiZiCount int) bool {

	mahjongMatrix := getMahjongMatrixWithCardList(cardList)
	removeThreeLinkCards(&mahjongMatrix)
	removeTheSameThreeCards(&mahjongMatrix)
	for i := 0; i < len(mahjongMatrix); i++ {
		for j := 0; j < len(mahjongMatrix[i]); j++ {
			if mahjongMatrix[i][j] > 0 {
				tempMahjong := mahjongMatrix
				needLaiZiCount := tempMahjong[i][j] % 2
				tempMahjong[i][j] = 0
				needLaiZiCount = getNeedLaiZiCountByMahjongMatrix(tempMahjong, needLaiZiCount)
				if needLaiZiCount <= laiZiCount {
					return true
				}
			}
		}
		needLaiZiCount := getNeedLaiZiCountByMahjongMatrix(mahjongMatrix, 2)
		if needLaiZiCount <= laiZiCount {
			return true
		}
	}
	return false
}

// 计算需要赖子的数量
func getNeedLaiZiCountByMahjongMatrix(mahjongMatrix MahjongMatrix, needLaiZiCount int) int {

	minLaiZiCount := needLaiZiCount
	if !checkMatrixAllElemEqualZero(mahjongMatrix) {
		for i := 0; i < len(mahjongMatrix); i++ {
			for j := 0; j < len(mahjongMatrix[i]); j++ {
				if mahjongMatrix[i][j] <= 0 {
					continue
				}
				if mahjongMatrix[i][j+1] > 0 {
					mahjongMatrix[i][j]--
					mahjongMatrix[i][j+1]--
					j--
					minLaiZiCount++
					continue
				}
				if mahjongMatrix[i][j+2] > 0 {
					mahjongMatrix[i][j]--
					mahjongMatrix[i][j+2]--
					j--
					minLaiZiCount++
					continue
				}
				if mahjongMatrix[i][j] == 1 {
					mahjongMatrix[i][j]--
					minLaiZiCount += 2
					continue
				}
				if mahjongMatrix[i][j] == 2 {
					mahjongMatrix[i][j] -= 2
					minLaiZiCount++
				}
			}
		}
	}
	return minLaiZiCount
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

// 打印cardList信息
func printCardList(cardList []int) {

	cardInfo := "{"
	for _, card := range cardList {
		cardType := card >> 12
		cardValue := card >> 8 & 0x0F
		cardInfo += strconv.Itoa(cardValue) + mahjongType[cardType] + " "
	}
	cardInfo += "}"
	fmt.Println(cardInfo)
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
