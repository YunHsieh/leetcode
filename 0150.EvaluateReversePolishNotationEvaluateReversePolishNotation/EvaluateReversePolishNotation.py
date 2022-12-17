class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        if len(tokens) == 1:
            return int(tokens[0])
        nums = []
        for token in tokens:
            if token not in ['+', '-', '*', '/']:
                nums.append(token)
                continue
            nums.append(eval(f'int({nums.pop(-2)} {token} {nums.pop()})'))
        return nums[0]
