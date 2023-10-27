package xml

import (
	sj "github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	inputXML := `
		<?xml version="1.0" encoding="UTF-8"?>
		<osm version="0.6" generator="CGImap 0.0.2">
			<bounds minlat="54.0889580" minlon="12.2487570" maxlat="54.0913900" maxlon="12.2524800"/>
			<node id="298884269" lat="54.0901746" lon="12.2482632" user="SvenHRO" uid="46882" visible="true" version="1"
				  changeset="676636" timestamp="2008-09-21T21:37:45Z"/>
			<node id="261728686" lat="54.0906309" lon="12.2441924" user="PikoWinter" uid="36744" visible="true" version="1"
				  changeset="323878" timestamp="2008-05-03T13:39:23Z"/>
			<node id="1831881213" version="1" changeset="12370172" lat="54.0900666" lon="12.2539381" user="lafkor" uid="75625"
				  visible="true" timestamp="2012-07-20T09:43:19Z">
				<tag k="name" v="Neu Broderstorf"/>
				<tag k="traffic_sign" v="city_limit"/>
			</node>
			<foo>bar</foo>
			<mixed attr="attribute">
				content
			</mixed>
		</osm>`

	// Build SimpleJSON
	json, err := sj.NewJson([]byte(`
		{
		  "osm": {
			"-version": "0.6",
			"-generator": "CGImap 0.0.2",
			"bounds": {
			  "-minlat": "54.0889580",
			  "-minlon": "12.2487570",
			  "-maxlat": "54.0913900",
			  "-maxlon": "12.2524800"
			},
			"node": [
			  {
				"-id": "298884269",
				"-lat": "54.0901746",
				"-lon": "12.2482632",
				"-user": "SvenHRO",
				"-uid": "46882",
				"-visible": "true",
				"-version": "1",
				"-changeset": "676636",
				"-timestamp": "2008-09-21T21:37:45Z"
			  },
			  {
				"-id": "261728686",
				"-lat": "54.0906309",
				"-lon": "12.2441924",
				"-user": "PikoWinter",
				"-uid": "36744",
				"-visible": "true",
				"-version": "1",
				"-changeset": "323878",
				"-timestamp": "2008-05-03T13:39:23Z"
			  },
			  {
				"-id": "1831881213",
				"-version": "1",
				"-changeset": "12370172",
				"-lat": "54.0900666",
				"-lon": "12.2539381",
				"-user": "lafkor",
				"-uid": "75625",
				"-visible": "true",
				"-timestamp": "2012-07-20T09:43:19Z",
				"tag": [
				  {
					"-k": "name",
					"-v": "Neu Broderstorf"
				  },
				  {
					"-k": "traffic_sign",
					"-v": "city_limit"
				  }
				]
			  }
			],
			"foo": "bar",
				"mixed": {
					"-attr": "attribute",
					"#content": "content"
				}
		  }
		}`))
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	expected, err := json.MarshalJSON()
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	// Then encode it in JSON
	result, err := ToJSON(strings.NewReader(inputXML))
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	expectedString := string(expected)
	resultString := result.String()
	assert.JSONEq(t, expectedString, resultString, "Expected %s, got %s", expectedString, resultString)
}
