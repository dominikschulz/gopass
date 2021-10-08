package cui

import (
	"context"
	"fmt"

	"github.com/gopasspw/gopass/pkg/ctxutil"
	"gopkg.in/AlecAivazis/survey.v1"
)

// GetSelection show a navigateable multiple-choice list to the user
// and returns the selected entry along with the action
func GetSelection(ctx context.Context, prompt string, choices []string) (string, int) {
	if ctxutil.IsAlwaysYes(ctx) || !ctxutil.IsInteractive(ctx) {
		return "impossible", 0
	}

	choices = append([]string{"abort"}, choices...)

	sel := ""
	if err := survey.AskOne(&survey.Select{
		Message: prompt,
		Options: choices,
	}, &sel, nil); err != nil {
		fmt.Println("Error: ", err)
		return "aborted", 0
	}
	i := -1
	for n, v := range choices {
		if v == sel {
			i = n
			break
		}
	}
	if i == 0 {
		return "aborted", 0
	}
	fmt.Printf("selection: %s - pos: %d\n", sel, i)
	return sel, i
	// for i, c := range choices {
	// 	fmt.Print(color.GreenString("[%  d]", i))
	// 	fmt.Printf(" %s\n", c)
	// }
	// fmt.Println()
	// var i int
	// for {
	// 	var err error
	// 	i, err = termio.AskForInt(ctx, prompt, 0)
	// 	if err == nil && i < len(choices) {
	// 		break
	// 	}
	// 	if errors.Is(err, termio.ErrAborted) {
	// 		return "aborted", 0
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }
	// fmt.Println(i)
	//return "default", i
}
