package consoleColors

import "github.com/fatih/color"

type CreatedColors struct {
	Red          func(a ...interface{}) string
	RedBold      func(a ...interface{}) string
	Success      func(a ...interface{}) string
	Green        func(a ...interface{}) string
	Yellow       func(a ...interface{}) string
	YellowBoldUl func(a ...interface{}) string
	FgCyan       func(a ...interface{}) string
}

func Colors() CreatedColors {
	return CreatedColors{
		Red:          color.New(color.FgRed).SprintFunc(),
		RedBold:      color.New(color.FgRed, color.Bold).SprintFunc(),
		Success:      color.New(color.FgGreen).SprintFunc(),
		Green:        color.New(color.FgGreen, color.BgBlack).Add(color.BlinkSlow).SprintFunc(),
		Yellow:       color.New(color.FgYellow).SprintFunc(),
		YellowBoldUl: color.New(color.FgYellow, color.Bold).Add(color.Underline).SprintFunc(),
		FgCyan:       color.New(color.FgCyan, color.Bold).SprintFunc(),
	}
}
