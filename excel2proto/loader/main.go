package main

import (
	"fmt"
	"os"

	pb "excel2proto/xlsp" // 路径与 go.mod 同名

	"google.golang.org/protobuf/proto"
)

func main1() {
	raw, err := os.ReadFile("../output/activity_weight.bytes")
	if err != nil {
		panic(err)
	}

	var list pb.ActivityWeightList
	if err := proto.Unmarshal(raw, &list); err != nil {
		panic(err)
	}
	fmt.Println(list.GetItems())

	m := make(map[uint32]uint32)
	for _, v := range list.Items {
		m[v.Id] = v.Weight
	}

	fmt.Println("反序列化成功，内容：", m)
}

var ans = make([][]int, 0)

func subsets(nums []int) [][]int {
	var curNum = []int{}
	dfs(0, nums, curNum)
	return ans
}

func dfs(cur int, nums []int, curNum []int) {

	ans = append(ans, curNum)
	//fmt.Println("ans", ans)

	for i := cur; i < len(nums); i++ {
		//fmt.Println("curNum", curNum)
		curNum = append(curNum, nums[i])
		dfs(i+1, nums, curNum)
		curNum = curNum[:len(curNum)-1]
	}
}

func main() {
	result := subsets([]int{1, 2, 3})
	fmt.Println("所有子集：", result)

}
