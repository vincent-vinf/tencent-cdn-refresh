package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

var rootCmd = &cobra.Command{
	Use: os.Args[0],
}
var PurgePathCMD = &cobra.Command{
	Use: "purge-path",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Println("paths is empty")

			return nil
		}
		secretId := cmd.Flag("secret-id").Value.String()
		secretKey := cmd.Flag("secret-key").Value.String()
		region := cmd.Flag("region").Value.String()

		var paths []*string
		for _, s := range args {
			if _, err := url.Parse(s); err != nil {
				return fmt.Errorf("error parsing url(%s): %w", s, err)
			}
			paths = append(paths, &s)
		}

		credential := common.NewCredential(secretId, secretKey)
		client, err := cdn.NewClient(credential, region, profile.NewClientProfile())
		if err != nil {
			return err
		}
		request := cdn.NewPurgePathCacheRequest()
		flushType := "delete"
		request.FlushType = &flushType
		request.Paths = paths
		response, err := client.PurgePathCache(request)
		if err != nil {
			return err
		}
		log.Printf("response: %s", response.ToJsonString())

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().String("secret-id", "", "")
	rootCmd.PersistentFlags().String("secret-key", "", "")
	rootCmd.PersistentFlags().String("region", regions.Shanghai, "")
	_ = rootCmd.MarkPersistentFlagRequired("secret-id")
	_ = rootCmd.MarkPersistentFlagRequired("secret-key")

	rootCmd.AddCommand(PurgePathCMD)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
