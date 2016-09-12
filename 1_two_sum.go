type Node struct {
    Val int
    Pos int
}

type ByVal []Node

func (tmp ByVal) Len() int {
    return len(tmp)
}

func (tmp ByVal) Swap(i,j int) {
    tmp[i],tmp[j] = tmp[j], tmp[i]
}

func (tmp ByVal) Less(i,j int) bool {
    return tmp[i].Val < tmp[j].Val
}

func twoSum(nums []int, target int) []int {
    var store []Node
    var ret []int
    for i:=0; i<len(nums); i++ {
        var tmp Node
        tmp.Val = nums[i]
        tmp.Pos = i
        store = append(store, tmp)
    }
    sort.Sort(ByVal(store))
    i:=0
    j:=len(store)-1
    for {
        if(i >= j) {
            break
        }
        if store[i].Val + store[j].Val == target {
            ret = append(ret, store[i].Pos, store[j].Pos)
            sort.Ints(ret)
            return ret
        } else if store[i].Val + store[j].Val > target {
            j--
        } else {
            i++
        }
    }
    return ret
}
