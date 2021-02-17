package main

func getShortestDist(beginWord string, endWord string,
					dict map[string]bool, wordToDist map[string]int) {
	wordToDist[beginWord] = 1
	que := []string{beginWord}
	step := 2
	for len(que) != 0 {
		sz := len(que)
		for ; sz > 0; sz-- {
			curr := que[0]
			que = que[1:]
			if curr == endWord {return}
			for i := 0; i < len(curr); i++ {
				for c := 'a'; c <= 'z'; c++ {
					nxt := curr[:i] + string(c) + curr[i+1:]
					if !dict[nxt] || wordToDist[nxt] != 0 {continue}
					wordToDist[nxt] = step
					que = append(que, nxt)
				}
			}
		}
		step++
	}
}

func getAllMinPath(res *[][]string, tmp []string, curr string,
					beginWord string, dist map[string]int, depth int) {
	if curr == beginWord {
		rev := make([]string, len(tmp))
		for k := len(tmp) - 1; k >= 0; k-- {
			rev[len(rev) - k - 1]  = tmp[k]
		}
		*res = append(*res, rev)
		return
	}

	if depth <= 0 {return}

	for i := 0; i < len(curr); i++ {
		for c := 'a'; c <= 'z'; c++ {
			nxt := curr[0:i] + string(c) + curr[i+1:]
			if nxt == curr || dist[nxt] != depth {continue}
			delete(dist, nxt)
			tmp = append(tmp, nxt)
			getAllMinPath(res, tmp, nxt, beginWord, dist, depth - 1)
			dist[nxt] = depth
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func insertFront(tmp []string, curr string) []string {
	nxtmp := append(tmp, "")
	copy(nxtmp[1:], nxtmp[0:])
	nxtmp[0] = curr
	return nxtmp
}


func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := [][]string{}
	dict := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		dict[word] = true
	}

	if wordList == nil || len(wordList) == 0 || !dict[endWord] {
		return res
	}

	wordToDist := make(map[string]int, len(wordList) + 1)
	getShortestDist(beginWord, endWord, dict, wordToDist)
	tmp := []string {endWord}
	getAllMinPath(&res, tmp, endWord, beginWord, wordToDist, wordToDist[endWord] - 1)

	return res
}

/*
func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string {"hot","dot","dog","lot","log","cog"}
	res := findLadders(beginWord, endWord, wordList)
	fmt.Println(res)
}
 */
