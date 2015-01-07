package main

type Register interface {
	Set(rune)
	Get(rune)
}

func Set(key rune, v Register) { v.Set(key) }
func Get(key rune, v Register) { v.Get(key) }
