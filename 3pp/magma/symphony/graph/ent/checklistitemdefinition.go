// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/checklistitemdefinition"
	"github.com/facebookincubator/symphony/graph/ent/workordertype"
)

// CheckListItemDefinition is the model entity for the CheckListItemDefinition schema.
type CheckListItemDefinition struct {
	config `gqlgen:"-" json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Index holds the value of the "index" field.
	Index int `json:"index,omitempty"`
	// EnumValues holds the value of the "enum_values" field.
	EnumValues *string `json:"enum_values,omitempty" gqlgen:"enumValues"`
	// HelpText holds the value of the "help_text" field.
	HelpText *string `json:"help_text,omitempty" gqlgen:"helpText"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckListItemDefinitionQuery when eager-loading is set.
	Edges              CheckListItemDefinitionEdges `json:"edges"`
	work_order_type_id *string
}

// CheckListItemDefinitionEdges holds the relations/edges for other nodes in the graph.
type CheckListItemDefinitionEdges struct {
	// WorkOrderType holds the value of the work_order_type edge.
	WorkOrderType *WorkOrderType
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkOrderTypeErr returns the WorkOrderType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckListItemDefinitionEdges) WorkOrderTypeErr() (*WorkOrderType, error) {
	if e.loadedTypes[0] {
		if e.WorkOrderType == nil {
			// The edge work_order_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workordertype.Label}
		}
		return e.WorkOrderType, nil
	}
	return nil, &NotLoadedError{edge: "work_order_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CheckListItemDefinition) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // title
		&sql.NullString{}, // type
		&sql.NullInt64{},  // index
		&sql.NullString{}, // enum_values
		&sql.NullString{}, // help_text
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*CheckListItemDefinition) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // work_order_type_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CheckListItemDefinition fields.
func (clid *CheckListItemDefinition) assignValues(values ...interface{}) error {
	if m, n := len(values), len(checklistitemdefinition.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	clid.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		clid.Title = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[1])
	} else if value.Valid {
		clid.Type = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field index", values[2])
	} else if value.Valid {
		clid.Index = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field enum_values", values[3])
	} else if value.Valid {
		clid.EnumValues = new(string)
		*clid.EnumValues = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field help_text", values[4])
	} else if value.Valid {
		clid.HelpText = new(string)
		*clid.HelpText = value.String
	}
	values = values[5:]
	if len(values) == len(checklistitemdefinition.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field work_order_type_id", value)
		} else if value.Valid {
			clid.work_order_type_id = new(string)
			*clid.work_order_type_id = strconv.FormatInt(value.Int64, 10)
		}
	}
	return nil
}

// QueryWorkOrderType queries the work_order_type edge of the CheckListItemDefinition.
func (clid *CheckListItemDefinition) QueryWorkOrderType() *WorkOrderTypeQuery {
	return (&CheckListItemDefinitionClient{clid.config}).QueryWorkOrderType(clid)
}

// Update returns a builder for updating this CheckListItemDefinition.
// Note that, you need to call CheckListItemDefinition.Unwrap() before calling this method, if this CheckListItemDefinition
// was returned from a transaction, and the transaction was committed or rolled back.
func (clid *CheckListItemDefinition) Update() *CheckListItemDefinitionUpdateOne {
	return (&CheckListItemDefinitionClient{clid.config}).UpdateOne(clid)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (clid *CheckListItemDefinition) Unwrap() *CheckListItemDefinition {
	tx, ok := clid.config.driver.(*txDriver)
	if !ok {
		panic("ent: CheckListItemDefinition is not a transactional entity")
	}
	clid.config.driver = tx.drv
	return clid
}

// String implements the fmt.Stringer.
func (clid *CheckListItemDefinition) String() string {
	var builder strings.Builder
	builder.WriteString("CheckListItemDefinition(")
	builder.WriteString(fmt.Sprintf("id=%v", clid.ID))
	builder.WriteString(", title=")
	builder.WriteString(clid.Title)
	builder.WriteString(", type=")
	builder.WriteString(clid.Type)
	builder.WriteString(", index=")
	builder.WriteString(fmt.Sprintf("%v", clid.Index))
	if v := clid.EnumValues; v != nil {
		builder.WriteString(", enum_values=")
		builder.WriteString(*v)
	}
	if v := clid.HelpText; v != nil {
		builder.WriteString(", help_text=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (clid *CheckListItemDefinition) id() int {
	id, _ := strconv.Atoi(clid.ID)
	return id
}

// CheckListItemDefinitions is a parsable slice of CheckListItemDefinition.
type CheckListItemDefinitions []*CheckListItemDefinition

func (clid CheckListItemDefinitions) config(cfg config) {
	for _i := range clid {
		clid[_i].config = cfg
	}
}
