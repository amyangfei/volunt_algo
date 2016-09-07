# coding: utf-8

import sys

class Solution(object):
    def increasingTriplet(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        low, high = sys.maxint, sys.maxint
        for num in nums:
            if num <= low:
                low = num
            elif num <= high:
                high = num
            else:
                return True
        return False
