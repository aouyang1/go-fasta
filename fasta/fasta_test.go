package fasta

import (
	"fmt"
	"reflect"
	"testing"
)



func TestIO(t *testing.T){

}


func TestQuery(t *testing.T){

	AF298042 := Fasta{entries: []Seq{ Seq{	name: "AF298042.1 Salvelinus alpinus alpinus haplotype ARCTIC_16 mitochondrial D-loop, partial sequence",
											sequence: "CCACTAATATGTACAATAATGAATATTGTATCTCAACAAATTAGTGTCATAATACATCTATGTATAATATTGCATATTATGTATTTACCCATATATACAATACCTGTATGATGAGTAGTACATCATATGTATTATCAACATAAGTGAATTTAAGCCCTCATATATCAGCATAAACCCAAGATTTACATAAGCTAAACACGTGATAATAACCAACTAGGTTGTTTTAACCTAGATAATTGCTACATTAACAAAACTCCAACTAACACGGGCTCCGTCTTTACCCACCAACTTTCAGCATCAGTCCTACTTAATGTAGTAAGAACCGACCAACGATTTATCAGTAGGCATACTCTTATTGATGGTGAGGGGCAAATATCGTATTAGGTAACATCTCGTGAACTATTCCTGGCATTTGGTTCCTAAGTCGAGGGCTATCCTTAAGAAACCAGCCCCTGAAAGCCGAATGTTAAGCATCTGGTTAATGGTGTCAATCTTATTGTTCGTTACCCACAAAGCCGGGCGTTCTCTTATATGCATAGGGTTCTCCTTT",
							}}}

	single_id := []string{"AF298042.1"}
	out1 := Query(single_id)
	
	if reflect.DeepEqual(out1, AF298042) != true {
		t.Errorf("Single Query of Accession Number from NCBI did not match expected output. Received:\n %v", out1)
	}

	AF298042_AY646679 := Fasta{entries : []Seq{ Seq{	name: "AF298042.1 Salvelinus alpinus alpinus haplotype ARCTIC_16 mitochondrial D-loop, partial sequence",
													sequence: "CCACTAATATGTACAATAATGAATATTGTATCTCAACAAATTAGTGTCATAATACATCTATGTATAATATTGCATATTATGTATTTACCCATATATACAATACCTGTATGATGAGTAGTACATCATATGTATTATCAACATAAGTGAATTTAAGCCCTCATATATCAGCATAAACCCAAGATTTACATAAGCTAAACACGTGATAATAACCAACTAGGTTGTTTTAACCTAGATAATTGCTACATTAACAAAACTCCAACTAACACGGGCTCCGTCTTTACCCACCAACTTTCAGCATCAGTCCTACTTAATGTAGTAAGAACCGACCAACGATTTATCAGTAGGCATACTCTTATTGATGGTGAGGGGCAAATATCGTATTAGGTAACATCTCGTGAACTATTCCTGGCATTTGGTTCCTAAGTCGAGGGCTATCCTTAAGAAACCAGCCCCTGAAAGCCGAATGTTAAGCATCTGGTTAATGGTGTCAATCTTATTGTTCGTTACCCACAAAGCCGGGCGTTCTCTTATATGCATAGGGTTCTCCTTT"},
												Seq{	name: "AY646679.1 Spironucleus barkhanus from wild Arctic charr small subunit ribosomal RNA gene, partial sequence",
													sequence: "AAGATTAAAGCCCTGCATGCCTATGTGTAGACAGTTATATTCATTATTGTGGAGCAAAAACGGCGAACAGCTCATTTATCAGTGGTAAGTGCATACAATGTATTTCGTTGGATAGTAACGGAAAATCTGTTAGTAATACATGAACTGTTTTTAGCATTATGTTAAAAATAATAGTAAGTGCGATTGTATATCTGCCACTGCAGCATCATCTTACGTTGGTGGGATATTTGCCTACCAAGGATTCGACGCTTACGGGGAATTAGGGTTTGACTCCGGAGAATGAGCATGAGAAACAGCTCATACATCTAAGGAAGGCAGCAGGCGCGGAAATTGCCCAATGTATCTTTTATACGAGGCAGTGACAAGAAATGGTAGGCACTTTTGTGCACTATCGAGGGTTAGTGGTATCTTTGCTAACCGTGACTCGTGGGCAAGCTCGGTGCCAGCAGCCGCGGTAATTCCGACACAGGGAGTTTTCCATTTGGTTGCTGCAGTTAAAAAGTTCGTAGTTTACTGACTCTTTCACTATAAGCAAAGCCGAATGCTCCAAGTTTTTTAGCAGTATTTATAGTATGAAATTATAGCGCGGCATTGAACGTAGTTTGGGGTACTCGATAGGGACAGGTGAAATAGGATGATCTATCGAAGACCCACGGTAGCGGAGGCTCCCAACGAAGTCCAAGTGTCACGATCAAGAACTAAAGTCAGGGGATAGACGACGATTAGACACCGTTTTATTCCTGACCCTAAACGATGTCGCCTAGCTGATGGGATTTTTTTCATTTGCCAAGAGAAATCGTAAGGTTTCAGACTCTGGGGGAAGTATGATCGCAAGGTTGAAACTTGAAGGGATTGACGGAGAGGTACCACCAGACGTGGAGTCTGCGGCTCAATTTGACTCAACACGCAAACATTACTAGGCCCAGAAGCTTTGAGGATTGACAGATGAGTGATCTTTCATGATTAAGTTGTTGGTGGTGCATGGCCGTTCTTAGTCCGTGATTTAAATTGTCTGCTTTATTGCGATAACGAACGAGACCTCTATCAGATTTATTATCTGAGACTGCTAGTGATGAACTAGAGGAAGGCAGAGGCAAAAACAGGTCTGTGATGCCCTTAGAAGCCCTAGGCCGCACGCGTACTACAATGGCAGGTTCATCGTGTTGCTTCCCTGAAAATGGTGGCAGTTCATTAAAACTTGTCGTGGTTAGGACTGAAGGTTGAAATTATCCTTCACGAATGAGGAATGTCTAGTAAGTGTAGGTTATGAATCTACGCTGATTACGTCCCTACCCCTTGTACACACCGCCCGTCGCTCCTACTGATTGGGAAGATCTGGTGAGTTATTCGGACCCATAGGTAAGCAATTATCTGTGGTAACAATTGCGAGCCAACTCTTCTAGAGGAAGG"},
												}}

	list_of_ids := []string{"AY646679.1","AF298042.1" }
	out2 := Query(list_of_ids)

	if reflect.DeepEqual(out2, AF298042_AY646679) != true {
		t.Errorf("Multiple Query of Accession Number from NCBI did not match expected output. Received:\n %v", out2)
	}
}

