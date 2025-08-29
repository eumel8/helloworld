package main

import (
        "bytes"
        "io"
        "os"
        "strings"
        "testing"
        "time"
)

func TestDecoratedBorder(t *testing.T) {
        tests := []struct {
                name  string
                width int
                char  string
                color string
        }{
                {"basic border", 10, "═", Cyan},
                {"different char", 5, "-", Red},
                {"zero width", 0, "*", Green},
                {"single char", 1, "#", Blue},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        oldStdout := os.Stdout
                        r, w, _ := os.Pipe()
                        os.Stdout = w

                        decoratedBorder(tt.width, tt.char, tt.color)

                        w.Close()
                        os.Stdout = oldStdout

                        var buf bytes.Buffer
                        io.Copy(&buf, r)
                        output := buf.String()

                        expectedChars := strings.Repeat(tt.char, tt.width)
                        if !strings.Contains(output, expectedChars) {
                                t.Errorf("Expected output to contain %q, got %q", expectedChars, output)
                        }
                        if !strings.Contains(output, tt.color) {
                                t.Errorf("Expected output to contain color %q, got %q", tt.color, output)
                        }
                        if !strings.Contains(output, Reset) {
                                t.Errorf("Expected output to contain Reset, got %q", output)
                        }
                })
        }
}

func TestCenterText(t *testing.T) {
        tests := []struct {
                name  string
                text  string
                width int
                color string
        }{
                {"basic centering", "Hello", 10, Red},
                {"exact width", "Test", 4, Green},
                {"empty text", "", 5, Blue},
                {"text longer than width", "VeryLongText", 15, Yellow},
                {"odd width", "Hi", 7, Purple},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        oldStdout := os.Stdout
                        r, w, _ := os.Pipe()
                        os.Stdout = w

                        centerText(tt.text, tt.width, tt.color)

                        w.Close()
                        os.Stdout = oldStdout

                        var buf bytes.Buffer
                        io.Copy(&buf, r)
                        output := buf.String()

                        if !strings.Contains(output, tt.text) {
                                t.Errorf("Expected output to contain %q, got %q", tt.text, output)
                        }
                        if tt.color != "" && !strings.Contains(output, tt.color) {
                                t.Errorf("Expected output to contain color %q, got %q", tt.color, output)
                        }
                        if tt.color != "" && !strings.Contains(output, Reset) {
                                t.Errorf("Expected output to contain Reset, got %q", output)
                        }
                })
        }
}

func TestAnimatedHello(t *testing.T) {
        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        done := make(chan bool)
        go func() {
                animatedHello()
                done <- true
        }()

        select {
        case <-done:
        case <-time.After(10 * time.Second):
                t.Fatal("animatedHello() took too long to complete")
        }

        w.Close()
        os.Stdout = oldStdout

        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()

        expectedMessage := "Hello World"
        if !strings.Contains(output, "H") && !strings.Contains(output, "e") && !strings.Contains(output, "l") {
                t.Errorf("Expected output to contain characters from %q, got %q", expectedMessage, output)
        }

        colors := []string{Red, Green, Yellow, Blue, Purple, Cyan}
        foundColor := false
        for _, color := range colors {
                if strings.Contains(output, color) {
                        foundColor = true
                        break
                }
        }
        if !foundColor {
                t.Error("Expected output to contain at least one color")
        }

        if !strings.Contains(output, Bold) {
                t.Error("Expected output to contain Bold formatting")
        }

        if !strings.Contains(output, Reset) {
                t.Error("Expected output to contain Reset")
        }
}

