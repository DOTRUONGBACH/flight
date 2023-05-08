// Code generated by ent, DO NOT EDIT.

package ent

import (
	"customer-service/ent/customer"
	"customer-service/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescUsername is the schema descriptor for username field.
	customerDescUsername := customerFields[1].Descriptor()
	// customer.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	customer.UsernameValidator = customerDescUsername.Validators[0].(func(string) error)
	// customerDescPassword is the schema descriptor for password field.
	customerDescPassword := customerFields[2].Descriptor()
	// customer.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	customer.PasswordValidator = customerDescPassword.Validators[0].(func(string) error)
	// customerDescFullname is the schema descriptor for fullname field.
	customerDescFullname := customerFields[3].Descriptor()
	// customer.FullnameValidator is a validator for the "fullname" field. It is called by the builders before save.
	customer.FullnameValidator = customerDescFullname.Validators[0].(func(string) error)
	// customerDescPhoneNumber is the schema descriptor for phone_number field.
	customerDescPhoneNumber := customerFields[4].Descriptor()
	// customer.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	customer.PhoneNumberValidator = customerDescPhoneNumber.Validators[0].(func(int) error)
	// customerDescEmail is the schema descriptor for email field.
	customerDescEmail := customerFields[5].Descriptor()
	// customer.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	customer.EmailValidator = customerDescEmail.Validators[0].(func(string) error)
	// customerDescIDCard is the schema descriptor for id_card field.
	customerDescIDCard := customerFields[6].Descriptor()
	// customer.IDCardValidator is a validator for the "id_card" field. It is called by the builders before save.
	customer.IDCardValidator = customerDescIDCard.Validators[0].(func(int) error)
	// customerDescMemberCard is the schema descriptor for member_card field.
	customerDescMemberCard := customerFields[7].Descriptor()
	// customer.MemberCardValidator is a validator for the "member_card" field. It is called by the builders before save.
	customer.MemberCardValidator = customerDescMemberCard.Validators[0].(func(int) error)
	// customerDescCreatedAt is the schema descriptor for created_at field.
	customerDescCreatedAt := customerFields[9].Descriptor()
	// customer.DefaultCreatedAt holds the default value on creation for the created_at field.
	customer.DefaultCreatedAt = customerDescCreatedAt.Default.(func() time.Time)
	// customerDescUpdatedAt is the schema descriptor for updated_at field.
	customerDescUpdatedAt := customerFields[10].Descriptor()
	// customer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customer.DefaultUpdatedAt = customerDescUpdatedAt.Default.(func() time.Time)
	// customer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customer.UpdateDefaultUpdatedAt = customerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customerDescID is the schema descriptor for id field.
	customerDescID := customerFields[0].Descriptor()
	// customer.DefaultID holds the default value on creation for the id field.
	customer.DefaultID = customerDescID.Default.(func() uuid.UUID)
}
