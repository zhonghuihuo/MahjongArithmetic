--[[--麻将矩阵 
local mahjongMatrix = {
	[0] = {[-1]=0,[0]=0,[1]=0,[2]=0,[3]=0,[4]=0,[5]=0,[6]=0,[7]=0,[8]=0,[9]=0,[10]=0,[11] = 0}, --万
	[1] = {[-1]=0,[0]=0,[1]=0,[2]=0,[3]=0,[4]=0,[5]=0,[6]=0,[7]=0,[8]=0,[9]=0,[10]=0,[11] = 0}, --筒
	[2] = {[-1]=0,[0]=0,[1]=0,[2]=0,[3]=0,[4]=0,[5]=0,[6]=0,[7]=0,[8]=0,[9]=0,[10]=0,[11] = 0}, --条
} ]] 

local mahjongType = {
	[0] = "万",
	[1] = "筒",
	[2] = "条",
}

-- 深拷贝
function deepCopy(sourceData)

	if type(sourceData) == "table" then
		local temp = {}
		for key, value in pairs(sourceData) do
			temp[key] = deepCopy(value)
		end
		return temp
	end
	return sourceData
end

-- 初始化一个麻将矩阵
function initMahjongMatrix()

	local mahjongMatrix = {}
	for i = 0, 3 do
		mahjongMatrix[i] = {}
		for j = -1, 11 do
			mahjongMatrix[i][j] = 0
		end
	end	
	return mahjongMatrix;
end

-- 打印麻将矩阵
function dumpMahjongMatrix(mahjongMatrix)

	local dumpInfo = "{ "
	for cardType, mahjongList in pairs(mahjongMatrix) do
		for mahjongValue, count in pairs(mahjongList) do
			for i = 1, count do
				dumpInfo = dumpInfo .. mahjongValue .. mahjongType[cardType] .. " "
			end
		end
	end
	dumpInfo = dumpInfo .. "}"
	print(dumpInfo)
end

-- 初始化一个麻将矩阵
function cardListConvertToMatrix(cardList)

	local mahjongMatrix = initMahjongMatrix()
	for _, card in pairs(cardList) do
		local cardType = card >> 12
		local cardValue = (card >> 8) & 0x0F
		mahjongMatrix[cardType][cardValue] = mahjongMatrix[cardType][cardValue] + 1
	end
	dumpMahjongMatrix(mahjongMatrix)
	return mahjongMatrix
end

-- 通过去除麻将矩阵中一个将之后的麻将矩阵列表
function getMahjongMatrixListByRemoveTwoCards(mahjongMatrix)

	local mahjongMatrixList = {}
	for cardType, mahjongList in pairs(mahjongMatrix) do
		for mahjongValue, count in pairs(mahjongList) do
			if count >= 2 then
				local temp = deepCopy(mahjongMatrix)
				temp[cardType][mahjongValue] = temp[cardType][mahjongValue] - 2
				table.insert(mahjongMatrixList, temp);
			end
		end
	end
	return mahjongMatrixList
end

-- 移除麻将矩阵中的句子
function removeThreeLinkCards(mahjongMatrix)

	for cardType, mahjongList in pairs(mahjongMatrix) do
		for mahjongValue, count in pairs(mahjongList) do
			for i=1, count do
				local mahjongValuePlusOneCount = mahjongList[mahjongValue+1]
				local mahjongValuePlusTwoCount = mahjongList[mahjongValue+2]
				if count > 0 and mahjongValuePlusOneCount > 0 and mahjongValuePlusTwoCount > 0 then
					mahjongList[mahjongValue] = mahjongList[mahjongValue] - 1
					mahjongList[mahjongValue+1] = mahjongList[mahjongValue+1] - 1
					mahjongList[mahjongValue+2] = mahjongList[mahjongValue+2] - 1
				end
			end
		end
	end
end

-- 检测克子
function removeTheSameThreeCards(mahjongMatrix)
	for cardType, mahjongList in pairs(mahjongMatrix) do
		for mahjongValue, count in pairs(mahjongList) do
			if count >= 3 then
				mahjongList[mahjongValue] = mahjongList[mahjongValue] - 3
			end
		end
	end
end

-- 检测矩阵中元素是否全部为0
function checkMatrixAllElemEqualZero(mahjongMatrix)

	for cardType, mahjongList in pairs(mahjongMatrix) do
		for mahjongValue, count in pairs(mahjongList) do
			if count ~= 0 then
				return false
			end
		end
	end
	return true
end

-- 检测是否胡牌
function checkHu(cardList)

	local mahjongMatrix = cardListConvertToMatrix(cardList)
	local mahjongCardList = getMahjongMatrixListByRemoveTwoCards(mahjongMatrix)
	for _, matrix in ipairs(mahjongCardList) do
		removeThreeLinkCards(matrix)
		removeTheSameThreeCards(matrix)
		local result = checkMatrixAllElemEqualZero(matrix)
		if result == true then
			return true
		end
	end
	return false
end

local cardList = {0x0101,0x0201,0x0202,0x0301,0x0302,0x0303,0x0401,0x0402,0x0501,0x1601,0x1701,0x1801,0x0901,0x0902}

local isHu = checkHu(cardList)
print("isHu ", isHu);













