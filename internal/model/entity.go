package model

import "reflect"

type EntityResult struct {
	Entity *Entity
}

type Entity struct {
	MessageData *MessageData
}

type MessageData struct {
	EntityID *EntityID
	Identity *Identity
}

type EntityID struct {
	DescriptiveLabel string
	UUID             string
}

type Identity struct {
	SpecificVehicle   *SpecificVehicle
	IdentityTimestamp string
}

type SpecificVehicle struct {
	IFF                *IFF
	CallSign           *CallSign
	DataLinkIdentifier *DataLinkIdentifier
}

type IFF struct {
	Mode1  *Mode1
	Mode2  *Mode2
	Mode3A *Mode3A
	Mode5  *Mode5
	ModeS  *ModeS
}

type Mode1 struct {
	Code string
}

type Mode2 struct {
	Code string
}

type Mode3A struct {
	Code string
}

type Mode5 struct {
	Mode5Indicator string
}

type ModeS struct {
	AircraftIdentifier string
}

type CallSign struct {
	Key string
}

type DataLinkIdentifier struct {
	TrackIdentifier *TrackIdentifier
}

type TrackIdentifier struct {
	TrackNumber string
}

func EntityResultToMap(entityJSON *EntityResult) map[string]any {
	m := make(map[string]any)
	m["id"] = entityJSON.Entity.MessageData.EntityID.UUID
	m["descriptiveLabel"] = entityJSON.Entity.MessageData.EntityID.DescriptiveLabel
	if entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode1 != nil {
		m["mode1"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode1.Code
	}
	if entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode2 != nil {
		m["mode2"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode2.Code
	}
	if entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode3A != nil {
		m["mode3"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode3A.Code
	}
	if entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode5 != nil {
		m["mode5"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.Mode5.Mode5Indicator
	}
	if entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.ModeS != nil {
		m["tailNumber"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.IFF.ModeS.AircraftIdentifier
	}
	if entityJSON.Entity.MessageData.Identity.SpecificVehicle.DataLinkIdentifier != nil {
		m["trackNumber"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.DataLinkIdentifier.TrackIdentifier.TrackNumber
	}
	m["callSign"] = entityJSON.Entity.MessageData.Identity.SpecificVehicle.CallSign.Key
	m["identityTimestamp"] = entityJSON.Entity.MessageData.Identity.IdentityTimestamp

	return m
}

func isEmpty(object interface{}) bool {
	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}
