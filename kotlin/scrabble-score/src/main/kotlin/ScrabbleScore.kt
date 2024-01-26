object ScrabbleScore {
    val letterMapping = mapOf<Char, Int>(
      'A' to 1,
      'B' to 3,
      'C' to 3,
      'D' to 2,
      'E' to 1,
      'F' to 4,
      'G' to 2,
      'H' to 4,
      'I' to 1,
      'J' to 8,
      'K' to 5,
      'L' to 1,
      'M' to 3,
      'N' to 1,
      'O' to 1,
      'P' to 3,
      'Q' to 10,
      'R' to 1,
      'S' to 1,
      'T' to 1,
      'U' to 1,
      'V' to 4,
      'Y' to 4,
      'W' to 4,
      'X' to 8,
      'Z' to 10,
    )

    fun scoreLetter(c: Char): Int {
        return letterMapping[c] ?: 0
    }

    fun scoreWord(word: String): Int {
      return word.uppercase().toList().map(::scoreLetter).sum()
    }
}
