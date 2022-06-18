package sum

func main() {

}

func Sum_array(number [5]int) int {
	ret := 0
	for _, value := range number {
		ret += value
	}
	return ret
}

func Sum_slice(number []int) int {
	ret := 0
	for _, value := range number {
		ret += value
	}
	return ret
}
