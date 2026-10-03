---
page_title: "routes.simple_route.incoming_port"
subcategory: "Load Balancing"
description: "Port match of the request can be a range or a specific port."
xcsh_docs: {"aliases": ["routes simple route incoming port"], "body_bytes": 4491, "body_sha256": "sha256:2a76270bbb414357cf9948b80c5776cefce38b3e59f2abfc0f45a38f14c060e3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2021321033003230-1330132333221000-2002303112001020-2221033213330222-2123322230122213-3213310100232300-3232300311302131-2032221321131323", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [{"anchor": "schema-routes--simple_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "incoming_port"], "schema_version": 1, "sections": [{"aliases": ["routes simple route incoming port no port match"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "incoming_port", "no_port_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes simple route incoming port port"], "anchor": "schema-routes--simple_route--incoming_port--port", "description": "Exclusive with Exact Port to match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "incoming_port", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["routes simple route incoming port port ranges"], "anchor": "schema-routes--simple_route--incoming_port--port_ranges", "description": "Exclusive with Port range to match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "incoming_port", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Port match of the request can be a range or a specific port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.incoming_port

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- routes.simple_route.incoming_port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/no_port_match/): complete subsection reference.

<a id="schema-routes--simple_route--incoming_port--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-routes--simple_route--incoming_port--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

## Next pages

- [routes.simple_route.incoming_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/no_port_match/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
