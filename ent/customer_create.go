// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer-service/ent/customer"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (cc *CustomerCreate) SetUsername(s string) *CustomerCreate {
	cc.mutation.SetUsername(s)
	return cc
}

// SetPassword sets the "password" field.
func (cc *CustomerCreate) SetPassword(s string) *CustomerCreate {
	cc.mutation.SetPassword(s)
	return cc
}

// SetFullname sets the "fullname" field.
func (cc *CustomerCreate) SetFullname(s string) *CustomerCreate {
	cc.mutation.SetFullname(s)
	return cc
}

// SetPhoneNumber sets the "phone_number" field.
func (cc *CustomerCreate) SetPhoneNumber(i int) *CustomerCreate {
	cc.mutation.SetPhoneNumber(i)
	return cc
}

// SetEmail sets the "email" field.
func (cc *CustomerCreate) SetEmail(s string) *CustomerCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetIDCard sets the "id_card" field.
func (cc *CustomerCreate) SetIDCard(i int) *CustomerCreate {
	cc.mutation.SetIDCard(i)
	return cc
}

// SetMemberCard sets the "member_card" field.
func (cc *CustomerCreate) SetMemberCard(i int) *CustomerCreate {
	cc.mutation.SetMemberCard(i)
	return cc
}

// SetDateOfBirth sets the "date_of_birth" field.
func (cc *CustomerCreate) SetDateOfBirth(t time.Time) *CustomerCreate {
	cc.mutation.SetDateOfBirth(t)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CustomerCreate) SetCreatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableCreatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CustomerCreate) SetUpdatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUpdatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CustomerCreate) SetID(u uuid.UUID) *CustomerCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableID(u *uuid.UUID) *CustomerCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := customer.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := customer.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := customer.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Customer.username"`)}
	}
	if v, ok := cc.mutation.Username(); ok {
		if err := customer.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Customer.username": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Customer.password"`)}
	}
	if v, ok := cc.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Customer.password": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Fullname(); !ok {
		return &ValidationError{Name: "fullname", err: errors.New(`ent: missing required field "Customer.fullname"`)}
	}
	if v, ok := cc.mutation.Fullname(); ok {
		if err := customer.FullnameValidator(v); err != nil {
			return &ValidationError{Name: "fullname", err: fmt.Errorf(`ent: validator failed for field "Customer.fullname": %w`, err)}
		}
	}
	if _, ok := cc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "Customer.phone_number"`)}
	}
	if v, ok := cc.mutation.PhoneNumber(); ok {
		if err := customer.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Customer.phone_number": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Customer.email"`)}
	}
	if v, ok := cc.mutation.Email(); ok {
		if err := customer.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Customer.email": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IDCard(); !ok {
		return &ValidationError{Name: "id_card", err: errors.New(`ent: missing required field "Customer.id_card"`)}
	}
	if v, ok := cc.mutation.IDCard(); ok {
		if err := customer.IDCardValidator(v); err != nil {
			return &ValidationError{Name: "id_card", err: fmt.Errorf(`ent: validator failed for field "Customer.id_card": %w`, err)}
		}
	}
	if _, ok := cc.mutation.MemberCard(); !ok {
		return &ValidationError{Name: "member_card", err: errors.New(`ent: missing required field "Customer.member_card"`)}
	}
	if v, ok := cc.mutation.MemberCard(); ok {
		if err := customer.MemberCardValidator(v); err != nil {
			return &ValidationError{Name: "member_card", err: fmt.Errorf(`ent: validator failed for field "Customer.member_card": %w`, err)}
		}
	}
	if _, ok := cc.mutation.DateOfBirth(); !ok {
		return &ValidationError{Name: "date_of_birth", err: errors.New(`ent: missing required field "Customer.date_of_birth"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Customer.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Customer.updated_at"`)}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(customer.Table, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Username(); ok {
		_spec.SetField(customer.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := cc.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := cc.mutation.Fullname(); ok {
		_spec.SetField(customer.FieldFullname, field.TypeString, value)
		_node.Fullname = value
	}
	if value, ok := cc.mutation.PhoneNumber(); ok {
		_spec.SetField(customer.FieldPhoneNumber, field.TypeInt, value)
		_node.PhoneNumber = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cc.mutation.IDCard(); ok {
		_spec.SetField(customer.FieldIDCard, field.TypeInt, value)
		_node.IDCard = value
	}
	if value, ok := cc.mutation.MemberCard(); ok {
		_spec.SetField(customer.FieldMemberCard, field.TypeInt, value)
		_node.MemberCard = value
	}
	if value, ok := cc.mutation.DateOfBirth(); ok {
		_spec.SetField(customer.FieldDateOfBirth, field.TypeTime, value)
		_node.DateOfBirth = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
