func solution(my_string string, queries [][]int) string {
    var arr = []rune(my_string)
    for _, data := range queries {
        d_len := data[1] - data[0]
        for i:=0; i<=d_len/2; i++ {
            arr[data[0] + i], arr[data[1] - i] = arr[data[1] - i], arr[data[0] + i] 
        }
    }
    return string(arr)
}