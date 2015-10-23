// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Incident",
			[]types.Field{
				types.Field{"ID", types.MakePrimitiveTypeRef(types.Int64Kind), false},
				types.Field{"Category", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Description", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"DayOfWeek", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Date", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Time", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"PdDistrict", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Resolution", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Address", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Geoposition", types.MakeTypeRef(ref.Parse("sha1-fb09d21d144c518467325465327d46489cff7c47"), 0), false},
				types.Field{"PdID", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-fb09d21d144c518467325465327d46489cff7c47"),
	})
	__mainPackageInFile_types_CachedRef = types.RegisterPackage(&p)
}

// Incident

type Incident struct {
	m   types.Map
	ref *ref.Ref
}

func NewIncident() Incident {
	return Incident{types.NewMap(
		types.NewString("ID"), types.Int64(0),
		types.NewString("Category"), types.NewString(""),
		types.NewString("Description"), types.NewString(""),
		types.NewString("DayOfWeek"), types.NewString(""),
		types.NewString("Date"), types.NewString(""),
		types.NewString("Time"), types.NewString(""),
		types.NewString("PdDistrict"), types.NewString(""),
		types.NewString("Resolution"), types.NewString(""),
		types.NewString("Address"), types.NewString(""),
		types.NewString("Geoposition"), NewGeoposition(),
		types.NewString("PdID"), types.NewString(""),
	), &ref.Ref{}}
}

type IncidentDef struct {
	ID          int64
	Category    string
	Description string
	DayOfWeek   string
	Date        string
	Time        string
	PdDistrict  string
	Resolution  string
	Address     string
	Geoposition GeopositionDef
	PdID        string
}

func (def IncidentDef) New() Incident {
	return Incident{
		types.NewMap(
			types.NewString("ID"), types.Int64(def.ID),
			types.NewString("Category"), types.NewString(def.Category),
			types.NewString("Description"), types.NewString(def.Description),
			types.NewString("DayOfWeek"), types.NewString(def.DayOfWeek),
			types.NewString("Date"), types.NewString(def.Date),
			types.NewString("Time"), types.NewString(def.Time),
			types.NewString("PdDistrict"), types.NewString(def.PdDistrict),
			types.NewString("Resolution"), types.NewString(def.Resolution),
			types.NewString("Address"), types.NewString(def.Address),
			types.NewString("Geoposition"), def.Geoposition.New(),
			types.NewString("PdID"), types.NewString(def.PdID),
		), &ref.Ref{}}
}

func (s Incident) Def() (d IncidentDef) {
	d.ID = int64(s.m.Get(types.NewString("ID")).(types.Int64))
	d.Category = s.m.Get(types.NewString("Category")).(types.String).String()
	d.Description = s.m.Get(types.NewString("Description")).(types.String).String()
	d.DayOfWeek = s.m.Get(types.NewString("DayOfWeek")).(types.String).String()
	d.Date = s.m.Get(types.NewString("Date")).(types.String).String()
	d.Time = s.m.Get(types.NewString("Time")).(types.String).String()
	d.PdDistrict = s.m.Get(types.NewString("PdDistrict")).(types.String).String()
	d.Resolution = s.m.Get(types.NewString("Resolution")).(types.String).String()
	d.Address = s.m.Get(types.NewString("Address")).(types.String).String()
	d.Geoposition = s.m.Get(types.NewString("Geoposition")).(Geoposition).Def()
	d.PdID = s.m.Get(types.NewString("PdID")).(types.String).String()
	return
}

var __typeRefForIncident types.TypeRef

func (m Incident) TypeRef() types.TypeRef {
	return __typeRefForIncident
}

func init() {
	__typeRefForIncident = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)
	types.RegisterFromValFunction(__typeRefForIncident, func(v types.Value) types.Value {
		return IncidentFromVal(v)
	})
}

func IncidentFromVal(val types.Value) Incident {
	// TODO: Do we still need FromVal?
	if val, ok := val.(Incident); ok {
		return val
	}
	// TODO: Validate here
	return Incident{val.(types.Map), &ref.Ref{}}
}

func (s Incident) InternalImplementation() types.Map {
	return s.m
}

func (s Incident) Equals(other types.Value) bool {
	if other, ok := other.(Incident); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s Incident) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Incident) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s Incident) ID() int64 {
	return int64(s.m.Get(types.NewString("ID")).(types.Int64))
}

func (s Incident) SetID(val int64) Incident {
	return Incident{s.m.Set(types.NewString("ID"), types.Int64(val)), &ref.Ref{}}
}

func (s Incident) Category() string {
	return s.m.Get(types.NewString("Category")).(types.String).String()
}

