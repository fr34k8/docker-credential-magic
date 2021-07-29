package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/docker-credential-magic/docker-credential-magic/pkg/magician"
)

type mutateSettings struct {
	Tag string
	IncludeHelpers []string
}

func main() {
	var mutate mutateSettings

	rootCmd := &cobra.Command{
		Use:   "docker-credential-magician",
		Short: "Augment images with various credential helpers (including magic)",
	}

	mutateCmd := &cobra.Command{
		Use: "mutate",
		Short: "Augment an image with one or more credential helpers",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ref := args[0]
			var opts []magician.MutateOption
			if tag := mutate.Tag; tag != "" {
				opts = append(opts, magician.MutateOptWithTag(tag))
			}
			if len(mutate.IncludeHelpers) > 0 {
				opts = append(opts, magician.MutateOptWithHelpers(mutate.IncludeHelpers))
			}
			return magician.Mutate(ref, opts...)
		},
	}
	mutateCmd.Flags().StringVarP(&mutate.Tag, "tag", "t", "", "push to custom location")
	mutateCmd.Flags().StringArrayVarP(&mutate.IncludeHelpers, "include", "i",
		[]string{}, "custom helpers to include")

	rootCmd.AddCommand(mutateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
}
