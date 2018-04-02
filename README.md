# go-fasta
## Command line based fasta housekeeping (merge, split, sort and summarize), also queries NCBI to obtain fasta files with sequence corresponding to given accession numbers.

# UNDER CONSTRUCTION - library works but I'm still working on documentation!
Still TO DO:
- test all of the cmd line functionality, make sure parallel parts work as expected
- format all the files with go fmt
- make the travis.yml file and have the build pass
- compiled version for download? /makefile?
- summary file throw warning if there are non ATGCN bp in the sequences


## What is in the repository?

## 1. go-fasta program 
This program is controlled through a series of user specified command line options which allow for:
* fasta file merger
* fasta file splitting
* fasta file creation from NCBI query
* fasta file summary
* fasta file sorting

### -m merge 
Takes a list of comma delimited file names (or a text file with a list of fasta file names) and merge them and write to -f or default name

Merge Fastas. The -m flage takes a comma delimited list of fasta filenames to be merged. The final fasta will contain the sequences in the order of the .fasta inputs.
You an also pass in a .txt filename which contains a list of filnames (all names specified on seprate lines).
Use in conjunction with the -f flag to alter the output file name (default: output.fasta).

```
EXAMPLE CODE
```

### -n ncbi 
Query NCBI. Takes a comma delimited list of unique NCBI IDs. The .fasta files associated with the accession IDs will be downloaded and saved to a .fasta file.
You an also pass in a .txt filename which contains a list of IDs (all specified on seprate lines). 
Use in conjunction with the -f flag to alter the output file name. Note: this will run significantly faster if not called in conjunction with the -summary flag, as this requires the data to be stored in memory instead of written directly to the file.

```
EXAMPLE CODE
```

### -a alphabetize
Alphabetize fasta. pass this flag name in conjunction with a -f flag. sequences in the -f specified file will be sorted alphabetically by sequence name.

```
EXAMPLE CODE
```

### -split 
Split Fasta. Pass this flag name in conjunction with a -f flag.
The Sequences in the -f specified file will be split into a set of fasta files, one for each sequence in the file.

```
EXAMPLE CODE
```

### -summary
Make a summary file of output. Pass this flag and a summary file will be constructed which
gives the following information for each sequence in the fasta produced: 
	sequence name	sequence length	percent gc content
IMPORTANT NOTE: summary is designed for use with nucleotide based fasta files, if you call it on a protein sequence fasta file the gc content column will be nonsense!

```
EXAMPLE CODE
```


### -f
File name. A .fasta or .txt filename. For use with -m -n -a -split and -summary flags to specify an output name.
If both a fasta and summary are needed, just passed a .fasta name and it will produce a summary file with the same name and a .txt extension.

```
EXAMPLE CODE
```



## 2. The fasta package 
* A copy of the fasta package's documentation is shown below. It can also be obtained using `godoc`, follwed by the path to the package (i.e. `godoc ./fasta`)

The fasta package is designed to provide a suite of functions for reading, writing and manipulating Fasta sequence files. This library can be imported into other go projects to allow for simplified use and creation of Fasta files. The Query functions also provide the ability to retrieve new sequence data in fasta format from the National Center for Biotechnology Information (NCBI) databases by providing a slice of unique sequence ids.
Importing this library provides the following specalized data structures and methods:


#### type `Seq`
* a struct with two fields, Name and Sequence to represent the two parts of a fasta file entry.

#### type `Fasta` 
* a slice of Seq types []Seq
* This structure represents a fasta file for the library, it is a set of Seq structures.
* The library's input/output functionality allows for efficient reading and writing of files in standard fasta format

#### func Read(filename string)
* This function takes one argument, a string specifying the name of a fasta file.
* The function returns an object of the type Fasta

#### func (fa \*Fasta) AddItem(item Seq)
* This method can be called on an existing Fasta type. It takes one argument, specifying a new Seq structure to be added to the Fasta (this is useful for merging multiple Fasta files or adding newly obtained data to an existing Fasta)
* The Fasta will be modified in place (this is a pointer reciever method).

#### func (fa Fasta) Write(filename)
* This method can be called on an existing Fasta type. It takes one argument, a filename (path optional) to which the Fasta will be written.
* The output is in standard Fasta format, with a header line prefaced by a '>' character and a sequence section with 60 characters of sequence per line.

#### func (fa \*Fasta) Sort()
* This method can be called on a Fasta type instance to sort the underlying sequences alphabetically, by the Name fields of the constituent Seqs.
* The Fasta will be modified in place (this is a pointer reciever method).

#### func Query(accession)
* This function takes a slice of strings as an argument, where each of the strings is an NCBI accession number. 
* It will query NCBI for these accession numbers, and return a Fasta type instance containing the a Seq struct corresponding to each of the accession numbers. 

#### func QueryToFile(accession []string, output string)
* This function can be used in lieu of the Query function in instances where the data are not required in memory, they can then be written directly to a file (this is more efficient as the data does not need to be processed into the Fasta structure and the string can be written straight to the file).
* The function takes two argumens. The first argument is a slice of strings where each of the strings is an NCBI accession number. The second argument is a string containing the desired output file name to which the sequences obtained in the NCBI query will be written.

#### func (fa Fasta) Summary()
* This method should be used with nucleotide Fasta structures only. 
* Calling this method will produce a slice of structs with three fields, corresponding to the name, length and percent GC content of the sequences in the Fasta


#### func (fa Fasta) WriteSummary(filename)
* This method has the same functionality as the Summary method, but instead of providing the output slice with the summary data in memory, it writes the summary directly to the file specified as a string in the method call.

