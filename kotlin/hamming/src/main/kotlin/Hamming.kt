object Hamming {

    fun compute(leftStrand: String, rightStrand: String): Int {
        if (leftStrand.length != rightStrand.length) {
          throw IllegalArgumentException("left and right strands must be of equal length")
        }

        val leftChars = leftStrand.toList()
        val rightChars = rightStrand.toList()
        var distance = 0

        leftChars.forEachIndexed { index, element ->
          if (element != rightChars[index]) {
            distance++
          }
        }

        return distance
    }
}
