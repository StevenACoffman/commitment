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
		gitCommitEmoji += " "
	}

	fmt.Println("📝  Enter your JIRA issue url (or hit enter for none).")
	issue := pkg.GetSingleLineInput(bufio.NewScanner(os.Stdin))
	if issue == "" {
		issue = "\"none\""
	}
	// TODO: else if not url, then prefix with JIRA issue

	fmt.Println("🤔  Enter your one-line summary:")
	subject := pkg.GetSingleLineInput(bufio.NewScanner(os.Stdin))

	fmt.Println("✍️  Enter/Paste your multiline summary. On any empty line, Ctrl-D to save it:")
	body := pkg.GetMultiLineInput(bufio.NewScanner(os.Stdin))

	fmt.Println("⚗️  Enter/Paste your multiline Test Plan. On any empty line, Ctrl-D to save it:")
	testPlan := pkg.GetMultiLineInput(bufio.NewScanner(os.Stdin))

	commitMessage := fmt.Sprintf("%s%s\n\n%s\n\nIssue: %s\n\nTest plan:\n\n%s\n",
		gitCommitEmoji,
		subject,
		body,
		issue,
		testPlan)

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
