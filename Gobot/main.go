package main

import (
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)

func main() {
        firmataAdaptor := firmata.NewTCPAdaptor("172.20.10.13:3030")
        led := gpio.NewLedDriver(firmataAdaptor, "2")

        work := func() {
                gobot.Every(1*time.Second, func() {
                        led.Toggle()
                })
        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led},
                work,
        )

        robot.Start()
}

