---
page_title: "routes.route_destination.query_params"
subcategory: ""
description: "routes.route_destination.query_params for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 3136, "body_sha256": "sha256:f064ba5582dc2b3e1c7f4d81782ed193cec67efaf8f8b081fd6a7c8ee15d79e2", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "docs/guides/resources--route--properties--routes--route_destination--query_params.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "query_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/query_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.query_params for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination.query_params

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
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

- [remove_all_params](resources--route--properties--routes--route_destination--query_params--remove_all_params.md): complete subsection reference.

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

- [retain_all_params](resources--route--properties--routes--route_destination--query_params--retain_all_params.md): complete subsection reference.

## Next pages

- [routes.route_destination.query_params.remove_all_params](resources--route--properties--routes--route_destination--query_params--remove_all_params.md)
- [routes.route_destination.query_params.retain_all_params](resources--route--properties--routes--route_destination--query_params--retain_all_params.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [xcsh_route](../resources/route.md)
