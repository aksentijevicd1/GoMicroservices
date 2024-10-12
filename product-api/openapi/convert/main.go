package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/invopop/yaml"
)

func main() {

	// Učitavanje Swagger 2.0 (openapi v2) fajla
	input, err := os.ReadFile("docs/swagger/swagger.yaml")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	fmt.Println("Swagger 2.0 file read successfully")

	// Parsiranje fajla u OpenAPI v2 strukturu
	var doc openapi2.T
	if err = yaml.Unmarshal(input, &doc); err != nil {
		fmt.Println("Error unmarshalling input file:", err)
		os.Exit(1)
	}

	fmt.Println("Swagger 2.0 file unmarshalled successfully")

	// Direktna konverzija u OpenAPI v3 bez validacije JSON roundtrip-a
	oa3, err := openapi2conv.ToV3(&doc)
	if err != nil {
		fmt.Println("Error converting to OpenAPI 3:", err)
		os.Exit(1)
	}

	fmt.Println("Conversion to OpenAPI 3 successful")

	// Dodavanje opcionalnih OpenAPI 3 podataka (primer: JWT autentifikacija)
	oa3.Security = openapi3.SecurityRequirements{
		{
			"bearerAuth": []string{},
		},
	}

	if oa3.Components.SecuritySchemes == nil {
		oa3.Components.SecuritySchemes = openapi3.SecuritySchemes{}
	}

	oa3.Components.SecuritySchemes["bearerAuth"] = &openapi3.SecuritySchemeRef{
		Value: &openapi3.SecurityScheme{
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	oa3.Servers = openapi3.Servers{
		&openapi3.Server{
			URL: "http://localhost:9090", // Tvoj lokalni server
		},
	}

	// Marshalling OpenAPI 3 dokumenta u YAML
	outputYAMLv3, err := yaml.Marshal(oa3)
	if err != nil {
		fmt.Println("Error marshalling OpenAPI 3 to YAML:", err)
		os.Exit(1)
	}

	// Sačuvaj konvertovani fajl
	err = os.WriteFile("openapi.yaml", outputYAMLv3, 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("OpenAPI 3 file written successfully")
}
