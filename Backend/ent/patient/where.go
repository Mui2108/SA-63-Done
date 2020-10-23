// Code generated by entc, DO NOT EDIT.

package patient

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/panupong/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CardID applies equality check predicate on the "Card_id" field. It's identical to CardIDEQ.
func CardID(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardID), v))
	})
}

// FirstName applies equality check predicate on the "First_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// LastName applies equality check predicate on the "Last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// Allergic applies equality check predicate on the "Allergic" field. It's identical to AllergicEQ.
func Allergic(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllergic), v))
	})
}

// Address applies equality check predicate on the "Address" field. It's identical to AddressEQ.
func Address(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Age applies equality check predicate on the "Age" field. It's identical to AgeEQ.
func Age(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// Birthday applies equality check predicate on the "Birthday" field. It's identical to BirthdayEQ.
func Birthday(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthday), v))
	})
}

// CardIDEQ applies the EQ predicate on the "Card_id" field.
func CardIDEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardID), v))
	})
}

// CardIDNEQ applies the NEQ predicate on the "Card_id" field.
func CardIDNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCardID), v))
	})
}

// CardIDIn applies the In predicate on the "Card_id" field.
func CardIDIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCardID), v...))
	})
}

// CardIDNotIn applies the NotIn predicate on the "Card_id" field.
func CardIDNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCardID), v...))
	})
}

// CardIDGT applies the GT predicate on the "Card_id" field.
func CardIDGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCardID), v))
	})
}

// CardIDGTE applies the GTE predicate on the "Card_id" field.
func CardIDGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCardID), v))
	})
}

// CardIDLT applies the LT predicate on the "Card_id" field.
func CardIDLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCardID), v))
	})
}

// CardIDLTE applies the LTE predicate on the "Card_id" field.
func CardIDLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCardID), v))
	})
}

// CardIDContains applies the Contains predicate on the "Card_id" field.
func CardIDContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCardID), v))
	})
}

// CardIDHasPrefix applies the HasPrefix predicate on the "Card_id" field.
func CardIDHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCardID), v))
	})
}

// CardIDHasSuffix applies the HasSuffix predicate on the "Card_id" field.
func CardIDHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCardID), v))
	})
}

// CardIDEqualFold applies the EqualFold predicate on the "Card_id" field.
func CardIDEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCardID), v))
	})
}

// CardIDContainsFold applies the ContainsFold predicate on the "Card_id" field.
func CardIDContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCardID), v))
	})
}

// FirstNameEQ applies the EQ predicate on the "First_name" field.
func FirstNameEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// FirstNameNEQ applies the NEQ predicate on the "First_name" field.
func FirstNameNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstName), v))
	})
}

// FirstNameIn applies the In predicate on the "First_name" field.
func FirstNameIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirstName), v...))
	})
}

// FirstNameNotIn applies the NotIn predicate on the "First_name" field.
func FirstNameNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirstName), v...))
	})
}

// FirstNameGT applies the GT predicate on the "First_name" field.
func FirstNameGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstName), v))
	})
}

// FirstNameGTE applies the GTE predicate on the "First_name" field.
func FirstNameGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstName), v))
	})
}

// FirstNameLT applies the LT predicate on the "First_name" field.
func FirstNameLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstName), v))
	})
}

// FirstNameLTE applies the LTE predicate on the "First_name" field.
func FirstNameLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstName), v))
	})
}

// FirstNameContains applies the Contains predicate on the "First_name" field.
func FirstNameContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstName), v))
	})
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "First_name" field.
func FirstNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstName), v))
	})
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "First_name" field.
func FirstNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstName), v))
	})
}

// FirstNameEqualFold applies the EqualFold predicate on the "First_name" field.
func FirstNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstName), v))
	})
}

// FirstNameContainsFold applies the ContainsFold predicate on the "First_name" field.
func FirstNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstName), v))
	})
}

// LastNameEQ applies the EQ predicate on the "Last_name" field.
func LastNameEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// LastNameNEQ applies the NEQ predicate on the "Last_name" field.
func LastNameNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastName), v))
	})
}

// LastNameIn applies the In predicate on the "Last_name" field.
func LastNameIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastName), v...))
	})
}

// LastNameNotIn applies the NotIn predicate on the "Last_name" field.
func LastNameNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastName), v...))
	})
}

// LastNameGT applies the GT predicate on the "Last_name" field.
func LastNameGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastName), v))
	})
}

// LastNameGTE applies the GTE predicate on the "Last_name" field.
func LastNameGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastName), v))
	})
}

// LastNameLT applies the LT predicate on the "Last_name" field.
func LastNameLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastName), v))
	})
}

// LastNameLTE applies the LTE predicate on the "Last_name" field.
func LastNameLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastName), v))
	})
}

