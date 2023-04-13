/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
  "github.com/amanoese/kmeango/pkg"
)


var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "kmeango",
  Short: "A brief description of your application",
  Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Run: func(cmd *cobra.Command, args []string) {
    var lines = *read_input()
    clumnsNum := len(strings.Split(lines[0]," "))
    datas := make([]pkg.Point,len(lines))

    for i, line := range lines {
      int_values := make([]int,clumnsNum)
      line_words := strings.Split(line," ")
      for w_i, word := range line_words {
        int_values[w_i],_ = strconv.Atoi(word)
      }
      p := pkg.Point{ Values: int_values }
      datas[i] = p
    }
		k, err := cmd.Flags().GetInt("kcotegory")
		if err != nil {
			fmt.Println(err)
		}
    results := pkg.Calc(datas,k)
    fmt.Println("Center : Value")
    for _, point := range results {
      category := ""
      values := ""
      for _,str := range point.CategoryValue {
        category = category + " " + strconv.Itoa(str)
      }
      for _,str := range point.Values {
        values = values + " " + strconv.Itoa(str)
      }
      fmt.Println(fmt.Sprintf("%6s", category),":",values)
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

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kmeango.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().IntP("kcotegory", "k", 3, "kmean cluster category number")
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

    // Search config in home directory with name ".kmeango" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".kmeango")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

func read_input() *[]string{
    sc:=bufio.NewScanner(os.Stdin)
    var lines []string
    for sc.Scan() {
        lines = append(lines,sc.Text())
    }
    return &lines

}
