package banner

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	reset = "\033[0m"
	bold  = "\033[1m"
	dim   = "\033[2m"
	green = "\033[38;5;84m"
	gray  = "\033[38;5;245m"
	cyan  = "\033[38;5;51m"
	pink  = "\033[38;5;213m"
)

func color(c int) string { return fmt.Sprintf("\033[38;5;%dm", c) }

var logo = []string{
	"██████╗ ███████╗██╗   ██╗      ██╗    ██╗███████╗██████╗ ",
	"██╔══██╗██╔════╝██║   ██║      ██║    ██║██╔════╝██╔══██╗",
	"██║  ██║█████╗  ██║   ██║█████╗██║ █╗ ██║█████╗  ██████╔╝",
	"██║  ██║██╔══╝  ╚██╗ ██╔╝╚════╝██║███╗██║██╔══╝  ██╔══██╗",
	"██████╔╝███████╗ ╚████╔╝       ╚███╔███╔╝███████╗██████╔╝",
	"╚═════╝ ╚══════╝  ╚═══╝         ╚══╝╚══╝ ╚══════╝╚═════╝ ",
	"",
	" ██████╗ ██████╗ ███╗   ██╗███████╗ ██████╗        █████╗ ██████╗ ██╗",
	"██╔════╝██╔═══██╗████╗  ██║██╔════╝██╔═══██╗      ██╔══██╗██╔══██╗██║",
	"██║     ██║   ██║██╔██╗ ██║███████╗██║   ██║█████╗███████║██████╔╝██║",
	"██║     ██║   ██║██║╚██╗██║╚════██║██║   ██║╚════╝██╔══██║██╔═══╝ ██║",
	"╚██████╗╚██████╔╝██║ ╚████║███████║╚██████╔╝      ██║  ██║██║     ██║",
	" ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝ ╚═════╝       ╚═╝  ╚═╝╚═╝     ╚═╝",
}

var gradient = []int{51, 45, 39, 33, 63, 99, 0, 135, 171, 207, 206, 205, 199}

func PrintBanner(addr string) {
	fmt.Println()
	for i, line := range logo {
		fmt.Println("  " + bold + color(gradient[i]) + line + reset)
	}

	w := 68
	hr := strings.Repeat("─", w)
	row := func(label, value, c string) {
		pad := w - 4 - utf8.RuneCountInString(label) - utf8.RuneCountInString(value)
		fmt.Printf("  %s│%s  %s%s%s%s%s%s%s  %s│%s\n",
			gray, reset, dim, label, reset, c, bold, value, reset+strings.Repeat(" ", pad), gray, reset)
	}

	fmt.Println()
	fmt.Println("  " + gray + "╭" + hr + "╮" + reset)
	row("● Status   ", "ONLINE", green)
	row("➜ Local    ", "http://localhost"+addr, cyan)
	row("◆ Cours    ", "dev-web-conso-api", pink)
	row("◷ Démarré  ", time.Now().Format("02/01/2006 15:04:05"), gray)
	fmt.Println("  " + gray + "╰" + hr + "╯" + reset)
	fmt.Println()
	fmt.Println("  " + dim + "Ctrl+C pour arrêter le serveur" + reset)
	fmt.Println()
}
