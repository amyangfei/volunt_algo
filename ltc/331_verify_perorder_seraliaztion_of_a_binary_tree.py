# coding: utf-8

class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        deg = 1
        for x in preorder.split(","):
            deg -= 1
            if deg < 0:
                return False
            if x != "#":
                deg += 2
        return deg == 0

        '''
        stack = []
        for x in preorder.split(","):
            stack.append(x)
            while len(stack) >= 3 and stack[-3] != "#" and stack[-2:] == ["#", "#"]:
                stack = stack[:-3] + ["#"]
        return len(stack) == 1 and stack[0] == "#"
        '''
