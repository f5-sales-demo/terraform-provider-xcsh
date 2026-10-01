---
page_title: "routes.route_destination.mirror_policy.percent"
subcategory: ""
description: "routes.route_destination.mirror_policy.percent for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2929, "body_sha256": "sha256:0b8ae0676e632ec7f815bbba7adbaa0fa21936cb8a0a644a56ab11e9c90a6d5b", "child_ids": [], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy:percent", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy", "path": "documentation/data-sources/route/properties/routes/route_destination/mirror_policy/percent/index.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy", "percent"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/mirror_policy/percent/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.mirror_policy.percent for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.mirror_policy.percent

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [routes.route_destination.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/mirror_policy/)
- routes.route_destination.mirror_policy.percent

<a id="section"></a>

Type: `"single"`. Computed.

Fraction used where sampling percentages are needed. Example sampled requests.

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

<a id="schema-routes--route_destination--mirror_policy--percent--denominator"></a>

### denominator property

Type: `"string"`. Computed.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Upstream description:

Denominator used in fraction where sampling percentages are needed. Example sampled requests

Use hundred as denominator Use ten thousand as denominator Use million as denominator.

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--route_destination--mirror_policy--percent--numerator"></a>

### numerator property

Type: `"number"`. Computed.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [routes.route_destination.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/mirror_policy/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
