package footballdata

import (
	"fmt"
	"strings"
	"time"
)

// Contains information about a soccerseason update sent to us by the Event API.
type SoccerSeasonUpdate struct {
	Timestamp time.Time
	Resource  string
	Id        uint64
	URI       string
	Updates   UpdateDescriptor
}

type FieldUpdateValues struct {
	OldValue string
	NewValue string
}

func (v FieldUpdateValues) String() string {
	return fmt.Sprintf("%s -> %s", v.OldValue, v.NewValue)
}

type FieldUpdate struct {
	Name   string
	Values *FieldUpdateValues
}

func (v FieldUpdate) String() string {
	r := []string{v.Name}
	if v.Values != nil {
		r = append(r, v.Values.String())
	}
	return strings.Join(r, "|")
}

type UpdateDescriptor string

func (d UpdateDescriptor) Fields() (retval []FieldUpdate) {
	// (FIELD_NAME|OLD_VALUE -> NEW_VALUE)(;($1))?
	retval = []FieldUpdate{}
	fieldStrs := strings.Split(string(d), ";")
	for _, fieldStr := range fieldStrs {
		field := FieldUpdate{}
		fieldStrSplit := strings.SplitN(fieldStr, "|", 2)
		field.Name = fieldStrSplit[0]
		if len(fieldStrSplit) > 1 {
			// OLD_VALUE -> NEW_VALUE
			valueSplit := strings.Split(fieldStrSplit[1], " -> ")
			field.Values = &FieldUpdateValues{
				OldValue: valueSplit[0],
				NewValue: valueSplit[1],
			}
		}
		retval = append(retval, field)
	}
	return
}
