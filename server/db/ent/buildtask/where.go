// Code generated by entc, DO NOT EDIT.

package buildtask

import (
	"github.com/bishopfox/sliver/server/db/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GUID applies equality check predicate on the "guid" field. It's identical to GUIDEQ.
func GUID(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGUID), v))
	})
}

// GUIDEQ applies the EQ predicate on the "guid" field.
func GUIDEQ(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGUID), v))
	})
}

// GUIDNEQ applies the NEQ predicate on the "guid" field.
func GUIDNEQ(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGUID), v))
	})
}

// GUIDIn applies the In predicate on the "guid" field.
func GUIDIn(vs ...uuid.UUID) predicate.BuildTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGUID), v...))
	})
}

// GUIDNotIn applies the NotIn predicate on the "guid" field.
func GUIDNotIn(vs ...uuid.UUID) predicate.BuildTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGUID), v...))
	})
}

// GUIDGT applies the GT predicate on the "guid" field.
func GUIDGT(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGUID), v))
	})
}

// GUIDGTE applies the GTE predicate on the "guid" field.
func GUIDGTE(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGUID), v))
	})
}

// GUIDLT applies the LT predicate on the "guid" field.
func GUIDLT(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGUID), v))
	})
}

// GUIDLTE applies the LTE predicate on the "guid" field.
func GUIDLTE(v uuid.UUID) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGUID), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.BuildTask) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.BuildTask) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
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
func Not(p predicate.BuildTask) predicate.BuildTask {
	return predicate.BuildTask(func(s *sql.Selector) {
		p(s.Not())
	})
}