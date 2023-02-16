package main

import "fmt"

func main() {
	var n, x, k int
	var channelList []string
	fmt.Scanf("%d %d %d\n", &n, &x, &k)
	for i := 0; i < n; i++ {
		var tmpStr string
		fmt.Scanf("%s\n", &tmpStr)
		channelList = append(channelList, tmpStr)
	}
	fmt.Println(finalChannelCalculator(n, x, k, channelList))
}

func finalChannelCalculator(n, x, k int, channelList []string) string {
	if x+k <= n {
		return channelList[x+k-1]
	} else {
		return channelList[((k+x-1)%n+n)%n]
	}
}
