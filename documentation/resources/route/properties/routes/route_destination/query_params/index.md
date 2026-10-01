---
page_title: "routes.route_destination.query_params"
subcategory: ""
description: "routes.route_destination.query_params for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 3733, "body_sha256": "sha256:6cde208c6561cb6959c2d0341579c66a0b2c8275386d0f313c41fd157ce7ca77", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/query_params/index.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["routes", "route_destination", "query_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/query_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.query_params for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [routes.route_destination.query_params.remove_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/remove_all_params/)
- [routes.route_destination.query_params.retain_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/retain_all_params/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
