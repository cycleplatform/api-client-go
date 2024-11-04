package main

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
)

func main() {
	spec, _ := os.ReadFile("../api-spec/dist/platform.yml")
	document, err := libopenapi.NewDocument(spec)

	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	v3Model, errors := document.BuildV3Model()

	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported",
			len(errors)))
	}

	// a, _ := v3Model.Index.SearchIndexForReference("#/components/schemas/ServerStatsCpuUsageTelemetry")
	// fmt.Println("TEST", a.Name)

	schemas := v3Model.Model.Components.Schemas

	for kv := schemas.Oldest(); kv != nil; kv = kv.Next() {
		fmt.Println("KEY", kv.Key)
		schema, err := kv.Value.BuildSchema()
		if err != nil {
			panic(err)
		}

		parseSchema(schema)
	}

	// paths := v3Model.Model.Paths.PathItems
	// for p := paths.Oldest(); p != nil; p = p.Next() {
	// 	fmt.Println(p.Key)
	// 	for o := p.Value.GetOperations().Oldest(); o != nil; o = o.Next() {
	// 		fmt.Printf("\t %s", strings.ToUpper(o.Key))
	// 		fmt.Printf("\t %s\n", o.Value.Summary)

	// 		o.Value.Parameters
	// 	}
	// }

	// for pair := paths.Oldest(); pair != nil; pair = pair.Next() {
	// 	if pair.Value.Get != nil {
	// 		for c := pair.Value.Get.Responses.Codes.Oldest(); c != nil; c = c.Next() {
	// 			if c.Value.Content == nil {
	// 				fmt.Printf("NOSDIOFJ %s %s\n", pair.Key, c.Key)
	// 				continue
	// 			}
	// 			for cn := c.Value.Content.Oldest(); cn != nil; cn = cn.Next() {
	// 				fmt.Printf("%s => %s => %s\n", pair.Key, c.Key, cn.Value.Schema.Schema().Title)
	// 			}
	// 		}
	// 	}

	// if pair.Value.Post != nil {
	// 	fmt.Printf("POST %s => %s\n", pair.Key, pair.Value.Post.OperationId)
	// }

	// fmt.Printf("%s => %s\n", pair.Key, pair.Value.Schema().Properties)
	// }

}
