package cmd

import (
	"os"

	cmd "github.com/can3p/kleiner/shared/cmd/cobra"
	"github.com/can3p/kleiner/shared/published"
	"github.com/can3p/qrcoder/generated/buildinfo"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "qrcoder",
		Short: "Generate a qr code for a give value",
		RunE: func(cmd *cobra.Command, args []string) error {
			value, err := cmd.Flags().GetString("value")

			if err != nil {
				return err
			}

			size, err := cmd.Flags().GetInt("size")

			if err != nil {
				return err
			}

			output, err := cmd.Flags().GetString("out")

			if err != nil {
				return err
			}

			if output == "-" {
				var png []byte
				png, err := qrcode.Encode(value, qrcode.Medium, size)

				if err != nil {
					return err
				}

				_, err = os.Stdout.Write(png)

				if err != nil {
					return err
				}

				return nil
			}

			return qrcode.WriteFile(value, qrcode.Medium, size, output)
		},
	}

	rootCmd.Flags().String("value", "", "value to encode")
	rootCmd.Flags().Int("size", 256, "QR Code size")
	rootCmd.Flags().String("out", "output.png", "Output file, - for std out")

	return rootCmd
}

var rootCmd = RootCommand()

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	info := buildinfo.Info()

	cmd.Setup(info, rootCmd)
	published.MaybeNotifyAboutNewVersion(info)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
