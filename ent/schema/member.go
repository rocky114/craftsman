package schema

import "github.com/facebook/ent"
import "github.com/facebook/ent/schema/field"

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
		field.String("nickname").Default(""),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
