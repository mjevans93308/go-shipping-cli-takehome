package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mjevans93308/platformscience/pkg/algs"
	"github.com/mjevans93308/platformscience/pkg/files"
	"github.com/mjevans93308/platformscience/pkg/models"
	"github.com/mjevans93308/platformscience/util/localctx"
	"github.com/mjevans93308/platformscience/util/logs"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func Execute() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "platsci",
		Short: "A CLI tool to process newline-separated files for Platform Science",
	}

	var n string
	var a string
	var o string

	var processCmd = &cobra.Command{
		Use:   "process",
		Short: "Process a pair of files containing names and addresses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runner(n, a, o)
		},
	}

	processCmd.Flags().StringVarP(&a, "addresses", "a", "", "Path to the file containing a newline-delimited list of addresses")
	processCmd.Flags().StringVarP(&n, "names", "n", "", "Path to the file containing a newline-delimited list of names")
	processCmd.Flags().StringVarP(&o, "output", "o", "", "Name of file to output save output into, not required")
	processCmd.MarkFlagRequired("addresses")
	processCmd.MarkFlagRequired("names")

	rootCmd.AddCommand(processCmd)

	return rootCmd
}

func processConcurrently(ctx *localctx.Localctx, g *errgroup.Group, nFile string, aFile string) (*models.FileContent, error) {
	var names []string
	var addresses []string

	// process files concurrently in case of very large files
	g.Go(func() error {
		lnames, err := files.ProcessFile(ctx, nFile)
		if err != nil {
			ctx.Logger.Errorf("Error encountered while processing name file: %s\n%s", nFile, err)
			return err
		}
		names = lnames
		return nil
	})
	g.Go(func() error {
		laddresses, err := files.ProcessFile(ctx, aFile)
		if err != nil {
			ctx.Logger.Errorf("Error encountered while processing address file: %s\n%s", aFile, err)
			return err
		}
		addresses = laddresses
		return nil
	})

	return &models.FileContent{
		Names:     names,
		Addresses: addresses,
	}, g.Wait()
}

func runner(n, a, output string) error {
	slog := logs.NewSlog()
	ctx := context.Background()

	// using errgroup over waitgroup for added safety
	// https://www.storj.io/blog/production-concurrency
	g, ctx := errgroup.WithContext(ctx)
	myCtx := localctx.NewLocalCtx(&ctx, slog)
	content, err := processConcurrently(myCtx, g, n, a)
	if err != nil {
		return err
	}
	// ensure we have a matching number of names and addresses
	if len(content.Addresses) != len(content.Names) {
		slog.Errorf("Number of supplied names does not match number of supplied address: %d != %d", len(content.Addresses), len(content.Names))
		return errors.New("Number of supplied names does not match number of supplied address")
	}
	result := algs.CalculateSS(myCtx, content)
	return handleOutput(myCtx, output, result)
}

func handleOutput(ctx *localctx.Localctx, output string, result []models.NameAddressSS) error {
	var text []string
	if output != "" {
		if strings.HasSuffix(output, ".txt") {
			i := 1
			for _, r := range result {
				text = append(text, fmt.Sprintf("%d. %s: %s - %f", i, r.Name, r.Address, r.SS))
				i++
			}
			input := strings.Join(text, "\n")
			err := os.WriteFile(output, []byte(input), 0644)
			if err != nil {
				ctx.Logger.Errorf("Error writing to file: ", err)
				return err
			}
		} else {
			ctx.Logger.Error("File output only supports `.txt` file types for now")
			return errors.New("File output only supports `.txt` file types for now")
		}
	} else {
		fmt.Println("Data:")
		i := 1
		for _, r := range result {
			fmt.Printf("%d. %s: %s - %f\n", i, r.Name, r.Address, r.SS)
			i++
		}
	}
	return nil
}
