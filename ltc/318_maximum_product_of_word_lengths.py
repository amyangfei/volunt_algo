# coding: utf-8

class Solution(object):
    def maxProduct(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        n = len(words)
        bits = [0] * n
        orda = ord('a')
        for i, word in enumerate(words):
            for c in word:
                bits[i] |= 1 << ord(c) - orda

        result = 0
        for i in range(n):
            for j in range(i+1, n):
                if not (bits[i] & bits[j]):
                    result = max(result, len(words[i]) * len(words[j]))
        return result
