class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        tmp = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

        decode_data = set()
        for word in words:
            decode_data.add(''.join(tmp[ord(i)-97] for i in word))
        return len(decode_data)
