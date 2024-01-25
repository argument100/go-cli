package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "My CLI application",
		Run:   runPrompt,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runPrompt(cmd *cobra.Command, args []string) {
	// ここでプロンプトを実装します
	// Cobra は直接的なプロンプトサポートを提供しません
	// したがって、通常の Go の入力読み取り機能を使用するか、サードパーティのライブラリを使用する必要があります

	fmt.Println("選択してください")
	fmt.Println("1: Aについて")
	fmt.Println("2: Bについて")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Aについての情報を表示します。")
	case 2:
		fmt.Println("Bについての情報を表示します。")
	default:
		fmt.Println("無効な選択です。")
	}
}

