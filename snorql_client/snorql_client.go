package main 

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

)

func main() {
	
	client := &http.Client{}
	
	
	/*
	
	   SELECT * WHERE {
                        {
                           <http://dbpedia.org/resource/Jupiter> <http://dbpedia.org/ontology/albedo> ?VAR_1
                        }
                       }
	
	*/
	
	//snorql_query := `http://dbpedia.org/snorql/?query=SELECT+*+WHERE+%7B%0D%0A+++%7B%0D%0A%3Chttp%3A%2F%2Fdbpedia.org%2Fresource%2FJupiter%3E+%3Chttp%3A%2F%2Fdbpedia.org%2Fontology%2Falbedo%3E+%3FVAR_1%0D%0A%7D%0D%0A%7D`
	snorql_query := `http://dbpedia.org/sparql/?query=SELECT+*+WHERE+%7B%0D%0A+++%7B%0D%0A%3Chttp%3A%2F%2Fdbpedia.org%2Fresource%2FJupiter%3E+%3Chttp%3A%2F%2Fdbpedia.org%2Fontology%2Falbedo%3E+%3FVAR_1%0D%0A%7D%0D%0A%7D`
	
	prr, _ := url.QueryUnescape(snorql_query)
	
//	query := `SELECT * WHERE {
//                        {
//                           <http://dbpedia.org/resource/Jupiter> <http://dbpedia.org/ontology/albedo> ?VAR_1
//                        }
//                       }`
	query := `SELECT * WHERE { {
<http://dbpedia.org/resource/Napoleon> <http://dbpedia.org/ontology/parent> ?VAR_1
} UNION 
{
<http://dbpedia.org/resource/Napoleon> <http://dbpedia.org/property/mother> ?VAR_1
} UNION 
{
<http://dbpedia.org/resource/Napoleon> <http://dbpedia.org/property/father> ?VAR_1
} }`

	
	query_esc := url.QueryEscape(query)
	sparql_request := fmt.Sprintf("http://dbpedia.org/sparql/?query=%s", query_esc)
	
	resp, err := client.Get(sparql_request)
	if err != nil {
		fmt.Printf("ERROR get : %s", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("response status : %s\n", resp.Status)
	fmt.Printf("response code : %d\n", resp.StatusCode)

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		fmt.Printf("ERROR reading body : %s", err)
		return
	}
	
	fmt.Println(buf.String())
	
	fmt.Printf("\n\n\n%s\n", prr)
	

}

