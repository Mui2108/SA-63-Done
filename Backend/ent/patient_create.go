// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/panupong/app/ent/gender"
	"github.com/panupong/app/ent/job"
	"github.com/panupong/app/ent/patient"
	"github.com/panupong/app/ent/title"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetCardID sets the Card_id field.
func (pc *PatientCreate) SetCardID(s string) *PatientCreate {
	pc.mutation.SetCardID(s)
	return pc
}

// SetFirstName sets the First_name field.
func (pc *PatientCreate) SetFirstName(s string) *PatientCreate {
	pc.mutation.SetFirstName(s)
	return pc
}

// SetLastName sets the Last_name field.
func (pc *PatientCreate) SetLastName(s string) *PatientCreate {
	pc.mutation.SetLastName(s)
	return pc
}

// SetAllergic sets the Allergic field.
func (pc *PatientCreate) SetAllergic(s string) *PatientCreate {
	pc.mutation.SetAllergic(s)
	return pc
}

// SetAddress sets the Address field.
func (pc *PatientCreate) SetAddress(s string) *PatientCreate {
	pc.mutation.SetAddress(s)
	return pc
}

// SetAge sets the Age field.
func (pc *PatientCreate) SetAge(i int) *PatientCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetBirthday sets the Birthday field.
func (pc *PatientCreate) SetBirthday(t time.Time) *PatientCreate {
	pc.mutation.SetBirthday(t)
	return pc
}

// SetGenderID sets the gender edge to Gender by id.
func (pc *PatientCreate) SetGenderID(id int) *PatientCreate {
	pc.mutation.SetGenderID(id)
	return pc
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pc *PatientCreate) SetNillableGenderID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetGenderID(*id)
	}
	return pc
}

// SetGender sets the gender edge to Gender.
func (pc *PatientCreate) SetGender(g *Gender) *PatientCreate {
	return pc.SetGenderID(g.ID)
}

// SetTitleID sets the title edge to Title by id.
func (pc *PatientCreate) SetTitleID(id int) *PatientCreate {
	pc.mutation.SetTitleID(id)
	return pc
}

// SetNillableTitleID sets the title edge to Title by id if the given value is not nil.
func (pc *PatientCreate) SetNillableTitleID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetTitleID(*id)
	}
	return pc
}

// SetTitle sets the title edge to Title.
func (pc *PatientCreate) SetTitle(t *Title) *PatientCreate {
	return pc.SetTitleID(t.ID)
}

// SetJobID sets the job edge to Job by id.
func (pc *PatientCreate) SetJobID(id int) *PatientCreate {
	pc.mutation.SetJobID(id)
	return pc
}

// SetNillableJobID sets the job edge to Job by id if the given value is not nil.
func (pc *PatientCreate) SetNillableJobID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetJobID(*id)
	}
	return pc
}

// SetJob sets the job edge to Job.
func (pc *PatientCreate) SetJob(j *Job) *PatientCreate {
	return pc.SetJobID(j.ID)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	if err := pc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientCreate) preSave() error {
	if _, ok := pc.mutation.CardID(); !ok {
		return &ValidationError{Name: "Card_id", err: errors.New("ent: missing required field \"Card_id\"")}
	}
	if v, ok := pc.mutation.CardID(); ok {
		if err := patient.CardIDValidator(v); err != nil {
			return &ValidationError{Name: "Card_id", err: fmt.Errorf("ent: validator failed for field \"Card_id\": %w", err)}
		}
	}
	if _, ok := pc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "First_name", err: errors.New("ent: missing required field \"First_name\"")}
	}
	if v, ok := pc.mutation.FirstName(); ok {
		if err := patient.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "First_name", err: fmt.Errorf("ent: validator failed for field \"First_name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.LastName(); !ok {
		return &ValidationError{Name: "Last_name", err: errors.New("ent: missing required field \"Last_name\"")}
	}
	if v, ok := pc.mutation.LastName(); ok {
		if err := patient.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "Last_name", err: fmt.Errorf("ent: validator failed for field \"Last_name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Allergic(); !ok {
		return &ValidationError{Name: "Allergic", err: errors.New("ent: missing required field \"Allergic\"")}
	}
	if v, ok := pc.mutation.Allergic(); ok {
		if err := patient.AllergicValidator(v); err != nil {
			return &ValidationError{Name: "Allergic", err: fmt.Errorf("ent: validator failed for field \"Allergic\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Address(); !ok {
		return &ValidationError{Name: "Address", err: errors.New("ent: missing required field \"Address\"")}
	}
	if v, ok := pc.mutation.Address(); ok {
		if err := patient.AddressValidator(v); err != nil {
			return &ValidationError{Name: "Address", err: fmt.Errorf("ent: validator failed for field \"Address\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "Age", err: errors.New("ent: missing required field \"Age\"")}
	}
	if _, ok := pc.mutation.Birthday(); !ok {
		return &ValidationError{Name: "Birthday", err: errors.New("ent: missing required field \"Birthday\"")}
	}
	return nil
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CardID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldCardID,
		})
		pa.CardID = value
	}
	if value, ok := pc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldFirstName,
		})
		pa.FirstName = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldLastName,
		})
		pa.LastName = value
	}
	if value, ok := pc.mutation.Allergic(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAllergic,
		})
		pa.Allergic = value
	}
	if value, ok := pc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAddress,
		})
		pa.Address = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
		pa.Age = value
	}
	if value, ok := pc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patient.FieldBirthday,
		})
		pa.Birthday = value
	}
	if nodes := pc.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.TitleTable,
			Columns: []string{patient.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.JobTable,
			Columns: []string{patient.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}

// PatientCreateBulk is the builder for creating a bulk of Patient entities.
type PatientCreateBulk struct {
	config
	builders []*PatientCreate
}

// Save creates the Patient entities in the database.
func (pcb *PatientCreateBulk) Save(ctx context.Context) ([]*Patient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Patient, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*PatientMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *PatientCreateBulk) SaveX(ctx context.Context) []*Patient {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}