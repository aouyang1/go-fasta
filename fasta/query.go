package fasta

import(
    "fmt"
    "net/http"
    "strings"
    "log"
    "io/ioutil"
    "os"
    "io"
)


type UID struct {
    all []string
}

func buildURL(accession UID) string {
    url_front := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id="
    
    // build the middle of the url from the input slice
    url_middle := ""
    for _, i := range accession.all {
        url_middle = fmt.Sprintf("%v,%v", url_middle, i)
    }
    
    url_end := "&rettype=fasta&retmode=text"

    url := []string{url_front, url_middle, url_end}

    return strings.Join(url, "")

}


// take the query unique IDs and get string response
func Query(accession UID) seq {
    //construct the url
    query_url := buildURL(accession)
    
    // make the http request
    resp, err := http.Get(query_url)
    if err != nil {
            log.Fatal(err)
    }        
    defer resp.Body.Close()

    // the the response data to variable
    body, err := ioutil.ReadAll(resp.Body)

    // parse data to string, pass to parser
    return ParseFasta(string(body))

}


func QueryToFile(accession UID, output string) error {
    // construct the url
    query_url := buildURL(accession)

    //make the file
    out, err := os.Create(output)    
    if err != nil {
        return err
    }
    defer out.Close()

    // make the http request
    resp, err := http.Get(query_url)
    if err != nil {
            log.Fatal(err)
    } 
    defer resp.Body.Close()

    // Write data to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}




// take an accession number, query NCBI and then return and use the data
// treat the return from NCBI like a file, send it to

// base url:
fetch := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi"

//
db := "?db="
// then fill in and add the db

id_base := "&id="
// then a comma delimited list of unique ids

// this says we want a fasta in text mode
fa_ext := "&rettype=fasta&retmode=text"


// as a test try to get these small charr sequences
// AY677181.1, in the nucleotide db, 
test1 := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id=AY677181.1&rettype=fasta&retmode=text"
// this is chr 1 from the genome, also a nucleotide db query
test2 := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id=NC_036838.1&rettype=fasta&retmode=text"

// AY646679.1
// AF298042.1
// NC_036838.1 chr one of genome

// make nucleotide the default, have option to switch to
// protein
// AQV08101.1
// both of the below work... can all protein sequences be queried through the nucleotide url?
test3 := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=protein&id=AQV08101.1&rettype=fasta&retmode=text"
test4 := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id=AQV08101.1&rettype=fasta&retmode=text"

//  AFF19513.2 <- a protein, but below the query to the nucleotide works fine
test4 := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id=AFF19513.2&rettype=fasta&retmode=text"


/*
// this is done using eutils - get a sense of the parts of the records needed
// https://www.ncbi.nlm.nih.gov/books/NBK25500/#chapter1.Downloading_Full_Records
// here are the components we need

Required Parameters

db - say if we are grabbing this from bioproject, biosample etc.
id - this is the id param (can pass in a comma delimited list, look at fmt options)
fasta - fa_extension to say we want a fasta returned
retmode - return mode - do you want a text or xml file




?db=

&rettype=fasta

ID types - how can we get an id number
	GenBank-Accn	RefSeq-Accn
	CM009399.1		NC_036838.1
ID must me a UID:
unique record identifier (UID), 
that unambiguously differentiates the record from all other records in the database.



efetch.fcgi?db=database&id=uid1,uid2,uid3&rettype=report_type&retmode=data_mode
elink.fcgi?dbfrom=initial_databasedb=target_database&id=uid1,uid2,uid3

nucleotide example:
https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=nucleotide&id=28864546,28800981


*/



/*

Nucleotide	GI number	nuccore



database_options:

Entrez_Database	UID_common_name	E-utility_Database_Name
BioProject	BioProject ID	bioproject
BioSample	BioSample ID	biosample
Biosystems	BSID	biosystems
Books	Book ID	books
Conserved Domains	PSSM-ID	cdd
dbGaP	dbGaP ID	gap
dbVar	dbVar ID	dbvar
Epigenomics	Epigenomics ID	epigenomics
EST	GI number	nucest
Gene	Gene ID	gene
Genome	Genome ID	genome
GEO Datasets	GDS ID	gds
GEO Profiles	GEO ID	geoprofiles
GSS	GI number	nucgss
HomoloGene	HomoloGene ID	homologene
MeSH	MeSH ID	mesh
NCBI C++ Toolkit	Toolkit ID	toolkit
NCBI Web Site	Web Site ID	ncbisearch
NLM Catalog	NLM Catalog ID	nlmcatalog
Nucleotide	GI number	nuccore
OMIA	OMIA ID	omia
PopSet	PopSet ID	popset
Probe	Probe ID	probe
Protein	GI number	protein
Protein Clusters	Protein Cluster ID	proteinclusters
PubChem BioAssay	AID	pcassay
PubChem Compound	CID	pccompound
PubChem Substance	SID	pcsubstance
PubMed	PMID	pubmed
PubMed Central	PMCID	pmc
SNP	rs number	snp
SRA	SRA ID	sra
Structure	MMDB-ID	structure
Taxonomy	TaxID	taxonomy
UniGene	UniGene Cluster ID	unigene
UniSTS	STS ID	unists





Database from which to retrieve records. 
The value must be a valid Entrez database name (default = pubmed). 
Currently EFetch does not support all Entrez databases. 
Please see Table 1 in Chapter 2 for a list of available databases.

Required Parameter – Used only when input is from a UID list
id
UID list. Either a single UID or a comma-delimited list of UIDs may be provided. 
All of the UIDs must be from the database specified by db. 
There is no set maximum for the number of UIDs that can be passed to EFetch, 
but if more than about 200 UIDs are to be provided, the request should be made 
using the HTTP POST method.

efetch.fcgi?db=protein&id=15718680,157427902,119703751
Required Parameters – Used only when input is from the Entrez History server
query_key
Query key. This integer specifies which of the UID lists attached to the given 
Web Environment will be used as input to EFetch. Query keys are obtained from 
the output of previous ESearch, EPost or ELInk calls. The query_key parameter 
must be used in conjunction with WebEnv.

WebEnv
Web Environment. This parameter specifies the Web Environment that contains 
the UID list to be provided as input to EFetch. Usually this WebEnv value is 
obtained from the output of a previous ESearch, EPost or ELink call. 
The WebEnv parameter must be used in conjunction with query_key.

efetch.fcgi?db=protein&query_key=<key>&WebEnv=<webenv string>
Optional Parameters – Retrieval
retmode
Retrieval mode. This parameter specifies the data format of the records returned, 
such as plain text, HMTL or XML. See Table 1 for a full list of allowed values for 
each database.

rettype
Retrieval type. This parameter specifies the record view returned, 
such as Abstract or MEDLINE from PubMed, or GenPept or FASTA from protein. 
Please see Table 1 for a full list of allowed values for each database.

FASTA	fasta	text


PubMed:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=pubmed&id=11850928,11482001

PubMed, version 2.0 XML:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=pubmed&id=11850928,11482001&version=2.0

Protein:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=protein&id=28800982,28628843

Nucleotide:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=nucleotide&id=28864546,28800981

Structure:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=structure&id=19923,12120

Taxonomy:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=taxonomy&id=9913,30521

UniSTS:

https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=unists&id=254085,254086

*/