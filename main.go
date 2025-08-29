package main

import (
        "fmt"
        "strings"
        "time"
)

const (
        Reset  = "\033[0m"
        Red    = "\033[31m"
        Green  = "\033[32m"
        Yellow = "\033[33m"
        Blue   = "\033[34m"
        Purple = "\033[35m"
        Cyan   = "\033[36m"
        White  = "\033[37m"
        Bold   = "\033[1m"
)

func decoratedBorder(width int, char string, color string) {
        fmt.Print(color)
        fmt.Println(strings.Repeat(char, width))
        fmt.Print(Reset)
}

func centerText(text string, width int, color string) {
        padding := (width - len(text)) / 2
        fmt.Print(color)
        fmt.Printf("%s%s%s\n", strings.Repeat(" ", padding), text, strings.Repeat(" ", width-len(text)-padding))
        fmt.Print(Reset)
}

func animatedHello() {
        message := "Hello World"
        colors := []string{Red, Green, Yellow, Blue, Purple, Cyan}

        for i := 0; i < 5; i++ {
                for j, char := range message {
                        color := colors[j%len(colors)]
                        fmt.Print(color + Bold + string(char) + Reset)
                        time.Sleep(50 * time.Millisecond)
                }
                fmt.Println()
                time.Sleep(500 * time.Millisecond)
        }
}

func decoratedOutput() {
        width := 50

        decoratedBorder(width, "═", Cyan+Bold)
        centerText("", width, "")
        centerText("🌟 WELCOME TO VIPE CODING 🌟", width, Yellow+Bold)
        centerText("", width, "")
        decoratedBorder(width, "─", Blue)
        centerText("", width, "")

        fmt.Print(Green + Bold)
        centerText("Animated Hello World:", width, Green+Bold)
        fmt.Print(Reset)
        centerText("", width, "")

        animatedHello()

        centerText("", width, "")
        decoratedBorder(width, "═", Cyan+Bold)
        centerText("✨ Decorated with style! ✨", width, Purple+Bold)
        decoratedBorder(width, "═", Cyan+Bold)
}

func main() {
        decoratedOutput()
}
