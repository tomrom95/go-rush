package models

import (
	"reflect"
)

const AltNameTag string = "go_rush_alt_name"

// CREDIT @xjtian & @emakeev
// This is not my own work

// FillIn assigns matching field values of src to dest
// A fields from src & dest are considered "matching" if they have
// identical names & Types
// FillIn will recursively inspect structures to find matching fields
// and return number of successfully set fields
func FillIn(src interface{}, dest interface{}) int {
	if src == nil || dest == nil {
		return 0 // nothing to do
	}
	vs, vd := reflect.ValueOf(src), reflect.ValueOf(dest)
	return fillIn(&vs, &vd)
}

func fillIn(vs *reflect.Value, vd *reflect.Value) int {

	var savedVd *reflect.Value

	ks, kd := vs.Kind(), vd.Kind()
	// First, "dereference" pointers and corresponding parameters
	for ks == reflect.Ptr && kd == reflect.Ptr {
		if vs.IsNil() { // nothing coming from source, return
			return 0
		}
		var vdo reflect.Value
		if vd.IsNil() { // check if we need initialize destination
			if !vd.CanSet() { // dest is not settable, return
				return 0
			}
			savedVd = vd
			vdo = reflect.New(vd.Type().Elem()).Elem()
			vd = &vdo

		} else {
			vdo = vd.Elem()
		}

		vso := vs.Elem()
		vs, vd = &vso, &vdo
		ks, kd = vs.Kind(), vd.Kind()
		if ks != kd && !vs.Type().ConvertibleTo(vd.Type()) {
			return 0
		}
	}

	if !vd.CanSet() {
		return 0 // dest is not settable after differencing - nothing to do
	}

	ts, td := vs.Type(), vd.Type()
	if ts.AssignableTo(td) { // value types match, we can assign directly
		vd.Set(*vs)
	} else if ts.ConvertibleTo(td) { // value types are compatible - assign
		vd.Set(vs.Convert(td))
	} else { // if it's a struct, iterate over all public fields
		count := 0 // keep count of total set fields recursively
		if ks == reflect.Struct && kd == reflect.Struct {
			count = convertStruct(vs, vd)
		}
		if ks == reflect.Map && kd == reflect.Map {
			count = convertMap(vs, vd)
		}
		// if no matching fields (count == 0), leave nil Ptr at nil
		if savedVd != nil && count > 0 {
			savedVd.Set(vd.Addr())
		}
		return count
	}

	// value was set directly, update nil Ptr if needed
	if savedVd != nil {
		savedVd.Set(vd.Addr())
	}
	return 1
}

func convertStruct(vs *reflect.Value, vd *reflect.Value) int {
	count := 0
	altSrcNames := map[string]int{}
	ts, td := vs.Type(), vd.Type()
	for i := 0; i < ts.NumField(); i++ {
		fts := ts.Field(i)
		if altNm := fts.Tag.Get(AltNameTag); len(altNm) > 0 {
			altSrcNames[altNm] = i
		}
	}
	for i := 0; i < td.NumField(); i++ {
		ftd := td.Field(i)
		fs := vs.FieldByName(ftd.Name)
		if !fs.IsValid() {
			if altNm := ftd.Tag.Get(AltNameTag); len(altNm) > 0 {
				fs = vs.FieldByName(altNm)
			}
			if !fs.IsValid() {
				sidx, ok := altSrcNames[ftd.Name]
				if !ok {
					continue
				}
				fs = vs.Field(sidx)
			}
		}
		vdf := vd.Field(i)
		if fs.Kind() == vdf.Kind() || fs.Type().ConvertibleTo(ftd.Type) {
			count += fillIn(&fs, &vdf)
		}
	}
	return count
}

func convertMap(vs *reflect.Value, vd *reflect.Value) int {
	count := 0
	if len(vs.MapKeys()) > 0 {
		vd.Set(reflect.MakeMap(vd.Type()))
	}
	for _, mapKey := range vs.MapKeys() {
		mapVal := vs.MapIndex(mapKey)
		if mapVal.Kind() == reflect.Ptr {
			if mapVal.IsNil() {
				continue
			}
			mapVal = mapVal.Elem()
		}
		copyVal := reflect.New(vd.Type().Elem()).Elem()
		if copyVal.Kind() == reflect.Ptr {
			copyValObj := reflect.New(copyVal.Type().Elem()).Elem()
			fillIn(&mapVal, &copyValObj)
			copyVal.Set(copyValObj.Addr())
		} else {
			fillIn(&mapVal, &copyVal)
		}
		vd.SetMapIndex(mapKey, copyVal)
	}
	return count
}
