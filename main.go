package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"

	"github.com/juicemia/gophercon-hackday/api"
)

var (
	board *raspi.Adaptor

	gp  *i2c.GrovePiDriver
	lcd *i2c.GroveLcdDriver
)

func main() {
	log.Info("booting...")
	board = raspi.NewAdaptor()

	gp = i2c.NewGrovePiDriver(board)
	lcd = i2c.NewGroveLcdDriver(board)

	msgs := make(chan string)

	work := func() {
		Message("Starting...")
		go func() {
			Message("system ready")

			for msg := range msgs {
				log.Debug("send it")
				Message(msg)
			}
		}()

	}

	log.Info("booting api server...")
	go api.NewServer(":9090", msgs)

	log.Info("starting robot...")
	robot := gobot.NewRobot("sensors",
		[]gobot.Connection{board},
		[]gobot.Device{gp, lcd},
		work,
	)

	robot.Start()
}
func Message(msg string) {
	fmt.Println(msg)

	lcd.Clear()
	lcd.Home()
	lcd.Write(msg)
}
