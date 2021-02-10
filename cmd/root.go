package cmd

import (
	"fmt"
	"os"

	"github.com/erindatkinson/pixawttr/internal/convert"
	"github.com/erindatkinson/pixawttr/internal/pixabay"
	"github.com/erindatkinson/pixawttr/internal/wttrin"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	useMetric bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pixawttr location [outfile]",
	Short: "A binary to make pretty weather updates",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		var outFile string
		switch numArgs := len(cmd.Flags().Args()); {
		case numArgs == 1:
			outFile = "outFile.png"
		case numArgs == 2:
			outFile = cmd.Flags().Args()[1]
		default:
			cmd.Usage()
			return
		}
		text, err := wttrin.GetWeather(cmd.Flags().Args()[0])
		if err != nil {
			hclog.Default().Error("couldn't get weather", "error", err)
			return
		}

		bgImage, err := pixabay.GetImage(viper.GetString("PixabayAPIKey"), text)
		if err != nil {
			hclog.Default().Error("couldn't download background", "error", err)
			return
		}

		weatherImg, err := wttrin.GetWeatherImage(cmd.Flags().Args()[0])
		if err != nil {
			hclog.Default().Error("couldn't download forecast image", "error", err)
			return
		}

		err = convert.Merge(bgImage, weatherImg, "outFile.png")
		if err != nil {
			hclog.Default().Error("couldn't merge images")
		}

		hclog.Default().Info("Merged image is at", "location", outFile)

		if err = os.Remove(bgImage); err != nil {
			hclog.Default().Error("Didn't clean up interim image", "image", bgImage, "error", err)
		}
		if err = os.Remove(weatherImg); err != nil {
			hclog.Default().Error("Didn't clean up interim image", "image", weatherImg, "error", err)
		}

	},
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pixawttr.yaml)")
	rootCmd.Flags().BoolVarP(&useMetric, "metric", "m", false, "use metric")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pixawttr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pixawttr")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
