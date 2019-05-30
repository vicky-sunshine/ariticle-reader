package main

import (
	"artread/article"
	"artread/hackernews"
	"artread/reddit"
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var showNum int

	var rootCmd = &cobra.Command{Use: "artread"}

	var cmdHackerNews = &cobra.Command{
		Use:   "hkns",
		Short: "Read hacker news",
		Run: func(cmd *cobra.Command, args []string) {
			hnr := hackernews.NewReader("https://hacker-news.firebaseio.com/v0")
			list, _ := hnr.TopArticles(showNum)
			for _, v := range list {
				ar, _ := hnr.GetArticle(v)
				fmt.Println(article.Summarized(ar))
			}
		},
	}

	var cmdRedditGolang = &cobra.Command{
		Use:   "rdgl",
		Short: "Read reddit /r/golang",
		Run: func(cmd *cobra.Command, args []string) {
			rdr := reddit.NewReader("https://www.reddit.com/r/golang")
			list, _ := rdr.TopArticles(showNum)
			for _, v := range list {
				ar, _ := rdr.GetArticle(v)
				fmt.Println(article.Summarized(ar))
			}
		},
	}

	rootCmd.AddCommand(cmdHackerNews, cmdRedditGolang)
	rootCmd.PersistentFlags().IntVarP(&showNum, "number", "n", 10, "Specify number of top articles")
	rootCmd.Execute()

}
