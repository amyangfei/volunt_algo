# coding: utf-8


def is_palindrome(s):
    start, end = 0, len(s)-1
    while start < end:
        if s[start] != s[end]:
            return False
        start +=1
        end -= 1
    return True


class Solution(object):
    def palindromePairs(self, words):
        """
        :type words: List[str]
        :rtype: List[List[int]]
        """
        m = {}
        s = set()
        for idx, word in enumerate(words):
            m[word] = idx
            s.add(len(word))

        result = []
        for idx, word in enumerate(words):
            rword = word[::-1]
            if rword in m and m[rword] != idx:
                result.append([idx, m[rword]])
            for sublen in filter(lambda x: x < len(word), s):
                if is_palindrome(rword[:len(rword)-sublen]) and rword[len(rword)-sublen:] in m:
                    result.append([idx, m[rword[len(rword)-sublen:]]])
                if is_palindrome(rword[sublen:]) and rword[:sublen] in m:
                    result.append([m[rword[:sublen]], idx])

        return result
