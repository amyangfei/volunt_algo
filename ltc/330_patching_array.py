# coding: utf-8

# Greedy

class Solution(object):
    def minPatches(self, nums, n):
        """
        :type nums: List[int]
        :type n: int
        :rtype: int
        """
        miss, patch, idx = 1, 0, 0
        while miss <= n:
            if idx < len(nums) and nums[idx] <= miss:
                miss += nums[idx]
                idx += 1
            else:
                miss += miss
                patch += 1
        return patch

