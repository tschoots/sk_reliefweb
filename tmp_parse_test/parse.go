package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	//"log"
	//"net/http"
	"os"
	//"strconv"
)

func main() {

	b := []byte(`{"href":"http:\/\/api.reliefweb.int\/v1\/reports?sort[]=id:asc&appname=sokrates.ai&profile=full&preset=analysis&offset=0&limit=2","time":4,"links":{"self":{"href":"http:\/\/api.reliefweb.int\/v1\/reports?offset=0&limit=2&preset=analysis"},"next":{"href":"http:\/\/api.reliefweb.int\/v1\/reports?offset=2&limit=2&preset=analysis"},"prev":{"href":"http:\/\/api.reliefweb.int\/v1\/reports?offset=0&limit=2&preset=analysis"}},"totalCount":650304,"count":2,"data":[{"id":"8","score":1,"href":"http:\/\/api.reliefweb.int\/v1\/reports\/8","fields":{"id":8,"title":"Afghanistan: District Population Per Health Facility","status":"published","body":"<font size=1 face=\"Arial\">Source: National Health Resources<\/font>\n<br><font size=1 face=\"Arial\">Assessment, September 2002<\/font>\n<br><font size=1 face=\"Arial\">MoPH and USAID\/AHSEP<\/font>\n","primary_country":{"href":"http:\/\/api.reliefweb.int\/v1\/countries\/13","id":13,"name":"Afghanistan","iso3":"afg","location":{"lat":33.84,"lon":66.03}},"country":[{"href":"http:\/\/api.reliefweb.int\/v1\/countries\/13","id":13,"name":"Afghanistan","iso3":"afg","location":{"lat":33.84,"lon":66.03},"primary":true}],"source":[{"href":"http:\/\/api.reliefweb.int\/v1\/sources\/360","id":360,"name":"Afghanistan Information Management Service","shortname":"AIMS","longname":"Afghanistan Information Management Service","homepage":"http:\/\/www.aims.org.af\/","type":{"id":272,"name":"International Organization"}}],"language":[{"id":267,"name":"English","code":"en"}],"theme":[{"id":4595,"name":"Health"}],"format":[{"id":12,"name":"Map"}],"url":"http:\/\/reliefweb.int\/node\/8","url_alias":"http:\/\/reliefweb.int\/map\/afghanistan\/afghanistan-district-population-health-facility","body-html":"<p><font size=1 face=\"Arial\">Source: National Health Resources<\/font>\n<br><font size=1 face=\"Arial\">Assessment, September 2002<\/font>\n<br><font size=1 face=\"Arial\">MoPH and USAID\/AHSEP<\/font><\/p>\n","date":{"original":"2002-11-15T05:00:00+00:00","changed":"2011-03-21T12:49:55+00:00","created":"2002-11-15T05:00:00+00:00"}}},{"id":"9","score":1,"href":"http:\/\/api.reliefweb.int\/v1\/reports\/9","fields":{"id":9,"title":"Kenya - Rachuonyo district: NGO\/CBO activity in flood-affected regions","status":"published","body":"<font size=1 face=\"Arial\">Prepared by UNDP Disaster MGT Unit &amp; OCHA\nRSO-CEA - May 2003<\/font>\n","file":[{"id":"724","description":"","url":"http:\/\/reliefweb.int\/sites\/reliefweb.int\/files\/resources\/00081348C06B1D1DC1256F2D0047FF53-ocha_kenRachuonyo190503.pdf","filename":"00081348C06B1D1DC1256F2D0047FF53-ocha_kenRachuonyo190503.pdf","mimetype":"application\/pdf","filesize":"234680","preview":{"url":"http:\/\/reliefweb.int\/sites\/reliefweb.int\/files\/resources-pdf-previews\/724-00081348C06B1D1DC1256F2D0047FF53-ocha_kenRachuonyo190503.png","url-large":"http:\/\/reliefweb.int\/sites\/reliefweb.int\/files\/styles\/attachment-large\/public\/resources-pdf-previews\/724-00081348C06B1D1DC1256F2D0047FF53-ocha_kenRachuonyo190503.png","url-small":"http:\/\/reliefweb.int\/sites\/reliefweb.int\/files\/styles\/attachment-small\/public\/resources-pdf-previews\/724-00081348C06B1D1DC1256F2D0047FF53-ocha_kenRachuonyo190503.png","url-thumb":"http:\/\/reliefweb.int\/sites\/reliefweb.int\/files\/styles\/m\/public\/resources-pdf-previews\/724-00081348C06B1D1DC1256F2D0047FF53-ocha_kenRachuonyo190503.png"}}],"primary_country":{"href":"http:\/\/api.reliefweb.int\/v1\/countries\/131","id":131,"name":"Kenya","iso3":"ken","location":{"lat":0.53,"lon":37.86}},"country":[{"href":"http:\/\/api.reliefweb.int\/v1\/countries\/131","id":131,"name":"Kenya","iso3":"ken","location":{"lat":0.53,"lon":37.86},"primary":true}],"source":[{"href":"http:\/\/api.reliefweb.int\/v1\/sources\/1503","id":1503,"name":"UN Office for the Coordination of Humanitarian Affairs","shortname":"OCHA","longname":"United Nations Office for the Coordination of Humanitarian Affairs","homepage":"http:\/\/www.unocha.org\/","type":{"id":272,"name":"International Organization"}}],"theme":[{"id":4587,"name":"Agriculture"},{"id":4593,"name":"Food and Nutrition"},{"id":4595,"name":"Health"},{"id":4604,"name":"Water Sanitation Hygiene"}],"format":[{"id":12,"name":"Map"}],"ocha_product":[{"id":12354,"name":"Thematic Map"}],"disaster":[{"id":6280,"name":"Horn of Africa: Floods - May 2003","glide":"FL-2003-0204-KEN","type":[{"id":4611,"name":"Flood","code":"FL","primary":true}]}],"disaster_type":[{"id":4611,"name":"Flood","code":"FL"}],"url":"http:\/\/reliefweb.int\/node\/9","url_alias":"http:\/\/reliefweb.int\/map\/kenya\/kenya-rachuonyo-district-ngocbo-activity-flood-affected-regions","body-html":"<p><font size=1 face=\"Arial\">Prepared by UNDP Disaster MGT Unit &amp; OCHA\nRSO-CEA - May 2003<\/font><\/p>\n","date":{"original":"2003-05-19T04:00:00+00:00","changed":"2011-04-12T10:11:30+00:00","created":"2003-05-19T04:00:00+00:00"}}}]}`)
	//b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("ERROR : ", err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})
	fmt.Println(m)
	for k, v := range m {

		switch vv := v.(type) {
		case string:
			//fmt.Printf(k, " is string ", vv)
			fmt.Printf("%s is a string\n\n", k)
		case int:
			//fmt.Println(k, " is int ", vv)
			fmt.Printf("%s is a int\n\n", k)
		case float64:
			fmt.Printf("%s is float\n\n", k)
		case []interface{}:
			fmt.Println(k, " is an array: ")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, " is an map")
		default:
			//fmt.Println(k, " is of a type i don't know how to handle")
			fmt.Println(k)
			fmt.Printf("default formate : %v, Go-syntax representation : %#v, Go-syntax representation of the type of the value : %T\n\n", k, k, k)

		}
	}

}
