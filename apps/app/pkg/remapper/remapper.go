package remapper

import (
	"reflect"
	"sort"
	"unsafe"
)

const (
	kindPlain  = 0
	kindStruct = iota
	kindSlice  = iota
)

type field struct {
	kind byte
	size byte
	srcO byte
	dstO byte
}

type Mapper struct {
	fields []field
}

func New(src, dst any) Mapper {
	// TODO: Fix panic messages
	st := reflect.TypeOf(src)
	if st.Kind() != reflect.Pointer {
		panic("not a ptr")
	}
	dt := reflect.TypeOf(dst)
	if dt.Kind() != reflect.Pointer {
		panic("not a ptr")
	}
	se := st.Elem()
	if se.Kind() != reflect.Struct {
		panic("not a struct")
	}
	sm := collectFields(se)
	de := dt.Elem()
	if de.Kind() != reflect.Struct {
		panic("not a struct")
	}
	dm := collectFields(de)
	fields := make([]field, 0)
	for k, s := range sm {
		d, ok := dm[k]
		if !ok {
			continue
		}
		if s.Type != d.Type {
			panic("type mismatch")
		}
		if s.Type.Kind() == reflect.Slice { // TODO: Support slices
			panic("slice field is not supported")
		}
		if s.Type.Kind() == reflect.Map { // TODO: Support maps
			panic("map field is not supported")
		}
		f := field{
			kind: 0,
			size: byte(s.Type.Size()),
			srcO: byte(s.Offset),
			dstO: byte(d.Offset),
		}
		fields = append(fields, f)
	}
	return Mapper{combineAdjacentFields(fields)}
}

func combineAdjacentFields(fields []field) []field {
	if len(fields) < 2 {
		return fields
	}

	sorted := make([]field, len(fields))
	copy(sorted, fields)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].srcO < sorted[j].srcO
	})

	var combined []field
	current := sorted[0]

	for i := 1; i < len(sorted); i++ {
		next := sorted[i]

		if areAdjacent(current, next) {
			current.size += next.size
		} else {
			combined = append(combined, current)
			current = next
		}
	}

	combined = append(combined, current)
	return combined
}

func areAdjacent(current, next field) bool {
	return current.kind == next.kind &&
		current.srcO+current.size == next.srcO &&
		current.dstO+current.size == next.dstO
}

func collectFields(t reflect.Type) map[string]reflect.StructField {
	res := map[string]reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		res[f.Name] = f
	}
	return res
}

func (m Mapper) Map(src, dst any) {
	sip := (*[2]uintptr)(unsafe.Pointer(&src))
	dip := (*[2]uintptr)(unsafe.Pointer(&dst))

	sp := unsafe.Pointer(sip[1])
	dp := unsafe.Pointer(dip[1])

	for _, f := range m.fields {
		sfp := unsafe.Pointer(uintptr(sp) + uintptr(f.srcO))
		dfp := unsafe.Pointer(uintptr(dp) + uintptr(f.dstO))
		switch f.kind {
		case kindPlain:
			switch f.size {
			case 1:
				*(*uint8)(dfp) = *(*uint8)(sfp)
			case 2:
				*(*uint16)(dfp) = *(*uint16)(sfp)
			case 4:
				*(*uint32)(dfp) = *(*uint32)(sfp)
			case 8:
				*(*uint64)(dfp) = *(*uint64)(sfp)
			default:
				copy(
					unsafe.Slice((*byte)(dfp), int(f.size)),
					unsafe.Slice((*byte)(sfp), int(f.size)),
				)
			}
		default:
			panic("not supported")
		}
	}
}