// LastNameContains applies the Contains predicate on the "Last_name" field.
func LastNameContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastName), v))
	})
}

// LastNameHasPrefix applies the HasPrefix predicate on the "Last_name" field.
func LastNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastName), v))
	})
}

// LastNameHasSuffix applies the HasSuffix predicate on the "Last_name" field.
func LastNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastName), v))
	})
}

// LastNameEqualFold applies the EqualFold predicate on the "Last_name" field.
func LastNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastName), v))
	})
}

// LastNameContainsFold applies the ContainsFold predicate on the "Last_name" field.
func LastNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastName), v))
	})
}

// AllergicEQ applies the EQ predicate on the "Allergic" field.
func AllergicEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllergic), v))
	})
}

// AllergicNEQ applies the NEQ predicate on the "Allergic" field.
func AllergicNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAllergic), v))
	})
}

// AllergicIn applies the In predicate on the "Allergic" field.
func AllergicIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAllergic), v...))
	})
}

// AllergicNotIn applies the NotIn predicate on the "Allergic" field.
func AllergicNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAllergic), v...))
	})
}

// AllergicGT applies the GT predicate on the "Allergic" field.
func AllergicGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAllergic), v))
	})
}

// AllergicGTE applies the GTE predicate on the "Allergic" field.
func AllergicGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAllergic), v))
	})
}

// AllergicLT applies the LT predicate on the "Allergic" field.
func AllergicLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAllergic), v))
	})
}

// AllergicLTE applies the LTE predicate on the "Allergic" field.
func AllergicLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAllergic), v))
	})
}

// AllergicContains applies the Contains predicate on the "Allergic" field.
func AllergicContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAllergic), v))
	})
}

// AllergicHasPrefix applies the HasPrefix predicate on the "Allergic" field.
func AllergicHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAllergic), v))
	})
}

// AllergicHasSuffix applies the HasSuffix predicate on the "Allergic" field.
func AllergicHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAllergic), v))
	})
}

// AllergicEqualFold applies the EqualFold predicate on the "Allergic" field.
func AllergicEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAllergic), v))
	})
}

// AllergicContainsFold applies the ContainsFold predicate on the "Allergic" field.
func AllergicContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAllergic), v))
	})
}

// AddressEQ applies the EQ predicate on the "Address" field.
func AddressEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "Address" field.
func AddressNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "Address" field.
func AddressIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "Address" field.
func AddressNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "Address" field.
func AddressGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "Address" field.
func AddressGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "Address" field.
func AddressLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "Address" field.
func AddressLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "Address" field.
func AddressContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "Address" field.
func AddressHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "Address" field.
func AddressHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "Address" field.
func AddressEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "Address" field.
func AddressContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// AgeEQ applies the EQ predicate on the "Age" field.
func AgeEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "Age" field.
func AgeNEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "Age" field.
func AgeIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "Age" field.
func AgeNotIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "Age" field.
func AgeGT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "Age" field.
func AgeGTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "Age" field.
func AgeLT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "Age" field.
func AgeLTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// BirthdayEQ applies the EQ predicate on the "Birthday" field.
func BirthdayEQ(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthday), v))
	})
}

// BirthdayNEQ applies the NEQ predicate on the "Birthday" field.
func BirthdayNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBirthday), v))
	})
}

// BirthdayIn applies the In predicate on the "Birthday" field.
func BirthdayIn(vs ...time.Time) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBirthday), v...))
	})
}

// BirthdayNotIn applies the NotIn predicate on the "Birthday" field.
func BirthdayNotIn(vs ...time.Time) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBirthday), v...))
	})
}

// BirthdayGT applies the GT predicate on the "Birthday" field.
func BirthdayGT(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBirthday), v))
	})
}

// BirthdayGTE applies the GTE predicate on the "Birthday" field.
func BirthdayGTE(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBirthday), v))
	})
}

// BirthdayLT applies the LT predicate on the "Birthday" field.
func BirthdayLT(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBirthday), v))
	})
}

// BirthdayLTE applies the LTE predicate on the "Birthday" field.
func BirthdayLTE(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBirthday), v))
	})
}

// HasGender applies the HasEdge predicate on the "gender" edge.
func HasGender() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenderWith applies the HasEdge predicate on the "gender" edge with a given conditions (other predicates).
func HasGenderWith(preds ...predicate.Gender) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTitle applies the HasEdge predicate on the "title" edge.
func HasTitle() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTitleWith applies the HasEdge predicate on the "title" edge with a given conditions (other predicates).
func HasTitleWith(preds ...predicate.Title) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasJob applies the HasEdge predicate on the "job" edge.
func HasJob() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JobTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobWith applies the HasEdge predicate on the "job" edge with a given conditions (other predicates).
func HasJobWith(preds ...predicate.Job) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JobInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		p(s.Not())
	})
}
