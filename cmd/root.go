package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-loggi/helper"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"time"
)

var filename string

type config struct {
	Interval int64    `yaml:"interval"`
	Strings  []string `yaml:"strings"`
}

func (c *config) getConf(filename string) *config {

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	if len(c.Strings) == 0 {
		log.Fatalf("Provide some strings!")
	}

	return c
}
func (c *config) GetRandomString() string {
	return c.Strings[rand.Intn(len(c.Strings))]
}

var rootCmd = &cobra.Command{
	Use:   "go-loggi",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.LoggerInit()
		var c config
		c.getConf(filename)

		log.Infof("Interval: %d seconds", c.Interval)
		log.Infof("filename: %s", filename)
		for range time.Tick(time.Second * time.Duration(c.Interval)) {
			go func() {
				log.Info(c.GetRandomString())

			}()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&filename, "filename", "f", "example.yaml", "pass filename")
}
