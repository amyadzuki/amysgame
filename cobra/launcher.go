// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// launcherCmd represents the launcher command
var launcherCmd = &cobra.Command{
	Use:   "launcher",
	Short: "The launcher for games by amyadzuki",
	Long: `The launcher for games by amyadzuki`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("launcher called")
	},
}

func init() {
	rootCmd.AddCommand(launcherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// launcherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// launcherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
