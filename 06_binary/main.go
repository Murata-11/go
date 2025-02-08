package main

import (
	"fmt"
	"strconv"
)

func intToBinaryArray(n int64) []string {
	// 整数を2進数文字列に変換
	binaryStr := strconv.FormatInt(n, 2)

	// 各ビットを []string に格納
	binaryArray := make([]string, len(binaryStr))
	for i, bit := range binaryStr {
		binaryArray[i] = fmt.Sprintf("[%d] %c", i, bit)
	}

	return binaryArray
}

func intToBinaryArray1Based(n int64) []string {
	// 整数を2進数文字列に変換
	binaryStr := strconv.FormatInt(n, 2)
	binaryLen := len(binaryStr)

	// インデックス 1 から使うために、+1 のスライスを確保
	binaryArray := make([]string, binaryLen+1)

	// 各ビットをスライスの 1 から格納（スライスの 0 は使わない）
	for i, bit := range binaryStr {
		binaryArray[i+1] = fmt.Sprintf("[%d] %c", i+1, bit)
	}

	return binaryArray
}

func main() {
	num := int64(10)
	binaryArray := intToBinaryArray(num)
	for i, bit := range binaryArray {
		fmt.Printf("%d: %s\n", i, bit)
	}
	binaryArray = intToBinaryArray1Based(num)
	for i, bit := range binaryArray {
		fmt.Printf("%d: %s\n", i, bit)
	}
}
