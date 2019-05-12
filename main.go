package main

import (
	"fmt"
	"hn-reader/article"
	"hn-reader/hackernews"
	"hn-reader/reddit"

	"github.com/spf13/cobra"
)

func main() {
	var showNum int

	var rootCmd = &cobra.Command{Use: "hn-reader"}

	var cmdHackerNews = &cobra.Command{
		Use:   "hkns",
		Short: "Read hacker news",
		Run: func(cmd *cobra.Command, args []string) {
			hnr := hackernews.NewHackerNewsReader()
			list, _ := hnr.TopArticles(showNum)
			for _, v := range list {
				ar, _ := hnr.GetArticle(v)
				fmt.Println(article.Summerized(ar))
			}
		},
	}

	var cmdRedditGolang = &cobra.Command{
		Use:   "rdgl",
		Short: "Read reddit /r/golang",
		Run: func(cmd *cobra.Command, args []string) {
			rdr := reddit.NewRedditReader("https://www.reddit.com/r/golang")
			list, _ := rdr.TopArticles(showNum)
			for _, v := range list {
				ar, _ := rdr.GetArticle(v)
				fmt.Println(article.Summerized(ar))
			}
		},
	}

	rootCmd.AddCommand(cmdHackerNews, cmdRedditGolang)
	rootCmd.PersistentFlags().IntVarP(&showNum, "number", "n", 10, "Specify number of top articles")
	rootCmd.Execute()

}
