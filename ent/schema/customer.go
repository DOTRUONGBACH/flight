package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("username").Unique().NotEmpty().Annotations(entgql.OrderField("USERNAME")),
		field.String("password").NotEmpty().Annotations(entgql.OrderField("PASSWORD")),
		field.String("fullname").NotEmpty().Annotations(entgql.OrderField("FULLNAME")),
		field.Int("phone_number").Positive().Annotations(entgql.OrderField("PHONE_NUMBER")),
		field.String("email").NotEmpty().Annotations(entgql.OrderField("EMAIL")),
		field.Int("id_card").Positive().Annotations(entgql.OrderField("ID_CARD")),
		field.Int("member_card").Positive().Annotations(entgql.OrderField("MEMBER_CARD")),
		field.Time("date_of_birth").Annotations(entgql.OrderField("DATE_OF_BIRTH")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
