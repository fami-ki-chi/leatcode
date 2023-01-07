func longestCommonPrefix(strs []string) string {

	pfx_len := 0
	max_pfx_len := 0

	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	sort.Strings(strs)
	first_str := strs[0]
	last_str := strs[len(strs)-1]

	if len(first_str) >= len(last_str) {
		max_pfx_len = len(last_str)
	} else {
		max_pfx_len = len(first_str)
	}

	for i := 0; i < max_pfx_len; i++ {
		if first_str[i] == last_str[i] {
			pfx_len = i + 1
		} else {
			break
		}
	}

	if pfx_len == 0 {
		return ""
	} else {
		return first_str[:(pfx_len)]
	}
}