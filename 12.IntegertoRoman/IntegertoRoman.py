class Solution:
    def intToRoman(self, num: int) -> str:
        roman = [{
            1: 'I',
            4: 'IV',
            5: 'V',
            9: 'IX',
        },{
            1: 'X',
            4: 'XL',
            5: 'L',
            9: 'XC',
        },{
            1: 'C',
            4: 'CD',
            5: 'D',
            9: 'CM',
        },{
            1: 'M',
        }]
        result = ''
        for i, n in enumerate(str(num)[::-1]):
            n = int(n)
            if roman[i].get(n):
                result = roman[i].get(n) + result
            elif n >= 5:
                n-=5
                result = roman[i][5] + roman[i][1]*n + result
            else:
                result = roman[i][1]*n + result
        return result
