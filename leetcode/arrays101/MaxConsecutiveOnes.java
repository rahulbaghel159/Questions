package leetcode.arrays101;

import java.lang.Math;

class MaxConsecutiveOnes{
    public int findMaxConsecutiveOnes(int[] nums) {
        if (nums.length == 0){
            return 0;
        }

        int maxOnes = 0;
        int currOnes = 0;
        int prevNum = 0;

        for (int i=0;i<nums.length;i++){
            if (prevNum == 1 && nums[i]==1){
                currOnes++;
            }else if (prevNum == 0 && nums[i]==1){
                currOnes = 1;
            }

            maxOnes = Math.max(maxOnes, currOnes);
            prevNum = nums[i];
        }

        return maxOnes;
    }
}