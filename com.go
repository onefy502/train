package main

import "github.com/fatih/color"

func printer(s string) {
	ch := make(chan int, 1)
	select {
		case ch <- 1: color.Cyan(s)
		case ch <- 1: color.Red(s)
		case ch <- 1: color.Yellow(s)
		case ch <- 1: color.Green(s)
		case ch <- 1: color.Blue(s)
		case ch <- 1: color.HiRed(s)
	}
	<-ch
	close(ch)
	/**
	rand.Seed(time.Now().UnixNano())
	rdm := rand.Intn(60)
	switch {
		case rdm <= 10:color.Cyan(s)
		case rdm <= 20:color.Red(s)
		case rdm <= 30:color.Yellow(s)
		case rdm <= 40:color.Green(s)
		case rdm <= 50:color.Blue(s)
		case rdm <= 60:color.HiRed(s)
	} */

}
