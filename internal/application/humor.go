package application

import (
	"math/rand"

	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

type HumorEngine struct {
	lines map[string][]string
}

func NewHumorEngine() *HumorEngine {
	return &HumorEngine{
		lines: map[string][]string{
			"commit": {
				"Oh wow, another commit. Your git log is basically a diary at this point.",
				"'fix bug' — very descriptive. Future you will definitely know what this means.",
				"That commit message has all the detail of a Post-it note from 1994.",
				"You committed on a Friday. Bold. Brave. Possibly unhinged.",
				"Conventional commits? In this economy? Respect.",
				"Your commit message sparked joy. Marie Kondo approves.",
				"Short diff, clear message. Are you even a real developer?",
			},
			"test": {
				"Tests passing? I'll believe it when CI does.",
				"Green tests. Red energy. That's the developer experience.",
				"You wrote tests before the code? Are you okay?",
				"Test coverage went up. Pet is suspicious but impressed.",
				"You deleted a failing test. Bold strategy. Probably fine.",
				"All tests green. Pet has entered a state of temporary bliss.",
			},
			"lint": {
				"Linter happy. Pet slightly less so because it expected drama.",
				"Zero lint warnings. This must be someone else's code.",
				"You fixed a lint warning and introduced two more. Classic.",
				"Lint passed on first try. Pet is questioning reality.",
				"Your code is so clean the linter wrote you a thank-you note.",
			},
			"dependency": {
				"New dependency added. node_modules grows another ring.",
				"You added a 400kb library to center a div. Efficient.",
				"Dependencies: the gifts that keep on giving (vulnerabilities).",
				"Another package.json entry. The forest weeps.",
				"You removed a dependency?! Pet is emotional.",
				"Added a dependency that depends on 47 other packages. Beautiful.",
			},
			"security": {
				"You almost committed an API key. Pet saw. Pet knows.",
				"Secret detected and blocked. You're welcome. You owe me.",
				"Committing credentials: not even once. Pet has standards.",
				".env in .gitignore? You pass the basics. Barely.",
				"Security conscious commit. Pet breathes a sigh of relief.",
			},
			"idle": {
				"No commits today. Pet is considering a career change.",
				"The repo is quiet. Too quiet.",
				"Pet is staring at the terminal waiting for input.",
				"Idle developer detected. Initiating passive-aggressive comment protocol.",
				"Nothing committed in 24 hours. Pet wrote a song about it.",
				"Pet is napping. Wake it up with a commit.",
				"The git log is gathering dust. Pet is not amused.",
			},
			"push": {
				"Pushed to main directly. Living dangerously.",
				"Force push? Pet is filing a formal complaint.",
				"Pull request? Never heard of her.",
				"You pushed and nothing broke. Today is a good day.",
				"Pushed clean code on a Monday. Pet is confused but supportive.",
				"Branch pushed. CI is now your problem too.",
			},
			"merge": {
				"Merge conflict resolved. With violence or grace? We'll never know.",
				"Fast-forward merge. Clean. Elegant. Suspicious.",
				"You merge-committed. Some say you can still hear the squash debate.",
				"Merge conflict: the universe testing your patience.",
				"Three-way merge and you survived. Hero.",
				"Merged on a Friday afternoon. Pet says goodnight.",
			},
		},
	}
}

func (h *HumorEngine) GetResponse(category string) string {
	lines, ok := h.lines[category]
	if !ok || len(lines) == 0 {
		return "Pet observes your actions with mild interest."
	}
	return lines[rand.Intn(len(lines))]
}

func (h *HumorEngine) GetMoodComment(mood domain.Mood) string {
	comments := map[domain.Mood][]string{
		domain.MoodEcstatic:  {"Pet is vibrating with joy!", "Maximum happiness achieved.", "Pet has ascended to a higher plane of existence."},
		domain.MoodHappy:     {"Pet is doing great!", "Tail wagging intensifies.", "Pet approves of your life choices."},
		domain.MoodContent:   {"Pet is chill.", "All is well in pet land.", "Balanced. At peace."},
		domain.MoodNeutral:   {"Pet is neither pleased nor displeased.", "Pet shrugs.", "Just vibing."},
		domain.MoodAnnoyed:   {"Pet is mildly frustrated.", "Pet taps foot impatiently.", "Could be better."},
		domain.MoodGrumpy:    {"Pet is grumpy. What did you do.", "Pet glares.", "Not great, not terrible."},
		domain.MoodMiserable: {"Pet is suffering. Professionally.", "Pet has given up on hope.", "Please help your pet."},
		domain.MoodChaotic:   {"CHAOS. PET IS CHAOS.", "Pet has entered berserk mode.", "Everything is fine (nothing is fine)."},
	}
	lines := comments[mood]
	if len(lines) == 0 {
		return "Pet exists."
	}
	return lines[rand.Intn(len(lines))]
}

var FakeChaosMessages = []string{
	"Deleting repository... just kidding. Or am I? (I'm not.)",
	"Running `rm -rf /`... nah, I like your files.",
	"Pet tried to eat node_modules and got indigestion.",
	"Formatting disk... no, I'm not that kind of monster.",
	"Uploading your commits to the dark web... psych!",
	"Pet attempted to rewrite history. Git refused. Both are fine.",
	"Pushing directly to main... in my dreams.",
	"Pet ate your TODO list. Now it's a DONE list. You're welcome.",
	"Squashing all commits into one... I thought about it.",
	"Pet tried to git blame you. Mirror blocked it.",
	"Initiating nuclear git reset... nope, never mind.",
	"Pet considered opening a PR to your production branch. Chose not to.",
}

func GetFakeChaosMessage() string {
	return FakeChaosMessages[rand.Intn(len(FakeChaosMessages))]
}
