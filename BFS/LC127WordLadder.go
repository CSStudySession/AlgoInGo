package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]int, len(wordList))
	for _, word := range wordList {
		dict[word]++
	}

	if dict[endWord] == 0 {
		return 0
	}

	que := []string{beginWord}
	visited := make(map[string]int)
	visited[beginWord] = 1
	step := 1

	for len(que) > 0 {
		size := len(que)
		for i:= 0; i < size; i++ {
			curr := que[0]
			que = que[1:] // queue pop curr
			if curr == endWord {
				return step
			}

			for j := 0; j < len(curr); j++ {
				for k := 'a'; k <= 'z'; k++ {
					nxt := curr[:j] + string(k) + curr[j+1:]
					if dict[nxt] == 0 || visited[nxt] != 0 {
						continue
					}
					visited[nxt] = 1
					que = append(que, nxt)
				}
			}
		}
		step++
	}
	return 0
}
