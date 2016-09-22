# coding: utf-8

class Solution(object):

    def get_max_sub_arr(self, nums, k):
        result, n = [], len(nums)
        for i in range(n):
            while result and n - i + len(result) > k and result[-1] < nums[i]:
                result.pop()
            if len(result) < k:
                result.append(nums[i])
        return result

    def maxNumber(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[int]
        """
        result = [0] * k
        for i in xrange(max(0, k-len(nums2)), min(k, len(nums1))+1):
            arr1 = self.get_max_sub_arr(nums1, i)
            arr2 = self.get_max_sub_arr(nums2, k-i)
            result = max(result, [max(arr1, arr2).pop(0) for _ in xrange(k)])
        return result
