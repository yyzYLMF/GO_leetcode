func singleNumber(nums []int) int {
    ret := 0
    for i:=0; i<len(nums); i++ {
        ret = ret^nums[i]
    }
    return ret
}