func (s Incident) SetCategory(val string) Incident {
	return Incident{s.m.Set(types.NewString("Category"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Description() string {
	return s.m.Get(types.NewString("Description")).(types.String).String()
}

func (s Incident) SetDescription(val string) Incident {
	return Incident{s.m.Set(types.NewString("Description"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) DayOfWeek() string {
	return s.m.Get(types.NewString("DayOfWeek")).(types.String).String()
}

func (s Incident) SetDayOfWeek(val string) Incident {
	return Incident{s.m.Set(types.NewString("DayOfWeek"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Date() string {
	return s.m.Get(types.NewString("Date")).(types.String).String()
}

func (s Incident) SetDate(val string) Incident {
	return Incident{s.m.Set(types.NewString("Date"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Time() string {
	return s.m.Get(types.NewString("Time")).(types.String).String()
}

func (s Incident) SetTime(val string) Incident {
	return Incident{s.m.Set(types.NewString("Time"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) PdDistrict() string {
	return s.m.Get(types.NewString("PdDistrict")).(types.String).String()
}

func (s Incident) SetPdDistrict(val string) Incident {
	return Incident{s.m.Set(types.NewString("PdDistrict"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Resolution() string {
	return s.m.Get(types.NewString("Resolution")).(types.String).String()
}

func (s Incident) SetResolution(val string) Incident {
	return Incident{s.m.Set(types.NewString("Resolution"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Address() string {
	return s.m.Get(types.NewString("Address")).(types.String).String()
}

func (s Incident) SetAddress(val string) Incident {
	return Incident{s.m.Set(types.NewString("Address"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Geoposition() Geoposition {
	return s.m.Get(types.NewString("Geoposition")).(Geoposition)
}

func (s Incident) SetGeoposition(val Geoposition) Incident {
	return Incident{s.m.Set(types.NewString("Geoposition"), val), &ref.Ref{}}
}

func (s Incident) PdID() string {
	return s.m.Get(types.NewString("PdID")).(types.String).String()
}

func (s Incident) SetPdID(val string) Incident {
	return Incident{s.m.Set(types.NewString("PdID"), types.NewString(val)), &ref.Ref{}}
}

// ListOfIncident

type ListOfIncident struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfIncident() ListOfIncident {
	return ListOfIncident{types.NewList(), &ref.Ref{}}
}

type ListOfIncidentDef []IncidentDef

func (def ListOfIncidentDef) New() ListOfIncident {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfIncident{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfIncident) Def() ListOfIncidentDef {
	d := make([]IncidentDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Incident).Def()
	}
	return d
}

func ListOfIncidentFromVal(val types.Value) ListOfIncident {
	// TODO: Do we still need FromVal?
	if val, ok := val.(ListOfIncident); ok {
		return val
	}
	// TODO: Validate here
	return ListOfIncident{val.(types.List), &ref.Ref{}}
}

func (l ListOfIncident) InternalImplementation() types.List {
	return l.l
}

func (l ListOfIncident) Equals(other types.Value) bool {
	if other, ok := other.(ListOfIncident); ok {
		return l.Ref() == other.Ref()
	}
	return false
}

func (l ListOfIncident) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfIncident) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfIncident.
var __typeRefForListOfIncident types.TypeRef

func (m ListOfIncident) TypeRef() types.TypeRef {
	return __typeRefForListOfIncident
}

func init() {
	__typeRefForListOfIncident = types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForListOfIncident, func(v types.Value) types.Value {
		return ListOfIncidentFromVal(v)
	})
}

func (l ListOfIncident) Len() uint64 {
	return l.l.Len()
}

func (l ListOfIncident) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfIncident) Get(i uint64) Incident {
	return l.l.Get(i).(Incident)
}

func (l ListOfIncident) Slice(idx uint64, end uint64) ListOfIncident {
	return ListOfIncident{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfIncident) Set(i uint64, val Incident) ListOfIncident {
	return ListOfIncident{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfIncident) Append(v ...Incident) ListOfIncident {
	return ListOfIncident{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfIncident) Insert(idx uint64, v ...Incident) ListOfIncident {
	return ListOfIncident{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfIncident) Remove(idx uint64, end uint64) ListOfIncident {
	return ListOfIncident{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfIncident) RemoveAt(idx uint64) ListOfIncident {
	return ListOfIncident{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfIncident) fromElemSlice(p []Incident) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfIncidentIterCallback func(v Incident, i uint64) (stop bool)

func (l ListOfIncident) Iter(cb ListOfIncidentIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Incident), i)
	})
}

type ListOfIncidentIterAllCallback func(v Incident, i uint64)

func (l ListOfIncident) IterAll(cb ListOfIncidentIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Incident), i)
	})
}

type ListOfIncidentFilterCallback func(v Incident, i uint64) (keep bool)

func (l ListOfIncident) Filter(cb ListOfIncidentFilterCallback) ListOfIncident {
	nl := NewListOfIncident()
	l.IterAll(func(v Incident, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}