func TestDecoratedOutput(t *testing.T) {
        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        done := make(chan bool)
        go func() {
                decoratedOutput()
                done <- true
        }()

        select {
        case <-done:
        case <-time.After(15 * time.Second):
                t.Fatal("decoratedOutput() took too long to complete")
        }

        w.Close()
        os.Stdout = oldStdout

        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()

        expectedTexts := []string{
                "WELCOME TO VIPE CODING",
                "Animated Hello World:",
                "Decorated with style!",
        }

        for _, text := range expectedTexts {
                if !strings.Contains(output, text) {
                        t.Errorf("Expected output to contain %q, got output of length %d", text, len(output))
                }
        }

        if !strings.Contains(output, "═") {
                t.Error("Expected output to contain border character ═")
        }

        if !strings.Contains(output, "─") {
                t.Error("Expected output to contain border character ─")
        }

        colors := []string{Cyan, Yellow, Green, Blue, Purple}
        for _, color := range colors {
                if !strings.Contains(output, color) {
                        t.Errorf("Expected output to contain color %q", color)
                }
        }
}

func TestMain(t *testing.T) {
        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        done := make(chan bool)
        go func() {
                main()
                done <- true
        }()

        select {
        case <-done:
        case <-time.After(15 * time.Second):
                t.Fatal("main() took too long to complete")
        }

        w.Close()
        os.Stdout = oldStdout

        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()

        if len(output) == 0 {
                t.Error("Expected main() to produce output")
        }

        if !strings.Contains(output, "WELCOME TO VIPE CODING") {
                t.Error("Expected main() to call decoratedOutput()")
        }
}

func TestColorConstants(t *testing.T) {
        constants := map[string]string{
                "Reset":  Reset,
                "Red":    Red,
                "Green":  Green,
                "Yellow": Yellow,
                "Blue":   Blue,
                "Purple": Purple,
                "Cyan":   Cyan,
                "White":  White,
                "Bold":   Bold,
        }

        for name, value := range constants {
                if value == "" {
                        t.Errorf("Expected %s constant to be non-empty", name)
                }
                if !strings.HasPrefix(value, "\033[") {
                        t.Errorf("Expected %s constant to be ANSI escape sequence, got %q", name, value)
                }
        }
}eumel@ollama:~/helloworld$
eumel@ollama:~/helloworld$ cat main_test.go
package main

import (
        "bytes"
        "io"
        "os"
        "strings"
        "testing"
        "time"
)

func TestDecoratedBorder(t *testing.T) {
        tests := []struct {
                name  string
                width int
                char  string
                color string
        }{
                {"basic border", 10, "═", Cyan},
                {"different char", 5, "-", Red},
                {"zero width", 0, "*", Green},
                {"single char", 1, "#", Blue},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        oldStdout := os.Stdout
                        r, w, _ := os.Pipe()
                        os.Stdout = w

                        decoratedBorder(tt.width, tt.char, tt.color)

                        w.Close()
                        os.Stdout = oldStdout

                        var buf bytes.Buffer
                        io.Copy(&buf, r)
                        output := buf.String()

                        expectedChars := strings.Repeat(tt.char, tt.width)
                        if !strings.Contains(output, expectedChars) {
                                t.Errorf("Expected output to contain %q, got %q", expectedChars, output)
                        }
                        if !strings.Contains(output, tt.color) {
                                t.Errorf("Expected output to contain color %q, got %q", tt.color, output)
                        }
                        if !strings.Contains(output, Reset) {
                                t.Errorf("Expected output to contain Reset, got %q", output)
                        }
                })
        }
}

func TestCenterText(t *testing.T) {
        tests := []struct {
                name  string
                text  string
                width int
                color string
        }{
                {"basic centering", "Hello", 10, Red},
                {"exact width", "Test", 4, Green},
                {"empty text", "", 5, Blue},
                {"text longer than width", "VeryLongText", 15, Yellow},
                {"odd width", "Hi", 7, Purple},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        oldStdout := os.Stdout
                        r, w, _ := os.Pipe()
                        os.Stdout = w

                        centerText(tt.text, tt.width, tt.color)

                        w.Close()
                        os.Stdout = oldStdout

                        var buf bytes.Buffer
                        io.Copy(&buf, r)
                        output := buf.String()

                        if !strings.Contains(output, tt.text) {
                                t.Errorf("Expected output to contain %q, got %q", tt.text, output)
                        }
                        if tt.color != "" && !strings.Contains(output, tt.color) {
                                t.Errorf("Expected output to contain color %q, got %q", tt.color, output)
                        }
                        if tt.color != "" && !strings.Contains(output, Reset) {
                                t.Errorf("Expected output to contain Reset, got %q", output)
                        }
                })
        }
}

