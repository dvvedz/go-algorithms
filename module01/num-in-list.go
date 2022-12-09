package module01

func NumInList(list []int, num int) bool {

	for _, l := range list {
		if num == l {
			// fmt.Println(l)
			return true
		}
	}
	return false
}
