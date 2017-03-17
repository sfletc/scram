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
)


var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare normalised alignment counts and standard errors for 2 read file sets",
	Long: `Compare normalised alignment counts and standard errors for 2 read file sets

For example:

scram2 compare -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -2 seq2a.fa,seq2b.fa,seq2c.fa -l 21,22,24 -o testAlign`,
	Run: func(cmd *cobra.Command, args []string) {
		a := scram2pkg.SeqLoad(strings.Split(fastaSet1,","), minLen, maxLen, minCount)
		b := scram2pkg.SeqLoad(strings.Split(fastaSet2,","), minLen, maxLen, minCount)
		c := scram2pkg.RefLoad(alignTo)
		for _, nt := range strings.Split(len,",") {
			nt,_ := strconv.Atoi(nt)
			switch {
			case noSplit == false:
				d := scram2pkg.AlignReads(a, c, nt)
				f := scram2pkg.CdpSplitCounts(d, a)
				e := scram2pkg.AlignReads(b, c, nt)
				g := scram2pkg.CdpSplitCounts(e, b)
				h := scram2pkg.Cdp(f, g)
				scram2pkg.CdpToCsv(h, nt, outFilePrefix)
			default:
				d := scram2pkg.AlignReads(a, c, nt)
				f := scram2pkg.CdpNoSplitCounts(d, a)
				e := scram2pkg.AlignReads(b, c, nt)
				g := scram2pkg.CdpNoSplitCounts(e, b)
				h := scram2pkg.Cdp(f, g)
				scram2pkg.CdpToCsv(h, nt, outFilePrefix)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(compareCmd)
	//compareCmd.Flags().StringVarP(&alignTo, "alignTo", "r", "","path/to/FASTA reference file")
	//compareCmd.Flags().StringVarP(&fastaSet1, "fastaSet1", "1", "","comma-seperated path/to/collapsed FASTA file set 1")
	compareCmd.Flags().StringVarP(&fastaSet2, "fastaSet2", "2", "","comma-seperated path/to/collapsed FASTA file set 2")
	//compareCmd.Flags().StringVarP(&len, "len", "l", "","comma-seperated read (sRNA) lengths to align")
	//compareCmd.Flags().StringVarP(&outFilePrefix, "outFilePrefix", "o", "","path/to/outfile prefix (len.csv will be appended)")
	//compareCmd.Flags().BoolVar(&noSplit, "noSplit", false, "Do not split alignment count for each read by the number of times it aligns")
	//compareCmd.Flags().Lookup("noSplit").NoOptDefVal="true"
	//compareCmd.Flags().IntVar(&minLen, "minLen", 18, "Minimum read length to include for RPMR normalization")
	//compareCmd.Flags().IntVar(&maxLen, "maxLen", 32, "Maximum read length to include for RPMR normalization")
	//compareCmd.Flags().Float64Var(&minCount, "minCount", 1.0, "Minimum read count for alignment and to include for RPMR normalization")

}
