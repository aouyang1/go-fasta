package main

// compile this, for use of the fasta package from a high level

import(
	"fmt"
	"./fasta"
	)

func main(){

	// parse the command line arguments
	flag.Parse() 

	// include arguments for:
		// -f reformat (to pretty print)
		// -s summary:
			// if passed then produce a summary file

		// -n ncbi (batch or single)
			// take either one string, multiple space delimited string or a text file
			// parse the above into a slice of accession numbers, query NCBI for the accession numbers
		
		// -f if passed, change the output file names 
			// for instance of both a .fasta and a summary, 
			// take this name and split on a . , take the first bit and append .fasta and .txt to it and use accordingly
		
		// -m merge (multiple, output file name)
			// take a list of files space delimited (or a .txt with filenames within)
			// merge them and write to -f or default name

		// -d split (single, output file names == fasta names)

			// for the split, have it take a list fasta struct, and split each member of 
			// the fasta into its own fasta struct, then take this and write each to a file
			// using the fasta.Write() function with the name of the sequence + ".fasta"
			// passed in as the second name.

}