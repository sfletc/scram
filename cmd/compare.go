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
	"time"
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
		t0:=time.Now()
		fmt.Println("\nLoading reads\n")
		a := scram2pkg.SeqLoad(strings.Split(fastaSet1,","), readFileType,adapter,minLen, maxLen, minCount, noNorm)
		b := scram2pkg.SeqLoad(strings.Split(fastaSet2,","), readFileType,adapter,minLen, maxLen, minCount, noNorm)

		fmt.Println("\nLoading reference\n")

		switch{
		case mir == false:
			c := scram2pkg.RefLoad(alignTo)

			for _, nt := range strings.Split(length,",") {
				nt,_ := strconv.Atoi(nt)
				fmt.Printf("\nAligning %v nt reads\n", nt)
				d := scram2pkg.AlignReads(a, c, nt)
				e := scram2pkg.AlignReads(b, c, nt)
				switch {
				case noSplit == false:

					f := scram2pkg.CompareSplitCounts(d, a)

					g := scram2pkg.CompareSplitCounts(e, b)
					h := scram2pkg.Compare(f, g)
					scram2pkg.CompareToCsv(h, nt, outFilePrefix)
				default:

					f := scram2pkg.CompareNoSplitCounts(d, a)

					g := scram2pkg.CompareNoSplitCounts(e, b)
					h := scram2pkg.Compare(f, g)
					scram2pkg.CompareToCsv(h, nt, outFilePrefix)
				}
			}
		default:
			c := scram2pkg.MirLoad(alignTo)
			d := scram2pkg.AlignMirnas(a, c)
			e := scram2pkg.AlignMirnas(b, c)
			switch {
			case noSplit ==false:
				f := scram2pkg.MirnaCompare(d,e,false)
				scram2pkg.CompareToCsv(f,0,outFilePrefix)
			default:
				f := scram2pkg.MirnaCompare(d,e,true)
				scram2pkg.CompareToCsv(f,0,outFilePrefix)
			}
		}

		t1 := time.Now()
		fmt.Printf("\nAlignment complete.  Total time taken = %s\n",t1.Sub(t0))
	},
}

func init() {
	RootCmd.AddCommand(compareCmd)
	compareCmd.Flags().StringVarP(&fastaSet2, "fastxSet2", "2", "","comma-separated path/to/read file set 2. GZIPped files must have .gz file extension")
	compareCmd.Flags().BoolVar(&mir, "mir", false, "Exact match reads to mature miRNAs")
}
