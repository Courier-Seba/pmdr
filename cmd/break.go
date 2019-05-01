package cmd

import (
    "fmt"

    "pmdr/cmd/blocks"

    "github.com/spf13/cobra"
)

// breakCmd represents the break command
var breakCmd = &cobra.Command{
    Use:   "break",
    Short: "Break timer",
    Long: `Launch the timer of a break. By default last 5 minutes`,
    Run: func(cmd *cobra.Command, args []string) {
        // Read the flags
        delay_input, _:= cmd.Flags().GetString("time")
        soundPath, _:= cmd.Flags().GetString("sound")

        // Call timer
        blocks.Timer(delay_input)
        // Timer Finish
        fmt.Println(`
             ____  ____  _____    _    _  __   _____     _______ ____  
            | __ )|  _ \| ____|  / \  | |/ /  / _ \ \   / / ____|  _ \ 
            |  _ \| |_) |  _|   / _ \ | ' /  | | | \ \ / /|  _| | |_) |
            | |_) |  _ <| |___ / ___ \| . \  | |_| |\ V / | |___|  _ < 
            |____/|_| \_\_____/_/   \_\_|\_\  \___/  \_/  |_____|_| \_\
        `)
        if soundPath != "" {
            fmt.Println("Press Ctrl+C to stop it")
            blocks.SoundPlayer(soundPath)
        } else {
            fmt.Println(">>No sound specified")
        }
    },
}

func init() {
    rootCmd.AddCommand(breakCmd)
    breakCmd.Flags().StringP("time", "t", "5m", "Time of the break")
    breakCmd.Flags().StringP("sound", "s", "", "Sound to play")
}
