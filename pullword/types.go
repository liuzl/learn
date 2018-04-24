package main

type TokenInfo struct {
	Freq  float32
	Poly  float32
	Flex  float32
	Left  map[string]float32
	Right map[string]float32
}
