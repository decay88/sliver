package command

import (
	"fmt"
	"sort"

	consts "sliver/client/constants"

	"github.com/desertbit/columnize"
	"github.com/desertbit/grumble"
)

func helpCmd(app *grumble.App, isShell bool) {
	printHelp(app)
	fmt.Println()
}

func printHelp(a *grumble.App) {
	config := columnize.DefaultConfig()
	config.Delim = "|"
	config.Glue = "  "
	config.Prefix = "  "
	// Group the commands by their help group if present.
	groups := make(map[string]*grumble.Commands)
	for _, c := range a.Commands().All() {
		key := c.HelpGroup
		if ActiveSliver.Sliver != nil {
			if ActiveSliver.Sliver.GetOS() != "windows" && key == consts.SliverWinHelpGroup {
				continue
			}
		} else {
			if key == consts.SliverHelpGroup || key == consts.SliverWinHelpGroup {
				continue
			}
		}
		if len(key) == 0 {
			key = "Commands:"
		}
		cc := groups[key]
		if cc == nil {
			cc = new(grumble.Commands)
			groups[key] = cc
		}
		cc.Add(c)
	}

	// Sort the map by the keys.
	var keys []string
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print each commands group.
	for _, headline := range keys {
		cc := groups[headline]
		cc.Sort()

		var output []string
		for _, c := range cc.All() {
			name := c.Name
			for _, a := range c.Aliases {
				name += ", " + a
			}
			output = append(output, fmt.Sprintf("%s | %v", name, c.Help))
		}

		if len(output) > 0 {
			fmt.Println()
			printHeadline(a.Config(), headline)
			fmt.Printf("%s\n", columnize.Format(output, config))
		}
	}
}

func headlinePrinter(c *grumble.Config) func(v ...interface{}) (int, error) {
	if c.NoColor || c.HelpHeadlineColor == nil {
		return fmt.Println
	}
	return c.HelpHeadlineColor.Println
}

func printHeadline(c *grumble.Config, s string) {
	println := headlinePrinter(c)
	if c.HelpHeadlineUnderline {
		println(s)
		u := ""
		for i := 0; i < len(s); i++ {
			u += "="
		}
		println(u)
	} else {
		println(s)
	}
}
