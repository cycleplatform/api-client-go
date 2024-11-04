package main

import (
	"fmt"

	"github.com/pb33f/libopenapi/datamodel/high/base"
)

type SchemaTemplate struct {
	Name    string
	Comment string
}

type SchemaTemplateField struct {
	Name     string
	Comment  string
	Type     string
	Optional bool
}

func parseSchema(s *base.Schema) {
	fmt.Printf("-- FOUND SCHEMA: %s\n", getSchemaName(s))

	if s.Properties != nil {
		for p := s.Properties.Oldest(); p != nil; p = p.Next() {
			fmt.Printf("\t - %s", p.Key)
			if p.Value.IsReference() {
				ref, _ := p.Value.GetReferenceOrigin().Index.SearchIndexForReference(p.Value.GetReference())
				fmt.Printf(" (reference) %s\n", ref.Name)
			} else if len(p.Value.Schema().AnyOf) > 0 {
				fmt.Printf(" [mult]")
				for _, aof := range p.Value.Schema().AnyOf {
					fmt.Printf("\n\t\t-- %s", aof.Schema().Type)
				}
				fmt.Println()
			} else {
				fmt.Println("(type)", p.Value.Schema().Type)
			}
		}
	} else {
		fmt.Println("\t -- no dependencies")
	}
}

func getSchemaName(s *base.Schema) string {
	if s.Title != "" {
		return s.Title
	}

	return s.Description
}
