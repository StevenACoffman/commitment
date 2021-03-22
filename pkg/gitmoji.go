package pkg

type Gitmoji struct {
	Emoji       string `json:"Emoji"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
	Name        string `json:"Name"`
}

func (g Gitmoji) String() string {
	return g.Emoji + " - " + g.Description
}

var Gitmojis = []Gitmoji{
	{
		Emoji:       "",
		Code:        "",
		Description: "None - Please omit emoji from git commit",
		Name:        "",
	},
	{
		Emoji:       "✨",
		Code:        ":sparkles:",
		Description: "Introducing new features.",
		Name:        "sparkles",
	},
	{Emoji: "🐛", Code: ":bug:", Description: "Fixing a bug.", Name: "bug"},
	{
		Emoji:       "🎨",
		Code:        ":art:",
		Description: "Improving UI and style files.",
		Name:        "art",
	},
	{
		Emoji:       "👌",
		Code:        ":ok_hand:",
		Description: "Updating Code due to Code review changes.",
		Name:        "ok-hand",
	},
	{
		Emoji:       "♻️",
		Code:        ":recycle:",
		Description: "Refactoring Code.",
		Name:        "recycle",
	},
	{
		Emoji:       "🔀",
		Code:        ":twisted_rightwards_arrows:",
		Description: "Merging branches.",
		Name:        "twisted-rightwards-arrows",
	},
	{
		Emoji:       "🚚",
		Code:        ":truck:",
		Description: "Moving or renaming files.",
		Name:        "truck",
	},
	{
		Emoji:       "⚡️",
		Code:        ":zap:",
		Description: "Improving performance.",
		Name:        "zap",
	},
	{
		Emoji:       "🔥",
		Code:        ":fire:",
		Description: "Removing Code or files.",
		Name:        "fire",
	},
	{
		Emoji:       "🚑",
		Code:        ":ambulance:",
		Description: "Critical hotfix.",
		Name:        "ambulance",
	},
	{
		Emoji:       "📝",
		Code:        ":pencil:",
		Description: "Writing docs.",
		Name:        "pencil",
	},
	{
		Emoji:       "🚀",
		Code:        ":rocket:",
		Description: "Deploying stuff.",
		Name:        "rocket",
	},
	{Emoji: "🎉", Code: ":tada:", Description: "Initial commit.", Name: "tada"},
	{
		Emoji:       "🔒",
		Code:        ":lock:",
		Description: "Fixing security issues.",
		Name:        "lock",
	},
	{
		Emoji:       "🏷️",
		Code:        ":label:",
		Description: "Releasing / Version tags.",
		Name:        "label",
	},
	{
		Emoji:       "🚧",
		Code:        ":construction:",
		Description: "Work in progress.",
		Name:        "construction",
	},
	{
		Emoji:       "💚",
		Code:        ":green_heart:",
		Description: "Fixing CI Build.",
		Name:        "green-heart",
	},
	{
		Emoji:       "🔧",
		Code:        ":wrench:",
		Description: "Changing configuration files.",
		Name:        "wrench",
	},
	{
		Emoji:       "🌐",
		Code:        ":globe_with_meridians:",
		Description: "Internationalization and localization.",
		Name:        "globe-with-meridians",
	},
	{
		Emoji:       "✏️",
		Code:        ":pencil2:",
		Description: "Fixing typos.",
		Name:        "pencil",
	},
	{
		Emoji:       "📦",
		Code:        ":package:",
		Description: "Updating compiled files or packages.",
		Name:        "package",
	},
	{
		Emoji:       "🍱",
		Code:        ":bento:",
		Description: "Adding or updating assets.",
		Name:        "bento",
	},
	{
		Emoji:       "♿️",
		Code:        ":wheelchair:",
		Description: "Improving accessibility.",
		Name:        "wheelchair",
	},
	{
		Emoji:       "🚸",
		Code:        ":children_crossing:",
		Description: "Improving user experience / usability.",
		Name:        "children-crossing",
	},
}
