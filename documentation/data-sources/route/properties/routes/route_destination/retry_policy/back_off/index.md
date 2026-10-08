---
page_title: "routes.route_destination.retry_policy.back_off"
subcategory: ""
description: "Specifies parameters that control retry back off."
xcsh_docs: {"aliases": ["routes route destination retry policy back off"], "body_bytes": 2522, "body_sha256": "sha256:69c72b637d70cd5047660302bfbc50e5fbc1e2ae4b113362938b217da5ea61c7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy:back_off", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy", "path": "documentation/data-sources/route/properties/routes/route_destination/retry_policy/back_off/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3102121022302111-0213310320213223-3302022031001103-0002012330323031-3203310302312010-0223223000221300-0112101003313012-3310120012220312", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "retry_policy", "back_off"], "schema_version": 1, "sections": [{"aliases": ["routes route destination retry policy back off base interval"], "anchor": "schema-routes--route_destination--retry_policy--back_off--base_interval", "description": "Specifies the base interval between retries in milliseconds.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy:back_off", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off", "base_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["routes route destination retry policy back off max interval"], "anchor": "schema-routes--route_destination--retry_policy--back_off--max_interval", "description": "Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The default is 10 times the base_interval.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy:back_off", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off", "max_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/retry_policy/back_off/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Specifies parameters that control retry back off.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["routeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.retry_policy.back_off

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/retry_policy/)
- routes.route_destination.retry_policy.back_off

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-routes--route_destination--retry_policy--back_off--base_interval"></a>

### base_interval property

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Type: `"number"`. Computed.

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
