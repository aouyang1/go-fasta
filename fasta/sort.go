// take a fasta struct and sort the entries into alphabetical order by their name
// improvement -> sort the sequence in place

package fasta

import (
	"sort"
)

// This method can be called on a Fasta type instance to sort the underlying sequences alphabetically, by the Name fields of the constituent Seqs.
// The Fasta will be modified in place (this is a pointer reciever method).
func (fa *Fasta) Sort() {
	// make a dict where the keys are the sequence names
	// and the values are a pointer to the seq structs
	fasta_dict := make(map[string]int)
	name_list := []string{}

	// value is a index in original entries struct,
	// avoids moving the seq into intermediate structure
	for i, s := range *fa {
		fasta_dict[s.Name] = i
		name_list = append(name_list, s.Name)
	}

	// quicksort of the keys (seq.name)
	sort.Strings(name_list)

	// make a new fasta
	out_fasta := Fasta{}
	// append the original seq to the output in the correct order
	for _, k := range name_list {
		original_pos := fasta_dict[k]
		// below looks a little weird,
		// to slice from the original we need to call the value pointer
		// wrapped in brackets (*fa)
		out_fasta.AddItem((*fa)[original_pos])
	}
	*fa = out_fasta
}