func TestSort(t *testing.T){

	unsorted := Fasta{entries : []Seq{Seq{name: "sdsdsd", sequence: "CAT"},
									Seq{name: "chr1", sequence: "GC"},
									Seq{name: "1chr", sequence: "ATGC"},
									Seq{name: "chr2", sequence: "AATT"}}}

	sorted := Fasta{entries : []Seq {Seq{name: "1chr", sequence: "ATGC"},
									Seq{name: "chr1", sequence: "GC"},
									Seq{name: "chr2", sequence: "AATT"},
									Seq{name: "sdsdsd", sequence: "CAT"}}}
	
	unsorted.Sort()			
	if reflect.DeepEqual(unsorted, sorted) != true {
		t.Errorf("Sorting of Fasta structure incorrect: %v\n want: %v.", unsorted, sorted)
	}
}

func TestSummary(t *testing.T){

}






// a test syntax example I have from elsewhere
func TestQueue(t *testing.T) {
	q := Queue{}
	compare_q := Queue{ord: []int{7, 8, 9}}
	compare_q2 := Queue{ord: []int{8, 9}}

	q.Add(7)
	q.Add(8)
	q.Add(9)

	if reflect.DeepEqual(q.ord, compare_q.ord) != true {
		t.Errorf("Adding to Queue incorrect: %v, want: %v.", q, compare_q)
	}
}