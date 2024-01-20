import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond(initDate: Any) {

    constructor(localDate: LocalDate) : this(localDate as Any)

    constructor(localDateTime: LocalDateTime) : this(localDateTime as Any)

    val date: LocalDateTime = when(initDate) {
      is LocalDate -> initDate.atStartOfDay()
      is LocalDateTime -> initDate
      else -> throw IllegalArgumentException("Unsupported type")
    }.plusSeconds(1_000_000_000)
}
