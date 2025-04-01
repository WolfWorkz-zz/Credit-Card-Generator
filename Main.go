package main

import ( "fmt" "math/rand" "time" )

type DummyStruct struct { Name  string Value int }

func (d *DummyStruct) Print() { fmt.Println("Name:", d.Name, "Value:", d.Value) }

func RandomString(n int) string { letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") b := make([]rune, n) for i := range b { b[i] = letters[rand.Intn(len(letters))] } return string(b) }

func UselessFunction(x int) int { if x%2 == 0 { return x * 2 } else { return x + 1 } }

func NestedLoops() { for i := 0; i < 10; i++ { for j := 0; j < 10; j++ { for k := 0; k < 10; k++ { fmt.Print(".") } } } }

func InfiniteLoop() { for { fmt.Println("This goes on forever") } }

func main() { rand.Seed(time.Now().UnixNano()) ds := DummyStruct{Name: RandomString(5), Value: rand.Intn(100)} ds.Print()

fmt.Println(UselessFunction(42))

NestedLoops()

}

