package colors

import "github.com/fatih/color"

func Bold(s string) string {
	return color.New(color.Bold).SprintFunc()(s)
}

func Red(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgRed, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiRed).SprintFunc()(s)
}

func Green(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgGreen, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiGreen).SprintFunc()(s)
}

func Yellow(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgYellow, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiYellow).SprintFunc()(s)
}

func Blue(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgBlue, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiBlue).SprintFunc()(s)
}

func Magenta(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgMagenta, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiMagenta).SprintFunc()(s)
}

func Cyan(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgCyan, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiCyan).SprintFunc()(s)
}

func Black(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgBlack, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiBlack).SprintFunc()(s)
}

func White(s string, args ...bool) string {
	if len(args) > 0 && args[0] {
		return color.New(color.FgWhite, color.Bold).SprintFunc()(s)
	}

	return color.New(color.FgHiWhite).SprintFunc()(s)
}
