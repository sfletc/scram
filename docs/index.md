<img src="https://user-images.githubusercontent.com/5491692/30258263-22944448-96fb-11e7-91ed-cefc31613523.png" width="200">

The SCRAM pipeline is developed by Stephen Fletcher at Bernie Carroll's Laboratory, School of Chemistry and Molecular Biosciences, University of Queensland, Australia

----
[1. Installation](#scram-pipeline-installation)

[2. Workflow](#scram-workflow)

[3. Worked examples of the SCRAM pipeline in Jupyter Notebook](#worked-example)

[4. SCRAM aligner command line options](#scram-cli-options)

[5. SCRAM plotting module command line options](#scram-plotting-module-cli-options)

----

## SCRAM Pipeline Installation

### 1. Use the scram_docker image

### What is Docker (and why is it useful for bioinformatics)?

Long story short - it takes the pain out of installing bioinformatics software if someone's already done it for you.  

Docker allows you to download images and run containers on a Windows, Mac or Linux PC.  All the software for a bioinformatics workflow can be loaded onto an image and tested before deployment.   It doesn't matter that the software is designed for Linux, and you're using a Windows PC - it'll run!!  

But what is an image or a container?  If you take a snapshot of your PC's hard-drive, that's essentially an image.  It has an OS, whatever software you have installed, and some data files.  When you 'run' an image, a container is generated.  This is similar to a running PC - it can be interacted with, and its installed software used.  Instead of being a discrete entity like a PC, the container is virtualized.  One image can be used to generate multiple containers, which can run similtaneously on a single PC.  Often containers are deleted once the application they are performing is complete.      

Redhat has a nice [outline](https://www.redhat.com/en/containers/what-is-docker):

> Container tools, including Docker, provide an image-based deployment model. This makes it easy to share an application, or set of services, with all of their dependencies across multiple environments. Docker also automates deploying the application (or combined sets of processes that make up an app) inside this container environment.

Rather than installing and troubleshooting software for QC, adapter trimming, read alignment, visualization and electronic lab book keeping, a single command can download a pre-configured image, and a single command can start the container.    

Why is Docker great for bioinformatics?  It makes installation of complex pipelines simple and hassle free (not that the SCRAM pipeline is overly complex).  It also supports reproducible science - the same software versions can be used across multiple systems, ensuring the same outputs for a given set of inputs.  For a publication, not only can the read files be made available online in the SRA for example, but the entire pipeline that generated the data for the paper can also be stored in the cloud, allowing anyone to reproduce the results, spot process errors, and suggest improvements.  

The SCRAM docker image has the following packages installed (among others).  Remember to cite the package authors if you use them!

1. [The SCRAM aligner and plotter](https://sfletc.github.io/scram/)
2. [Jupyter Notebook](https://jupyter.org/)
3. [FastQC](https://www.bioinformatics.babraham.ac.uk/projects/fastqc/)
4. [FastX-toolkit](http://hannonlab.cshl.edu/fastx_toolkit/)
5. [Blast+](https://blast.ncbi.nlm.nih.gov/Blast.cgi?PAGE_TYPE=BlastDocs&DOC_TYPE=Download)

The SCRAM docker file (which is used to build the image) is [here](https://github.com/sfletc/scram_docker/blob/master/Dockerfile) if you want to take a look.

### Prerequisites

A mid-range PC or up - laptop or desktop.  Ideally a minimum of 8GB RAM, though 16GB+ is better for larger projects.  Remember, unless it's a Linux machine, it's running both its own OS and the SCRAM Docker container with associated OS at the same time.

Easiest is a Linux machine (e.g. running Ubuntu).  Next are Windows 10 Pro or similar machines, and Apple PCs running OS X (1-click installs).  Windows 7 and (I think) Win 10 Home machines require a bit more playing around (Docker Toolbox and Virtual Box instead of Hyper-V)

### Install Docker CE

#### Ubuntu 

Via the Terminal:

``` sudo apt-get install docker.io```

You'll need to use ```sudo``` or [follow this guide](https://askubuntu.com/questions/477551/how-can-i-use-docker-without-sudo)

#### OS X (Apple Mac)

Full instructions are [here](https://store.docker.com/editions/community/docker-ce-desktop-mac)

#### Windows 10 (Pro or Education)

Full instructions are [here](https://store.docker.com/editions/community/docker-ce-desktop-windows)

#### Windows 10 (Home) or Windows 7/8

Unfortunately it's a bit more work  - full instructions are [here](https://docs.docker.com/toolbox/toolbox_install_windows/)

### Start Docker for the first time (Windows 10 and Mac - skip for Linux)

After starting docker, right click on the Docker icon in the taskbar / dock, and click ```settings```.  You'll probably want to un-tick the ```Start Docker when you log in``` check box.  

Click on the ```Shared Drives``` tab (left), and ensure the drive with your project data is shared.  Sometimes antivirus products (like Kaspersky) can interfere, so their settings may need altering.  

Next, click on the ```Advanced``` tab and give Docker sufficient CPUs and memory (RAM).  All CPUs and approximately 4GB memory *remaining* for the host OS should be OK, but check how much RAM the host OS and background processes are actually using if this is an issue.

Multiple Docker images (and containers if you retain them) can take up a lot of disk space, so ensure that there is sufficient spare.  

### Download the SCRAM docker image

With Docker running in the background, open a terminal (Mac and Linux), Powershell (Windows 10 Pro), or the Docker terminal (Windows 10 Home, Windows7/8) and enter:

```docker pull sfletcher/scram_docker```

```sudo``` may be needed for Linux.  A decent internet connection is required, as there is a fair bit to download.  If the software is updated, it's likely only a portion of this will have to be (automatically) downloaded again.  

### 2a. Download scram binary:

Pre-compiled binaries are can be found [here](https://github.com/sfletc/scram/releases) (*nix binaries may need to be made executable with ```chmod +x /path/to/binary```):


Execute with the full binary name (e.g. scram_osx) rather than scram

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

## SCRAM Workflow

![workflow](https://user-images.githubusercontent.com/5491692/30198690-57a70788-94b2-11e7-82ba-a36280e2310b.png)

## Worked Example

For a detailed worked example of the SCRAM pipeline's capabilities, see the following link (which uses the scram_docker image):

[Jupyter notebook on nbviewer](https://nbviewer.jupyter.org/github/sfletc/scram_worked_example/blob/master/scram_demonstration.ipynb)

## SCRAM aligner CLI options

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

## SCRAM plotting module CLI options

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



    
    
