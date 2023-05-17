package cmd

import (
	"fmt"
	"log"

	"github.com/iamrajiv/neon-db-cli/apis/api_keys"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "neondb",
	Short: "Neon Database CLI",
}

// Get a list of API keys
var listAPIKeys = &cobra.Command{
	Use:   "listapikeys",
	Short: "List API keys for your Neon account",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := api_keys.ListAPIKeys()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	},
}

// Create an API key
var createAPIKey = &cobra.Command{
	Use:   "createapikey",
	Short: "Create a new API key for your Neon account",
	Run: func(cmd *cobra.Command, args []string) {
		keyName := cmd.Flag("keyname").Value.String()

		output, err := api_keys.CreateAPIKey(keyName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	},
}

// Revoke an API key
var revokeAPIKey = &cobra.Command{
	Use:   "revokeapikey",
	Short: "Revoke an API key for your Neon account",
	Run: func(cmd *cobra.Command, args []string) {
		keyID := cmd.Flag("keyid").Value.String()

		output, err := api_keys.RevokeAPIKey(keyID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	},
}

func init() {
	createAPIKey.Flags().String("keyname", "", "Name of the API key")
	revokeAPIKey.Flags().String("keyid", "", "ID of the API key to revoke")

	createAPIKey.MarkFlagRequired("keyname")
	revokeAPIKey.MarkFlagRequired("keyid")

	rootCmd.AddCommand(listAPIKeys)
	rootCmd.AddCommand(createAPIKey)
	rootCmd.AddCommand(revokeAPIKey)
}

func GetRootCommand() *cobra.Command {
	return rootCmd
}
