---
page_title: "routes.route_destination.query_params"
subcategory: ""
description: "Handling of incoming query parameters in simple route."
xcsh_docs: {"aliases": ["routes route destination query params"], "body_bytes": 3063, "body_sha256": "sha256:c3b34abd1ff2f8b5c3e5fe4566545e55246ba66996eab1a5f969b20740a813e7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/query_params/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0210013123232110-1101203121320003-0232332230203020-2203202001132332-1123010132101312-1300303103231023-2212021131213210-0023212010013031", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "schema-routes--route_destination--query_params--replace_params", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--query_params--replace_params", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "query_params"], "schema_version": 1, "sections": [{"aliases": ["routes route destination query params remove all params"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "query_params", "remove_all_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination query params replace params"], "anchor": "schema-routes--route_destination--query_params--replace_params", "description": "Exclusive with", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "query_params", "replace_params"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination query params retain all params"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "query_params", "retain_all_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/query_params/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Handling of incoming query parameters in simple route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["routeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.query_params

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.query_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Handling of incoming query parameters in simple route.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]"
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [remove_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/remove_all_params/): complete subsection reference.

<a id="schema-routes--route_destination--query_params--replace_params"></a>

### replace_params property

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [retain_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/retain_all_params/): complete subsection reference.
