package cmd

import (
	"github.com/rbutcher/aoc2024/internal/solution"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	verboseFlagKey = "verbose"
	dayFlagKey     = "day"
	partFlagKey    = "part"
)

var rootCmd = &cobra.Command{
	Use:   "Run the Advent of Code problem of the day",
	Short: "Run the Advent of Code problem of the day",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel := zerolog.InfoLevel
		if viper.GetBool(verboseFlagKey) {
			logLevel = zerolog.DebugLevel
		}

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.SetGlobalLevel(logLevel)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		solutions := []solution.Constructor{
			solution.NewDay1,
			solution.NewDay2,
			solution.NewDay3,
			solution.NewDay4,
			solution.NewDay5,
		}

		part := viper.GetInt(partFlagKey)
		day := viper.GetInt(dayFlagKey)
		l := log.With().Int(dayFlagKey, day).Int(partFlagKey, part).Logger()
		if day <= 0 || day > len(solutions) {
			l.Fatal().Msgf("selected day is not available")
		}

		l.Debug().Msg("starting solution")
		slnCtor := solutions[day-1]
		sln := slnCtor()

		var result string
		var err error
		if part == 2 {
			result, err = sln.Part2()
		} else {
			result, err = sln.Part1()
		}
		l.Debug().Msg("done with solution")

		if err != nil {
			return err
		}

		l.Info().Str("result", result).Send()
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolP(verboseFlagKey, "v", false, "Enable verbose logging")
	_ = viper.BindPFlag(verboseFlagKey, rootCmd.Flags().Lookup(verboseFlagKey))

	rootCmd.Flags().IntP(dayFlagKey, "d", 1, "Select which day to run")
	_ = viper.BindPFlag(dayFlagKey, rootCmd.Flags().Lookup(dayFlagKey))

	rootCmd.Flags().IntP(partFlagKey, "p", 1, "Select which part to run")
	_ = viper.BindPFlag(partFlagKey, rootCmd.Flags().Lookup(partFlagKey))
}

func Execute() error {
	return rootCmd.Execute()
}
