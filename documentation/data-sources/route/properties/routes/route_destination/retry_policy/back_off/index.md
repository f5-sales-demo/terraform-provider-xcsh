---
page_title: "routes.route_destination.retry_policy.back_off"
subcategory: ""
description: "routes.route_destination.retry_policy.back_off for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2947, "body_sha256": "sha256:b78770a92ef9768e51f06eecb4d293eb5e0d37a8e76f39e4de7fac8466d80edc", "child_ids": [], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy:back_off", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy", "path": "documentation/data-sources/route/properties/routes/route_destination/retry_policy/back_off/index.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes", "route_destination", "retry_policy", "back_off"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/retry_policy/back_off/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.retry_policy.back_off for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Type: `"number"`. Computed.

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

- [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/retry_policy/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
