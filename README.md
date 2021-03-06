## SCRAM - Fast and simple small RNA read alignment

SCRAM (Small Complementary RNA Mapper) is a small RNA-focused index-free aligner that rapidly maps small RNA reads exactly to reference sequences, with an integrated visualization package generating interactive and publication-quality plots.

Details on installation and usage can be found [here](https://sfletc.github.io/scram/). 

To use the [Docker image](https://hub.docker.com/r/sfletcher/scram_docker/~/dockerfile/), see this [Gist](https://gist.github.com/sfletc/b70911d0de13bd4cde86f08b6ca32026). 


## Workflow

![workflow](https://cloud.githubusercontent.com/assets/5491692/25421042/4793d476-2a9e-11e7-9f41-9412f40f23f8.png)


## scram CLI options

### Profile alignment of 1 set of read files (likely biological replicates) to one or more reference sequences

```./scram profile ```

Align reads of length l from 1 read file set to all sequences in a reference file

For example:

```scram profile -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -l 21,22,24 -o testAlign```

Usage:
  ```scram profile [flags]```

#### Required Flags: ####

```-r, --alignTo```         : Path/to/FASTA reference file

```-1, --fastxSet1```       : Comma-separated path/to/read file set 1. GZIPped files must have .gz file extension

```-l, --length```          : Comma-separated read (sRNA) lengths to align

```-o, --outFilePrefix```   : Path/to/outfile prefix (_len.csv will be appended)



#### Optional Flags: ####

```-t, --readFileType```    : Read file type: cfa (collapsed FASTA), fa (FASTA), fq (FASTQ), clean (BGI clean.fa). (default "cfa")

```    --adapter```         : 3' adapter sequence to trim - FASTA & FASTQ only (default "nil")  

```    --maxLen```             : Maximum read length to include for RPMR normalization (default 32)

```    --minCount```         : Minimum read count for alignment and to include for RPMR normalization (default 1)

```    --minLen```             : Minimum read length to include for RPMR normalization (default 18)

```    --noSplit```                : Do not split alignment count for each read by the number of times it aligns


### Compare alignment of 2 sets of read files (likely biological replicates) to multiple reference sequences

```./scram compare```
 
Compare normalised alignment counts and standard errors for 2 read file sets

For example:

```scram compare -r ref.fa -1 seq1a.fa,seq1b.fa,seq1c.fa -2 seq2a.fa,seq2b.fa,seq2c.fa -l 21,22,24 -o testAlign```

Usage:
```scram compare [flags]```

#### Additional Required Flags: ####

```-2, --fastxSet2```       : Comma-separated path/to/read file set 2. GZIPped files must have .gz file extension

## scram_plot.py CLI options

### Profile plot ###

```%run scram_plot.py profile ```

#### Required Flags: ####

```-a, --alignment``` : sRNA alignment file prefix used by SCRAM profile (i.e. exclude _21.csv, _22.csv, _24.csv)

```-l, --length``` : Comma-separated list of sRNA lengths to plot. SCRAM alignment files must be available for each sRNA length

```-s, --search``` : Full header or substring of header. *Without flag, all headers will be plotted*

```-cutoff``` : Min. alignment RPMR from the most abundant profile (if multi) to generate plot
  
```-ylim``` :  +/- y axis limit
                          
```-win``` : Smoothing window size (default=auto)

```-pub``` : Remove all labels from profiles for editing for publication

```-png``` : Export plot/s as 300 dpi .png file/s

```-bin_reads``` : For plotting large profiles (i.e. chromosomes).  Assigns reads 10,000 bins prior to smoothing. X-axis shows bin, not reference position

### Compare plot ###

```%run scram_plot.py compare ```

#### Required Flags: ####

```-a, --alignment``` : sRNA alignment file prefix used by SCRAM compare (i.e. exclude _21.csv, _22.csv, _24.csv)

```-l, --length``` : Comma-separated list of sRNA lengths to plot. SCRAM alignment files must be available for each sRNA length


#### Optional Flags ####

```-plot_type``` : Bokeh plot type to display (log, log_error or all)

```-xlab``` : x label - corresponds to -s1 treatment in SCRAM arguments. Used to generate .png file name

```-ylab``` : y label - corresponds to -s2 treatment in SCRAM arguments. Used to generate .png file name

```-html``` : If not using Jupyter Notebook, output interactive plot to browser as save to .html

```-pub``` : Remove all labels from profiles for editing for publication

```-png``` : Export plot/s as 300 dpi .png file/s

## Worked Example

- For a detailed worked example of scram's capabilities, see the following link (which uses the scram_docker image):

[Jupyter notebook on nbviewer](https://nbviewer.jupyter.org/github/sfletc/scram_worked_example/blob/master/scram_demonstration.ipynb)

## Installation

### 1. Use the scram_docker image

- The scram aligner and scram_plot.py plotting script are installed, along with Jupyter notebook, on the minimal Miniconda base.
- You'll need docker installed. Ensure your project drive is shared and you've got a decent about of RAM (i.e. 8 GB+) available.

    1. Navigate to your project base directory. Your host project files (i.e. collapsed FASTA read and FASTA reference files in sub-directories) will be mounted.
    
        Bash shell
        ```
        docker run -it --rm  -v `pwd`:/work -p 8888:8888 sfletcher/scram_docker
        ```
        Windows PowerShell
        ```
        docker run -it --rm  -v ${PWD}:/work -p 8888:8888 sfletcher/scram_docker
        ```
    2. Copy generated link with token into your browser.  

    3. From a Jupyter notebook file, the scram aligner can be invoked by:
        ```
        !scram
        ```
        And the scram_plot.py script by:
        ```
        %run /scram_plot/scram_plot.py
        ```

### 2a. Download scram binary:

- Pre-compiled binaries are can be found at (*nix binaries may need to be made executable with ```chmod +x /path/to/binary```):

	[Mac OSX (64bit)](https://bitbucket.org/stevefl/scram/downloads/scram_osx)
	
	[Linux (64 bit)](https://bitbucket.org/stevefl/scram/downloads/scram_linux)
	
	[Windows (64 bit)](https://bitbucket.org/stevefl/scram/downloads/scram_win)

- Execute with the full binary name (e.g. scram_osx) rather than scram

### 2b. Or build from source:

- Go(lang) 1.8+ is required
    
    1. Install via ```go get```
    
        ```
        go get github.com/sfletc/scram github.com/sfletc/scramPkg github.com/spf13/cobra github.com/spf13/viper github.com/montanaflynn/stats
        ```
    2. Navigate to scram directory containing main.go (e.g. ```GOPATH/src/github.com/sfletc/scram/```)
        
        ```go install```
    3. scram will be in the ```GOPATH/bin``` directory
    
### 3. Install the scram_plot package and dependencies:

- Python 3.5+ is required 
        
    ```git clone https://github.com/sfletc/scram_plot.git```
    
    ```cd scram_plot```
    
    ```python setup.py install```
    
    
