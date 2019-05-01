package cmd

import (
    "fmt"

    "pmdr/cmd/blocks"

    "github.com/spf13/cobra"
)

// workCmd represents the work command
var workCmd = &cobra.Command{
    Use:   "work",
    Short: "Work timer",
    Long: `Launcs a timer for work. Defaults its 25 mins`,
    Run: func(cmd *cobra.Command, args []string) {
        // Read the flags
        delay_input, _:= cmd.Flags().GetString("time")
        soundPath, _:= cmd.Flags().GetString("sound")

        // Call timer
        blocks.Timer(delay_input)
        // Timer Finish
        fmt.Println(`
                      _______  _______  _          ______   _______  _        _______ 
            |\     /|(  ___  )(  ____ )| \    /\  (  __  \ (  ___  )( (    /|(  ____ \
            | )   ( || (   ) || (    )||  \  / /  | (  \  )| (   ) ||  \  ( || (    \/
            | | _ | || |   | || (____)||  (_/ /   | |   ) || |   | ||   \ | || (__    
            | |( )| || |   | ||     __)|   _ (    | |   | || |   | || (\ \) ||  __)   
            | || || || |   | || (\ (   |  ( \ \   | |   ) || |   | || | \   || (      
            | () () || (___) || ) \ \__|  /  \ \  | (__/  )| (___) || )  \  || (____/\
            (_______)(_______)|/   \__/|_/    \/  (______/ (_______)|/    )_)(_______/
        `)
        fmt.Println("Press Ctrl+C to stop it")
        blocks.SoundPlayer(soundPath)
    },
}

func init() {
    rootCmd.AddCommand(workCmd)
    workCmd.Flags().StringP("time", "t", "25m", "Time of work")
    workCmd.Flags().StringP("sound", "s", "", "Sound to play")
}
