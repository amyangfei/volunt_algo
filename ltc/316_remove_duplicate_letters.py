# coding: utf-8

import string

class Solution(object):
    def removeDuplicateLetters(self, s):
        """
        :type s: str
        :rtype: str
        """
        visit = {c: False for c in string.ascii_lowercase}
        count = {c: 0 for c in string.ascii_lowercase}

        for c in s:
            count[c] += 1

        result = []
        for c in s:
            count[c] -= 1
            if visit[c]:
                continue
            while result and count[result[-1]] > 0 and c < result[-1]:
                visit[result.pop()] = False
            result.append(c)
            visit[c] = True
        return ''.join(result)
