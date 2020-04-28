package leetcodego


func twoSum(nums []int, target int) []int {
    twoSumMap := make(map[int]int, len(nums))
    for i, num := range nums{
        //fmt.Println(twoSumMap)
		val,ok:=twoSumMap[target-num]
		if ok{
			return []int{val,i}
		}else{
			twoSumMap[num]=i
			continue
		}
	}
    return []int{}
}
