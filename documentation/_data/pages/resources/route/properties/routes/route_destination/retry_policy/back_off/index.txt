---
page_title: "routes.route_destination.retry_policy.back_off"
subcategory: ""
description: "Specifies parameters that control retry back off."
xcsh_docs: {"aliases": ["routes route destination retry policy back off"], "body_bytes": 2783, "body_sha256": "sha256:7c2934ed03795c9b8072cb30c2364a7aa71cf92c1014e577ee3fb3348c47ff3f", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "path": "documentation/resources/route/properties/routes/route_destination/retry_policy/back_off/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3333100010113321-1102030132022101-1120023010023101-1211232223320201-0213303320303033-1213213133220120-2100010200020111-3302303223023103", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "retry_policy", "back_off"], "schema_version": 1, "sections": [{"aliases": ["routes route destination retry policy back off base interval"], "anchor": "schema-routes--route_destination--retry_policy--back_off--base_interval", "description": "Specifies the base interval between retries in milliseconds.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off", "base_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["routes route destination retry policy back off max interval"], "anchor": "schema-routes--route_destination--retry_policy--back_off--max_interval", "description": "Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The default is 10 times the base_interval.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off", "max_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/retry_policy/back_off/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Specifies parameters that control retry back off.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["routeCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.retry_policy.back_off

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/)
- routes.route_destination.retry_policy.back_off

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies parameters that control retry back off.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
back_off {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--route_destination--retry_policy--back_off--base_interval"></a>

### base_interval property

Type: `"number"`. Optional.

Specifies the base interval between retries in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="schema-routes--route_destination--retry_policy--back_off--max_interval"></a>

### max_interval property

Type: `"number"`. Optional.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Additional upstream details:

The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
