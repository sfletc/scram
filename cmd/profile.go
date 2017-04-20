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
	"github.com/spf13/cobra"
	"github.com/sfletc/scram2pkg"
	"strings"
	"strconv"
	"fmt"
	"os"
)

// profileCmd represents the profile command


var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Align reads of length l from 1 read file set to all sequences in a reference file",
	Long: `Align reads of length l from 1 read file set to all sequences in a reference file

For example:

scram2 profile -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -l 21,22,24 -o testAlign

`,
	Run: func(cmd *cobra.Command, args []string) {
		if readFileType != "cfa" && readFileType != "fa" && readFileType != "fq" {
			fmt.Println("\nCan't parse read file type " + readFileType)
			os.Exit(1)
		}
		a := scram2pkg.SeqLoad(strings.Split(fastaSet1, ","), readFileType,minLen, maxLen, minCount)
		c := scram2pkg.RefLoad(alignTo)
		for _, nt := range strings.Split(length, ",") {
			nt, _ := strconv.Atoi(nt)
			d := scram2pkg.AlignReads(a, c, nt)
			switch {
			case noSplit == false:
				e := scram2pkg.ProfileSplit(d, a)
				scram2pkg.ProfileToCsv(e, c, nt, outFilePrefix)
			default:
				e := scram2pkg.ProfileNoSplit(d, a)
				scram2pkg.ProfileToCsv(e, c, nt, outFilePrefix)

			}
		}
	},
}
func init() {
	RootCmd.AddCommand(profileCmd)

}
