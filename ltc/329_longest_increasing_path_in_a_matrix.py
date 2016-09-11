# coding: utf-8

class Solution(object):
    def longestIncreasingPath(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        """
        if not matrix:
            return 0
        m, n = len(matrix), len(matrix[0])
        dp = [[0] * n for _ in range(m)]
        return max([self.dfs(matrix, m, n, i, j, dp) for i in range(m) for j in range(n)])

    def dfs(self, matrix, m, n, x, y, dp):
        if dp[x][y] > 0:
            return dp[x][y]
        for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            xx, yy = x + dx, y + dy
            if xx >= 0 and xx < m and yy >= 0 and yy < n and matrix[x][y] < matrix[xx][yy]:
                dp[x][y] = max(dp[x][y], self.dfs(matrix, m, n, xx, yy, dp))
        dp[x][y] += 1
        return dp[x][y]
