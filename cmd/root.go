package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/erindatkinson/pixawttr/internal/config"
	"github.com/erindatkinson/pixawttr/internal/convert"
	"github.com/erindatkinson/pixawttr/internal/pixabay"
	"github.com/erindatkinson/pixawttr/internal/wttrin"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	unit, output, query, location string
)

func rootRun(cmd *cobra.Command, args []string) {
	weatherFn, err := wttrin.GetWeatherImage(location, unitInFarenheight())
	if err != nil {
		hclog.L().Error("error pulling weather image", "error", err)
		return
	}
	defer os.Remove(weatherFn)

	if query == "" {
		query, err = wttrin.GetWeather(location, unitInFarenheight())
		if err != nil {
			hclog.L().Error("error pulling weather text", "error", err)
			return
		}
	}
	bgFn, err := pixabay.GetImage(viper.GetString("PixabayAPIKey"), query)
	if err != nil {
		hclog.L().Error("error pulling bg image", "error", err)
		return
	}
	defer os.Remove(bgFn)

	err = convert.Merge(bgFn, weatherFn, output)
	if err != nil {
		hclog.L().Error("error merging images", "error", err)
		return
	}

}

func unitInFarenheight() bool {
	switch strings.ToLower(unit) {
	case "f":
		return true
	case "c":
		return false
	default:
		hclog.L().Error("unit must be one of 'c' or 'f'", "unit", unit)
		os.Exit(1)
	}
	return false
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pixawttr [options]",
	Short: "A binary to make pretty weather updates",
	Long:  ``,
	Run:   rootRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.Flags().StringVarP(&unit, "unit", "u", "c", "unit for temperature [c|f]")
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "query for picture")
	rootCmd.Flags().StringVarP(&output, "output", "o", "outfile.png", "file to output to")
	rootCmd.Flags().StringVarP(&location, "location", "l", "Golden", "The location for weather")
}
