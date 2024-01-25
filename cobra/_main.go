package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "myapp",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("あなたの名前は何ですか？ ")
			name, _ := reader.ReadString('\n')
			fmt.Printf("こんにちは、%sさん！\n", name)
		},
	}

	rootCmd.Execute()
}

