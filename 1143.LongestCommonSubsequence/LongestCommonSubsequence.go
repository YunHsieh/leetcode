package main

/**
e.g.:
abcde, ace
[
	[0, 0, 0, 0],
	[0, 0+1, 0+1, 0+1],
	[0, 0+1, 0+1, 0+1],
	[0, 0+1, 0+2, 0+2],
	[0, 0+1, 0+2, 0+2],
	[0, 0+1, 0+2, 0+3],
]
*/

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}


func LongestCommonSubsequence(text1 string, text2 string) int {
	// make the nested slice
    dp:=make([][]int, len(text1)+1)
	for j:=0;j<len(text1)+1;j++{
		dp[j] = make([]int, len(text2)+1)
	}
	// DP
	for i:=1;i<=len(text1);i++{
		for j:=1;j<=len(text2);j++{
			if text1[i-1] == text2[j-1]{
				dp[i][j] = dp[i-1][j-1] +1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j]) 
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
