package main

func CreateMap(s string) map[string]int {
	res := make(map[string]int)
	for i := 0; i < len(s); i++ {
		res[string(s[i])] += 1
	}
	return res
}
