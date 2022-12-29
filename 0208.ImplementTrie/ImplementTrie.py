class Trie:

    def __init__(self):
        self.words = {}
        self.trie = {}

    def insert(self, word: str) -> None:
        self.words[word] = 'x'
        tmp = self.trie
        for w in word:
            tmp[w] = tmp.get(w, {})
            tmp = tmp[w]

    def search(self, word: str) -> bool:
        return bool(self.words.get(word))

    def startsWith(self, prefix: str) -> bool:
        tmp = self.trie
        for w in prefix:
            tmp = tmp.get(w, None)
            if not isinstance(tmp, dict):
                return False
        return True
