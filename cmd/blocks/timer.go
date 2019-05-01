package blocks

import (
    "time"
)


func Timer (delay_input string) {
    // Set timer
    delay, _:= time.ParseDuration(delay_input)
    timerWork := time.NewTimer(delay)
    // Launch timer
    <-timerWork.C
}
