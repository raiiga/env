package env

import (
	. "github.com/raiiga/env/internal"
	"reflect"
	"strings"
)

func Marshal(entity any) error {
	typeOf, valueOf := reflect.TypeOf(entity).Elem(), reflect.ValueOf(entity).Elem()

	for i, l := 0, typeOf.NumField(); i < l; i++ {
		if lookup, ok := typeOf.Field(i).Tag.Lookup(EntryTag); ok {
			if err := marshal(lookup, valueOf.Field(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

func marshal(lookup string, fieldValue reflect.Value) error {
	params, split := map[string]string{}, strings.Split(lookup, EntrySeparator)

	for _, s := range split {
		if i := strings.Split(strings.TrimSpace(s), EntryColon); len(i) == 2 {
			params[strings.TrimSpace(i[0])] = strings.TrimSpace(i[1])
		}
	}

	if len(params) == 0 {
		return nil
	}

	return NewEntry(params).Assign(fieldValue)
}
