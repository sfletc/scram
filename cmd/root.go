//The MIT License (MIT)
//
//Copyright © 2017 Stephen Fletcher
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)
const version =  "0.1.2"
var cfgFile string
var fastaSet1 string
var fastaSet2 string
var length string
var alignTo string
var outFilePrefix string
var noSplit bool
var minLen int
var maxLen int
var minCount float64

var RootCmd = &cobra.Command{
	Use:   "scram2",
	Short: "The ultra-fast siRNA aligner v"+version,
	Long: `The ultra-fast siRNA aligner v`+version,

}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&alignTo, "alignTo", "r", "","path/to/FASTA reference file")
	RootCmd.PersistentFlags().StringVarP(&fastaSet1, "fastaSet1", "1", "","comma-seperated path/to/collapsed FASTA file set 1")
	RootCmd.PersistentFlags().StringVarP(&length, "length", "l", "","comma-seperated read (sRNA) lengths to align")
	RootCmd.PersistentFlags().StringVarP(&outFilePrefix, "outFilePrefix", "o", "","path/to/outfile prefix (len.csv will be appended)")
	RootCmd.PersistentFlags().BoolVar(&noSplit, "noSplit", false, "Do not split alignment count for each read by the number of times it aligns")
	RootCmd.PersistentFlags().Lookup("noSplit").NoOptDefVal="true"
	RootCmd.PersistentFlags().IntVar(&minLen, "minLen", 18, "Minimum read length to include for RPMR normalization")
	RootCmd.PersistentFlags().IntVar(&maxLen, "maxLen", 32, "Maximum read length to include for RPMR normalization")
	RootCmd.PersistentFlags().Float64Var(&minCount, "minCount", 1.0, "Minimum read count for alignment and to include for RPMR normalization")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".scram2") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
