package gui

import (
	"os"
	"strings"

	"github.com/jesseduffield/lazygit/pkg/utils"
)

func (gui *Gui) prepareEncodings() {
	locale := os.Getenv("LC_CTYPE")
	if locale == "" {
		locale = os.Getenv("LANG")
	}

	usingASCII := strings.Contains(locale, "8859")

	if usingASCII {
		gui.g.ASCII = true
		gui.encodedStrings = utils.EncodedStrings{
			UpArrow:    "^",
			DownArrow:  "v",
			LeftArrow:  "<-",
			RightArrow: "->",
		}
	} else {
		gui.encodedStrings = utils.EncodedStrings{
			UpArrow:    "↑",
			DownArrow:  "↓",
			LeftArrow:  "←",
			RightArrow: "→",
		}
	}
}
