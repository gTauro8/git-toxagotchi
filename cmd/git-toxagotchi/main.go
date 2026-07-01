package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gTauro8/git-toxagotchi/internal/application"
	appconfig "github.com/gTauro8/git-toxagotchi/internal/infrastructure/config"
	"github.com/gTauro8/git-toxagotchi/internal/infrastructure/storage"
	"github.com/gTauro8/git-toxagotchi/internal/tui"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "git-toxagotchi",
		Short: "A terminal ASCII pet companion that reacts to your Git habits",
		Long:  "Git-Toxagotchi is a Tamagotchi-style pet that lives in your terminal and evolves based on the quality of your git commits.",
	}

	root.AddCommand(initCmd())
	root.AddCommand(statusCmd())
	root.AddCommand(watchCmd())
	root.AddCommand(feedCmd())
	root.AddCommand(playCmd())
	root.AddCommand(achievementsCmd())
	root.AddCommand(analyzeCmd())
	root.AddCommand(hookCmd())
	root.AddCommand(shareCmd())
	root.AddCommand(configCmd())
	root.AddCommand(resetCmd())

	return root
}

func openStore() (*storage.SQLiteStore, *appconfig.Config, error) {
	cfg, err := appconfig.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("load config: %w", err)
	}

	dbDir := cfg.DBPath[:strings.LastIndex(cfg.DBPath, "/")]
	if dbDir != "" {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return nil, nil, fmt.Errorf("create db dir: %w", err)
		}
	}

	store, err := storage.NewSQLiteStore(cfg.DBPath)
	if err != nil {
		return nil, nil, fmt.Errorf("open store: %w", err)
	}
	return store, cfg, nil
}

func initCmd() *cobra.Command {
	var name string
	var skipWizard bool
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create a new pet (runs setup wizard on first launch)",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Run the config wizard if no config file exists yet and not skipped.
			cfg, err := appconfig.Load()
			if err != nil {
				return err
			}
			_, statErr := os.Stat(appconfig.Path())
			if os.IsNotExist(statErr) && !skipWizard {
				result, err := tui.RunWizard(cfg)
				if err != nil {
					return fmt.Errorf("config wizard: %w", err)
				}
				if result != nil {
					cfg = result
				}
			}

			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			if name != "" {
				cfg.PetName = name
			}

			svc := application.NewService(store)
			pet, err := svc.GetOrCreatePet(cfg.PetName)
			if err != nil {
				return err
			}
			fmt.Printf("🐣 %s is alive! Stage: %s | Energy: %d | Mood: %s\n", pet.Name, pet.Stage, pet.Energy, pet.Mood)
			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name for your pet (overrides wizard)")
	cmd.Flags().BoolVar(&skipWizard, "no-wizard", false, "Skip the setup wizard")
	return cmd
}

func statusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show your pet's current status",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			pet, err := svc.GetPet()
			if err != nil {
				return err
			}
			if pet == nil {
				fmt.Println("No pet found. Run: git-toxagotchi init")
				return nil
			}

			fmt.Printf("Name:       %s\n", pet.Name)
			fmt.Printf("Stage:      %s\n", pet.Stage)
			fmt.Printf("Mood:       %s\n", pet.Mood)
			fmt.Printf("Energy:     %d/100\n", pet.Energy)
			fmt.Printf("Hunger:     %d/100\n", pet.Hunger)
			fmt.Printf("Stress:     %d/100\n", pet.Stress)
			fmt.Printf("Trust:      %d/100\n", pet.Trust)
			fmt.Printf("Chaos:      %d/100\n", pet.Chaos)
			fmt.Printf("Experience: %d\n", pet.Experience)
			return nil
		},
	}
}

func watchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Open the interactive TUI",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			pet, err := svc.GetPet()
			if err != nil {
				return err
			}

			m := tui.NewModel(pet, svc)
			p := tea.NewProgram(m, tea.WithAltScreen())
			_, err = p.Run()
			return err
		},
	}
}

func feedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "feed",
		Short: "Feed your pet",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			pet, err := svc.GetPet()
			if err != nil || pet == nil {
				fmt.Println("No pet found. Run: git-toxagotchi init")
				return nil
			}

			msg, err := svc.FeedPet(pet)
			if err != nil {
				return err
			}
			fmt.Println(msg)
			return nil
		},
	}
}

func playCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: "Play with your pet",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			pet, err := svc.GetPet()
			if err != nil || pet == nil {
				fmt.Println("No pet found. Run: git-toxagotchi init")
				return nil
			}

			msg, err := svc.PlayWithPet(pet)
			if err != nil {
				return err
			}
			fmt.Println(msg)
			return nil
		},
	}
}

func achievementsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "achievements",
		Short: "Show achievements",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			achs, err := svc.GetAchievements()
			if err != nil {
				return err
			}

			fmt.Println("Achievements:")
			for _, a := range achs {
				status := "  [ ]"
				if a.Unlocked {
					status = "  [x]"
				}
				fmt.Printf("%s %s %s - %s\n", status, a.Icon, a.Name, a.Description)
			}
			return nil
		},
	}
}

func analyzeCmd() *cobra.Command {
	var hookMode bool
	c := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze staged changes and update pet",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			pet, err := svc.GetPet()
			if err != nil || pet == nil {
				fmt.Println("No pet found. Run: git-toxagotchi init")
				return nil
			}

			msg := gitOutput("log", "-1", "--pretty=%B")
			diff := gitOutput("diff", "--cached")

			response, err := svc.ProcessCommit(context.Background(), pet, msg, diff)
			if err != nil {
				return err
			}
			fmt.Println(response)
			// In hook mode, exit with code 1 if secrets were detected.
			if hookMode {
				a := application.NewAnalyzer()
				analysis := a.AnalyzeCommit(msg, diff)
				if analysis.SecretsDetected {
					fmt.Fprintln(os.Stderr, "\n🚨 Commit blocked: possible secret detected in staged files.")
					fmt.Fprintln(os.Stderr, "   Review the diff with `git diff --cached` and remove sensitive data.")
					fmt.Fprintln(os.Stderr, "   Use --no-verify to bypass (at your own risk).")
					os.Exit(1)
				}
			}
			return nil
		},
	}
	c.Flags().BoolVar(&hookMode, "hook", false, "Run in pre-commit hook mode (blocks on secrets)")
	return c
}

func gitOutput(args ...string) string {
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func hookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hook",
		Short: "Manage Git hooks",
	}
	cmd.AddCommand(hookInstallCmd())
	cmd.AddCommand(hookUninstallCmd())
	return cmd
}

func hookInstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install the pre-commit hook (opt-in)",
		RunE: func(cmd *cobra.Command, args []string) error {
			hookPath, err := gitHookPath()
			if err != nil {
				return err
			}
			script := "#!/bin/sh\ngit-toxagotchi analyze --hook\n"
			if err := os.WriteFile(hookPath, []byte(script), 0755); err != nil {
				return fmt.Errorf("write hook: %w", err)
			}
			fmt.Printf("✅ Pre-commit hook installed at %s\n", hookPath)
			fmt.Println("   To uninstall: git-toxagotchi hook uninstall")
			return nil
		},
	}
}

func hookUninstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Remove the pre-commit hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			hookPath, err := gitHookPath()
			if err != nil {
				return err
			}
			if _, err := os.Stat(hookPath); os.IsNotExist(err) {
				fmt.Println("No hook found.")
				return nil
			}
			if err := os.Remove(hookPath); err != nil {
				return fmt.Errorf("remove hook: %w", err)
			}
			fmt.Println("✅ Pre-commit hook removed.")
			return nil
		},
	}
}

func gitHookPath() (string, error) {
	out := gitOutput("rev-parse", "--git-dir")
	if out == "" {
		return "", fmt.Errorf("not inside a Git repository")
	}
	return filepath.Join(out, "hooks", "pre-commit"), nil
}

func shareCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "share",
		Short: "Generate a badge and markdown snippet for your README",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()

			svc := application.NewService(store)
			pet, err := svc.GetPet()
			if err != nil || pet == nil {
				fmt.Println("No pet found. Run: git-toxagotchi init")
				return nil
			}

			home, _ := os.UserHomeDir()
			outDir := filepath.Join(home, ".local", "share", "git-toxagotchi")
			if err := os.MkdirAll(outDir, 0755); err != nil {
				return fmt.Errorf("create output dir: %w", err)
			}

			badgePath := filepath.Join(outDir, "badge.svg")
			svg := generateBadgeSVG(string(pet.Stage), string(pet.Mood))
			if err := os.WriteFile(badgePath, []byte(svg), 0644); err != nil {
				return fmt.Errorf("write badge: %w", err)
			}

			markdown := fmt.Sprintf("![Git-Toxagotchi](%s)", badgePath)
			fmt.Printf("✅ Badge saved to: %s\n\n", badgePath)
			fmt.Println("Add this to your README.md:")
			fmt.Println()
			fmt.Println(markdown)
			fmt.Println()
			fmt.Println("(Future: git-toxagotchi share --gist to publish to GitHub Gist)")
			return nil
		},
	}
}

func resetCmd() *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Delete your pet and start over",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !force {
				fmt.Print("⚠️  This will delete your pet and all progress. Are you sure? [y/N] ")
				var answer string
				fmt.Scanln(&answer)
				if strings.ToLower(strings.TrimSpace(answer)) != "y" {
					fmt.Println("Cancelled.")
					return nil
				}
			}
			store, _, err := openStore()
			if err != nil {
				return err
			}
			defer store.Close()
			if err := store.DeleteAllData(); err != nil {
				return err
			}
			fmt.Println("🪦 Pet deleted. Run `git-toxagotchi init` to start over.")
			return nil
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation prompt")
	return cmd
}

func generateBadgeSVG(stage, mood string) string {
	label := "git-toxagotchi"
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="220" height="20">
  <rect rx="3" width="220" height="20" fill="#555"/>
  <rect rx="3" x="120" width="100" height="20" fill="#4c1"/>
  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text x="60" y="14">%s</text>
    <text x="170" y="14">%s | %s</text>
  </g>
</svg>`, label, stage, mood)
}

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		// Default: show config when called with no subcommand.
		RunE: func(cmd *cobra.Command, args []string) error {
			return showConfig()
		},
	}
	cmd.AddCommand(configShowCmd())
	cmd.AddCommand(configSetCmd())
	cmd.AddCommand(configEditCmd())
	return cmd
}

func configShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return showConfig()
		},
	}
}

func showConfig() error {
	cfg, err := appconfig.Load()
	if err != nil {
		return err
	}
	fmt.Printf("pet_name:      %s\n", cfg.PetName)
	fmt.Printf("theme:         %s\n", cfg.Theme)
	fmt.Printf("llm_enabled:   %v\n", cfg.LLMEnabled)
	fmt.Printf("hook_blocking: %v\n", cfg.HookBlocking)
	fmt.Printf("db_path:       %s\n", cfg.DBPath)
	fmt.Printf("\nConfig file: %s\n", appconfig.Path())
	return nil
}

func configSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a config value",
		Long: `Set a single config value. Available keys:
  pet_name      — name of your pet
  theme         — color theme (default, dracula, nord, solarized)
  llm_enabled   — enable LLM feedback (true/false)
  hook_blocking — block dangerous commits in pre-commit hook (true/false)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key, value := args[0], args[1]
			cfg, err := appconfig.Load()
			if err != nil {
				return err
			}
			switch key {
			case "pet_name":
				cfg.PetName = value
			case "theme":
				valid := map[string]bool{"default": true, "dracula": true, "nord": true, "solarized": true}
				if !valid[value] {
					return fmt.Errorf("unknown theme %q — choose from: default, dracula, nord, solarized", value)
				}
				cfg.Theme = value
			case "llm_enabled":
				cfg.LLMEnabled = value == "true" || value == "1" || value == "yes"
			case "hook_blocking":
				cfg.HookBlocking = value == "true" || value == "1" || value == "yes"
			default:
				return fmt.Errorf("unknown key %q", key)
			}
			if err := appconfig.Save(cfg); err != nil {
				return err
			}
			fmt.Printf("✅ %s = %s\n", key, value)
			return nil
		},
	}
}

func configEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Re-open the interactive setup wizard",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := appconfig.Load()
			if err != nil {
				return err
			}
			result, err := tui.RunWizard(cfg)
			if err != nil {
				return err
			}
			if result == nil {
				fmt.Println("Cancelled — config unchanged.")
				return nil
			}
			fmt.Printf("✅ Config saved to %s\n", appconfig.Path())
			return nil
		},
	}
}
