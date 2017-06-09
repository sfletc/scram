## Workflow

![workflow](https://cloud.githubusercontent.com/assets/5491692/25421042/4793d476-2a9e-11e7-9f41-9412f40f23f8.png)


## scram2 CLI options

### Profile alignment of 1 set of read files (likely biological replicates) to one or more reference sequences

```./scram2 profile -h```

Align reads of length l from 1 read file set to all sequences in a reference file

For example:

```scram2 profile -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -l 21,22,24 -o testAlign```

Usage:
  ```scram2 profile [flags]```

Required Flags:

```-r, --alignTo```         : Path/to/FASTA reference file

```-1, --fastxSet1```       : Comma-seperated path/to/read file set 1. GZIPped files must have .gz file extension

```-l, --length```          : Comma-seperated read (sRNA) lengths to align

```-o, --outFilePrefix```   : Path/to/outfile prefix (_len.csv will be appended)



Optional Flags:

```-t, --readFileType```    : Read file type: cfa (collapsed FASTA), fa (FASTA), fq (FASTQ), clean (BGI clean.fa). (default "cfa")

```    --adapter```         : 3' adapter sequence to trim - FASTA & FASTQ only (default "nil")  

```    --maxLen```             : Maximum read length to include for RPMR normalization (default 32)

```    --minCount```         : Minimum read count for alignment and to include for RPMR normalization (default 1)

```    --minLen```             : Minimum read length to include for RPMR normalization (default 18)

```    --noSplit```                : Do not split alignment count for each read by the number of times it aligns






### Compare alignment of 2 sets of read files (likely biological replicates) to multiple reference sequences

```
./scram2 compare -h
Compare normalised alignment counts and standard errors for 2 read file sets

For example:

scram2 compare -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -2 seq2a.fa,seq2b.fa,seq2c.fa -l 21,22,24 -o testAlign

Usage:
  scram2 compare [flags]

Flags:
  -2, --fastxSet2 string   comma-separated path/to/read file set 2. GZIPped files must have .gz file extension

Global Flags:
      --adapter string         3' adapter sequence to trim - FASTA & FASTQ only (default "nil")
  -r, --alignTo string         path/to/FASTA reference file
  -1, --fastxSet1 string       comma-seperated path/to/read file set 1. GZIPped files must have .gz file extension
  -l, --length string          comma-seperated read (sRNA) lengths to align
      --maxLen int             Maximum read length to include for RPMR normalization (default 32)
      --minCount float         Minimum read count for alignment and to include for RPMR normalization (default 1)
      --minLen int             Minimum read length to include for RPMR normalization (default 18)
      --noSplit                Do not split alignment count for each read by the number of times it aligns
  -o, --outFilePrefix string   path/to/outfile prefix (len.csv will be appended)
  -t, --readFileType string    Read file type: cfa (collapsed FASTA), fa (FASTA), fq (FASTQ), clean (BGI clean.fa). (default "cfa")
```



## Worked Example

- For a detailed worked example of scram2's capabilities, see the following link (which uses the scram2_docker image):

[Jupyter notebook on nbviewer](https://nbviewer.jupyter.org/github/sfletc/scram2_worked_example/blob/master/scram2_demonstration.ipynb)

## Installation

### 1. Use the scram2_docker image

- The scram2 aligner and scram2_plot.py plotting script are installed, along with Jupyter notebook, on the minimal Miniconda base.
- You'll need docker installed. Ensure your project drive is shared and you've got a decent about of RAM (i.e. 8 GB+) available.

    1. Navigate to your project base directory. Your host project files (i.e. collapsed FASTA read and FASTA reference files in sub-directories) will be mounted.
    
        Bash shell
        ```
        docker run -it --rm  -v `pwd`:/work -p 8888:8888 sfletcher/scram2_docker
        ```
        Windows PowerShell
        ```
        docker run -it --rm  -v ${PWD}:/work -p 8888:8888 sfletcher/scram2_docker
        ```
    2. Copy generated link with token into your browser.  

    3. From a Jupyter notebook file, the scram2 aligner can be invoked by:
        ```
        !scram2
        ```
        And the scram2_plot.py script by:
        ```
        %run /scram2_plot/scram2_plot.py
        ```

### 2a. Download scram2 binary:

- Pre-compiled binaries are can be found at (*nix binaries may need to be made executable with ```chmod +x /path/to/binary```):

	[Mac OSX (64bit)](https://bitbucket.org/stevefl/scram2/downloads/scram2osx)
	
	[Linux (64 bit)](https://bitbucket.org/stevefl/scram2/downloads/scram2linux)
	
	[Windows (64 bit)](https://bitbucket.org/stevefl/scram2/downloads/scram2win)

- Execute with the full binary name (e.g. scram2osx) rather than scram2

### 2b. Or build from source:

- Go(lang) 1.8+ is required
    
    1. Install via ```go get```
    
        ```
        go get github.com/sfletc/scram2 github.com/sfletc/scram2pkg github.com/spf13/cobra github.com/spf13/viper github.com/montanaflynn/stats
        ```
    2. Navigate to scram2 directory containing main.go (e.g. ```GOPATH/src/github.com/sfletc/scram2/```)
        
        ```go install```
    3. scram2 will be in the ```GOPATH/bin``` directory
    
### 3. Install the scram2_plot package and dependencies:

- Python 3.5+ is required 
        
    ```git clone https://github.com/sfletc/scram2_plot.git```
    
    ```cd scram2_plot```
    
    ```python setup.py install```
    
    
