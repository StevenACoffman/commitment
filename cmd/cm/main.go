package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/StevenACoffman/commitment/pkg"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Git Commit Emoji",
		Items: pkg.Gitmojis,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	gitCommitEmoji := pkg.Gitmojis[i].Emoji
	if gitCommitEmoji != "" {
		// add some padding space before actual one line subject
		gitCommitEmoji = gitCommitEmoji + " "
	}

	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("📝  Enter your JIRA issue url (or hit enter for none).")
	issue := pkg.GetSingleLineInput(scn)
	if issue == "" {
		issue = "\"none\""
	}
	// TODO: else if not url, then prefix with JIRA issue

	fmt.Println("🤔  Enter your one-line summary:")
	subject := pkg.GetSingleLineInput(scn)

	fmt.Println("✍️  Enter/Paste your multiline summary. On any empty line, Ctrl-] and Enter to save it:")

	body := pkg.GetMultiLineInput(scn)

	fmt.Println("⚗️  Enter/Paste your multiline Test Plan description. On any empty line, Ctrl-] and Enter to save it:")
	testPlan := pkg.GetMultiLineInput(scn)

	commitMessage := fmt.Sprintf("%s%s\n\n%s\n\nIssue: %s\n\nTest plan:\n\n%s\n",
		gitCommitEmoji,
		subject,
		strings.Join(body, "\n"), issue,
		strings.Join(testPlan, "\n"))

	pipe := strings.NewReader(commitMessage)

	flags := pkg.GetFlags()
	commitFlags := append(flags, "-F", "-")
	gitArgs := append([]string{"commit"}, commitFlags...)
	// Get the flags to this program, and pass through to git
	// for git command
	// insert flags after "commit" arg, before "-F"
	// go run main.go -s -S -a should work
	commit := exec.Command("git", gitArgs...)
	commit.Stdin = pipe

	// Run and get the output of commit.
	res, _ := commit.Output()

	fmt.Println(string(res))
}
