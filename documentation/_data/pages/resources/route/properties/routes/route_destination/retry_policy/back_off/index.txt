---
page_title: "routes.route_destination.retry_policy.back_off"
subcategory: ""
description: "Specifies parameters that control retry back off."
xcsh_docs: {"aliases": ["routes route destination retry policy back off"], "body_bytes": 3172, "body_sha256": "sha256:2ac7dd7079c295f707d0839ed78aed622213a9a2585fc4c03694cf2c3ca55f5b", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "path": "documentation/resources/route/properties/routes/route_destination/retry_policy/back_off/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3333100010113321-1102030132022101-1120023010023101-1211232223320201-0213303320303033-1213213133220120-2100010200020111-3302303223023103", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "retry_policy", "back_off"], "schema_version": 1, "sections": [{"aliases": ["base interval"], "anchor": "schema-routes--route_destination--retry_policy--back_off--base_interval", "description": "Specifies the base interval between retries in milliseconds.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off", "base_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["max interval"], "anchor": "schema-routes--route_destination--retry_policy--back_off--max_interval", "description": "Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The default is 10 times the base_interval.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off", "max_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/retry_policy/back_off/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies parameters that control retry back off.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The default is 10 times the base\_interval.

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

## Next pages

- [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
