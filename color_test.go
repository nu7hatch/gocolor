package color

import(
	"testing"
	"fmt"
	"os"
)

func TestAllColorsAndEffects(t *testing.T) {
	Println(Black, "Black")
	Println(Red, "Red")
	Println(Green, "Green")
	Println(Yellow, "Yellow")
	Println(Blue, "Blue")
	Println(Magenta, "Magenta")
	Println(Cyan, "Cyan")
	Println(White, "White")
	Print(Red, "Red\n")
	Printf(Green, "Green %s\n", "yeah!")
	fmt.Println(Sprintf(Cyan, "Cyan %s", "yeah!"))
	fmt.Print(Sprintln(Magenta, "Magenta"))
	Fprint(Blue, os.Stderr, "Blue\n")
	Fprintf(Red, os.Stderr, "Red %s\n", "yeah!")
	Fprintln(Yellow, os.Stderr, "Yellow")
	Println(Red + Bright, "Bright red")
	Println(Blue + Italic, "Italic blue")
	Println(Underline + Yellow, "Underline yellow")
	Println(Blink + Red, "Blink red")
	Println(Inverse + White, "Invert white")
	Println(Hide + Black, "Hide black")
}