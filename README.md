# commitment
Interactive commit message generator for the command line.

### Installation

```
go get -u github.com/StevenACoffman/commitment/cmd/cm
```

### Usage

Commitment (or `cm`) will passthrough any and all the flags you give it to `git commit`, but then will
interactively help you generate a nicely formatted git commit message.

So for instance, if you type `cm -s -S -a` you will be prompted to optionally pick an emoji, a jira issue, one-line
summary, a multiline summary, and a multiline Test Plan.

```
✔ ✨ - Introducing new features.
📝 Enter your JIRA issue (or hit enter for none).

🤔  Enter your one-line summary:
Add new emoji stuff
✍️  Enter/Paste your multiline summary. On any empty line, Ctrl-] and Enter to save it:
I think it will be great
^]
⚗️  Enter/Paste your multiline Test Plan description. On any empty line, Ctrl-] and Enter to save it:
Run it and see!
^]

```
This will result in a commit message like:
``` 
✨ Add new emoji stuff

I think it will be great

Issue: none

Test Plan:
Run it and see
```
