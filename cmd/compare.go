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


var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare normalised alignment counts and standard errors for 2 read file sets",
	Long: `Compare normalised alignment counts and standard errors for 2 read file sets

For example:

scram2 compare -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -2 seq2a.fa,seq2b.fa,seq2c.fa -l 21,22,24 -o testAlign`,
	Run: func(cmd *cobra.Command, args []string) {
		if readFileType != "cfa" && readFileType != "fa" && readFileType != "fq" && readFileType != "clean" {
			fmt.Println("\nCan't parse read file type " + readFileType)
			os.Exit(1)
		}
		a := scram2pkg.SeqLoad(strings.Split(fastaSet1,","), readFileType,adapter,minLen, maxLen, minCount)
		b := scram2pkg.SeqLoad(strings.Split(fastaSet2,","), readFileType,adapter,minLen, maxLen, minCount)
		c := scram2pkg.RefLoad(alignTo)
		for _, nt := range strings.Split(length,",") { 
			nt,_ := strconv.Atoi(nt)
			switch {
			case noSplit == false:
				d := scram2pkg.AlignReads(a, c, nt)
				f := scram2pkg.CompareSplitCounts(d, a)
				e := scram2pkg.AlignReads(b, c, nt)
				g := scram2pkg.CompareSplitCounts(e, b)
				h := scram2pkg.Compare(f, g)
				scram2pkg.CompareToCsv(h, nt, outFilePrefix)
			default:
				d := scram2pkg.AlignReads(a, c, nt)
				f := scram2pkg.CompareNoSplitCounts(d, a)
				e := scram2pkg.AlignReads(b, c, nt)
				g := scram2pkg.CompareNoSplitCounts(e, b)
				h := scram2pkg.Compare(f, g)
				scram2pkg.CompareToCsv(h, nt, outFilePrefix)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(compareCmd)
	compareCmd.Flags().StringVarP(&fastaSet2, "fastxSet2", "2", "","comma-seperated path/to/read file set 2. GZIPped files must have .gz file extension")
}
