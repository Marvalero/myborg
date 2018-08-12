package cmd

import (
	"fmt"
	"log"

	"github.com/Marvalero/myborg/myborg/secret"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Actions related to secrets",
	Long:  `Use actions related to secrets`,
}

var listSecretsCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets",
	Long:  `Lists secrets in your bucket`,
	Run: func(cmd *cobra.Command, args []string) {
		secret.List()
	},
}

var decryptSecretsCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt secrets",
	Long:  `decrypts secrets in your bucket`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypt secrets called")
	},
}

var encryptSecretsCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt secrets",
	Long:  `encrypts secrets and uploads it to your bucket`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatalln("Encrypt Secrets Usage: myborg secrets encrypt <name> <credentials>")
		}
		secret.Encrypt(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(secretsCmd)
	secretsCmd.AddCommand(encryptSecretsCmd)
	secretsCmd.AddCommand(decryptSecretsCmd)
	secretsCmd.AddCommand(listSecretsCmd)

	secretsCmd.Flags().StringVar(&cfgFile, "secrets-key-ring", "", "secrets key ring")
	secretsCmd.Flags().StringVar(&cfgFile, "secrets-key", "", "secrets key")
	secretsCmd.Flags().StringVar(&cfgFile, "secrets-bucket", "", "secrets bucket")
	viper.BindPFlag("secrets-key", secretsCmd.Flags().Lookup("secrets-key"))
	viper.BindPFlag("secrets-key-ring", secretsCmd.Flags().Lookup("secrets-key-ring"))
	viper.BindPFlag("secrets-bucket", secretsCmd.Flags().Lookup("secrets-bucket"))

}
