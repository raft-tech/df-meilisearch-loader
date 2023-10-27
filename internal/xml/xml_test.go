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
	<Entity xmlns="https://www.vdl.afrl.af.mil/programs/oam" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xsi:schemaLocation="https://www.vdl.afrl.af.mil/programs/oam">
		<SecurityInformation>
			<Classification>U</Classification>
			<OwnerProducer>
				<GovernmentIdentifier>USA</GovernmentIdentifier>
			</OwnerProducer>
		</SecurityInformation>
		<MessageHeader>
			<SystemID>
				<UUID>4405E2F580FA4BEB893AA9230C62ABE7</UUID>
				<DescriptiveLabel>CTC</DescriptiveLabel>
			</SystemID>
			<Timestamp>2023-03-07T22:34:14.000007382Z</Timestamp>
			<SchemaVersion>002.1</SchemaVersion>
			<Mode>LIVE</Mode>
			<ServiceID>
				<UUID>AC840245C1DB4B379B525D8229118D43</UUID>
				<DescriptiveLabel>CTC</DescriptiveLabel>
			</ServiceID>
		</MessageHeader>
		<ObjectState>NEW</ObjectState>
		<MessageData>
			<EntityID>
				<UUID>263F5083FEC1459E9C7E946EE2C8AE6B</UUID>
				<DescriptiveLabel>28767</DescriptiveLabel>
			</EntityID>
			<CreationTimestamp>
				<DateTime>2023-03-07T22:34:14.000007382Z</DateTime>
			</CreationTimestamp>
			<Source>
				<SystemID>
					<UUID>4405E2F580FA4BEB893AA9230C62ABE7</UUID>
					<DescriptiveLabel>CTC</DescriptiveLabel>
				</SystemID>
				<SourceSpecificData>
					<TrackQuality>10</TrackQuality>
				</SourceSpecificData>
				<SourceType>FUSION_SYSTEM</SourceType>
			</Source>
			<EntityStatus>POTENTIAL</EntityStatus>
			<Identity>
				<Standard>
					<StandardIdentity>ASSUMED_FRIEND</StandardIdentity>
					<Confidence>100.0</Confidence>
					<ExerciseIdentityData>
						<ExerciseIdentity>ASSUMED_FRIEND</ExerciseIdentity>
					</ExerciseIdentityData>
				</Standard>
				<Environment>
					<Environment>AIR</Environment>
					<Confidence>100.0</Confidence>
				</Environment>
				<Platform>
					<PlatformType>0</PlatformType>
					<PlatformTypeCategory>AIR</PlatformTypeCategory>
					<Confidence>100.0</Confidence>
				</Platform>
				<Specific>
					<SpecificType>0</SpecificType>
					<SpecificTypeCategory>AIR</SpecificTypeCategory>
					<Confidence>100.0</Confidence>
				</Specific>
				<SpecificVehicle>
					<IFF>
						<Mode3A>
							<Code>1002</Code>
							<Enabled>true</Enabled>
						</Mode3A>
						<Mode4>
							<Mode4Indicator>NOT_INTERROGATED</Mode4Indicator>
						</Mode4>
						<Mode5>
							<NationalOrigin>0</NationalOrigin>
							<PIN>0</PIN>
							<Mode5Indicator>NOT_INTERROGATED</Mode5Indicator>
						</Mode5>
						<ModeC>
							<Code>360</Code>
							<Enabled>true</Enabled>
						</ModeC>
					</IFF>
					<CallSign>
						<Key/>
						<SystemName>TBD</SystemName>
					</CallSign>
					<Confidence>100.0</Confidence>
				</SpecificVehicle>
				<SelfReportedIdentity>false</SelfReportedIdentity>
				<DifferenceIndicator>false</DifferenceIndicator>
				<IdentityTimestamp>2023-03-07T22:34:14.000007382Z</IdentityTimestamp>
			</Identity>
			<Kinematics>
				<Position>
					<FixedPositionType>
						<FixedPoint>
							<Latitude>0.7561529419480809</Latitude>
							<Longitude>-1.8002823793348235</Longitude>
							<Altitude>10929.8882</Altitude>
							<Timestamp>2023-03-07T22:34:14.000007382Z</Timestamp>
						</FixedPoint>
					</FixedPositionType>
				</Position>
				<Velocity>
					<NorthSpeed>106.4826</NorthSpeed>
					<EastSpeed>-181.8195</EastSpeed>
				</Velocity>
			</Kinematics>
			<Strength>
				<StrengthValue>
					<Minimum>9</Minimum>
					<Maximum>9</Maximum>
				</StrengthValue>
			</Strength>
			<ActivityBy>
				<Activity>0</Activity>
				<ActivityCategory>AIR</ActivityCategory>
			</ActivityBy>
		</MessageData>
	</Entity>`

	// Build SimpleJSON
	json, err := sj.NewJson([]byte(`
	{
	  "Entity": {
		"MessageData": {
		  "Strength": {
			"StrengthValue": {
			  "Minimum": "9",
			  "Maximum": "9"
			}
		  },
		  "ActivityBy": {
			"Activity": "0",
			"ActivityCategory": "AIR"
		  },
		  "EntityID": {
			"UUID": "263F5083FEC1459E9C7E946EE2C8AE6B",
			"DescriptiveLabel": "28767"
		  },
		  "CreationTimestamp": {
			"DateTime": "2023-03-07T22:34:14.000007382Z"
		  },
		  "Source": {
			"SystemID": {
			  "UUID": "4405E2F580FA4BEB893AA9230C62ABE7",
			  "DescriptiveLabel": "CTC"
			},
			"SourceSpecificData": {
			  "TrackQuality": "10"
			},
			"SourceType": "FUSION_SYSTEM"
		  },
		  "EntityStatus": "POTENTIAL",
		  "Identity": {
			"IdentityTimestamp": "2023-03-07T22:34:14.000007382Z",
			"Standard": {
			  "StandardIdentity": "ASSUMED_FRIEND",
			  "Confidence": "100.0",
			  "ExerciseIdentityData": {
				"ExerciseIdentity": "ASSUMED_FRIEND"
			  }
			},
			"Environment": {
			  "Environment": "AIR",
			  "Confidence": "100.0"
			},
			"Platform": {
			  "PlatformType": "0",
			  "PlatformTypeCategory": "AIR",
			  "Confidence": "100.0"
			},
			"Specific": {
			  "SpecificType": "0",
			  "SpecificTypeCategory": "AIR",
			  "Confidence": "100.0"
			},
			"SpecificVehicle": {
			  "IFF": {
				"Mode5": {
				  "NationalOrigin": "0",
				  "PIN": "0",
				  "Mode5Indicator": "NOT_INTERROGATED"
				},
				"ModeC": {
				  "Code": "360",
				  "Enabled": "true"
				},
				"Mode3A": {
				  "Code": "1002",
				  "Enabled": "true"
				},
				"Mode4": {
				  "Mode4Indicator": "NOT_INTERROGATED"
				}
			  },
			  "CallSign": {
				"Key": "",
				"SystemName": "TBD"
			  },
			  "Confidence": "100.0"
			},
			"SelfReportedIdentity": "false",
			"DifferenceIndicator": "false"
		  },
		  "Kinematics": {
			"Position": {
			  "FixedPositionType": {
				"FixedPoint": {
				  "Latitude": "0.7561529419480809",
				  "Longitude": "-1.8002823793348235",
				  "Altitude": "10929.8882",
				  "Timestamp": "2023-03-07T22:34:14.000007382Z"
				}
			  }
			},
			"Velocity": {
			  "EastSpeed": "-181.8195",
			  "NorthSpeed": "106.4826"
			}
		  }
		},
		"-xmlns": "https://www.vdl.afrl.af.mil/programs/oam",
		"-xsi": "http://www.w3.org/2001/XMLSchema-instance",
		"-schemaLocation": "https://www.vdl.afrl.af.mil/programs/oam",
		"SecurityInformation": {
		  "Classification": "U",
		  "OwnerProducer": {
			"GovernmentIdentifier": "USA"
		  }
		},
		"MessageHeader": {
		  "SystemID": {
			"UUID": "4405E2F580FA4BEB893AA9230C62ABE7",
			"DescriptiveLabel": "CTC"
		  },
		  "Timestamp": "2023-03-07T22:34:14.000007382Z",
		  "SchemaVersion": "002.1",
		  "Mode": "LIVE",
		  "ServiceID": {
			"UUID": "AC840245C1DB4B379B525D8229118D43",
			"DescriptiveLabel": "CTC"
		  }
		},
		"ObjectState": "NEW"
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
