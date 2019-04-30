package cmd

import (
    "fmt"
    "time"

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

        // Set timer
        delay, _:= time.ParseDuration(delay_input)
        timerWork := time.NewTimer(delay)
        // Launch timer
        <-timerWork.C
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
    },
}

func init() {
    rootCmd.AddCommand(workCmd)
    workCmd.Flags().StringP("time", "t", "25m", "Tiempo del timer de trabajo")
}