func TestAnimatedHello(t *testing.T) {
        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        done := make(chan bool)
        go func() {
                animatedHello()
                done <- true
        }()

        select {
        case <-done:
        case <-time.After(10 * time.Second):
                t.Fatal("animatedHello() took too long to complete")
        }

        w.Close()
        os.Stdout = oldStdout

        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()

        expectedMessage := "Hello World"
        if !strings.Contains(output, "H") && !strings.Contains(output, "e") && !strings.Contains(output, "l") {
                t.Errorf("Expected output to contain characters from %q, got %q", expectedMessage, output)
        }

        colors := []string{Red, Green, Yellow, Blue, Purple, Cyan}
        foundColor := false
        for _, color := range colors {
                if strings.Contains(output, color) {
                        foundColor = true
                        break
                }
        }
        if !foundColor {
                t.Error("Expected output to contain at least one color")
        }

        if !strings.Contains(output, Bold) {
                t.Error("Expected output to contain Bold formatting")
        }

        if !strings.Contains(output, Reset) {
                t.Error("Expected output to contain Reset")
        }
}

func TestDecoratedOutput(t *testing.T) {
        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        done := make(chan bool)
        go func() {
                decoratedOutput()
                done <- true
        }()

        select {
        case <-done:
        case <-time.After(15 * time.Second):
                t.Fatal("decoratedOutput() took too long to complete")
        }

        w.Close()
        os.Stdout = oldStdout

        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()

        expectedTexts := []string{
                "WELCOME TO VIPE CODING",
                "Animated Hello World:",
                "Decorated with style!",
        }

        for _, text := range expectedTexts {
                if !strings.Contains(output, text) {
                        t.Errorf("Expected output to contain %q, got output of length %d", text, len(output))
                }
        }

        if !strings.Contains(output, "═") {
                t.Error("Expected output to contain border character ═")
        }

        if !strings.Contains(output, "─") {
                t.Error("Expected output to contain border character ─")
        }

        colors := []string{Cyan, Yellow, Green, Blue, Purple}
        for _, color := range colors {
                if !strings.Contains(output, color) {
                        t.Errorf("Expected output to contain color %q", color)
                }
        }
}

func TestMain(t *testing.T) {
        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        done := make(chan bool)
        go func() {
                main()
                done <- true
        }()

        select {
        case <-done:
        case <-time.After(15 * time.Second):
                t.Fatal("main() took too long to complete")
        }

        w.Close()
        os.Stdout = oldStdout

        var buf bytes.Buffer
        io.Copy(&buf, r)
        output := buf.String()

        if len(output) == 0 {
                t.Error("Expected main() to produce output")
        }

        if !strings.Contains(output, "WELCOME TO VIPE CODING") {
                t.Error("Expected main() to call decoratedOutput()")
        }
}

func TestColorConstants(t *testing.T) {
        constants := map[string]string{
                "Reset":  Reset,
                "Red":    Red,
                "Green":  Green,
                "Yellow": Yellow,
                "Blue":   Blue,
                "Purple": Purple,
                "Cyan":   Cyan,
                "White":  White,
                "Bold":   Bold,
        }

        for name, value := range constants {
                if value == "" {
                        t.Errorf("Expected %s constant to be non-empty", name)
                }
                if !strings.HasPrefix(value, "\033[") {
                        t.Errorf("Expected %s constant to be ANSI escape sequence, got %q", name, value)
                }
        }
}
