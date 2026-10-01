---
page_title: "routes.service_policy"
subcategory: ""
description: "routes.service_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1006, "body_sha256": "sha256:04123a8a680bbfab73a81010ac2aaf2b66923ef3d9b3af9b05204d1a10e7a626", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "child_ids": [], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "docs/guides/data-sources--route--properties--routes--service_policy.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "service_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/service_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.service_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.service_policy

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- routes.service_policy

<a id="section"></a>

Type: `"single"`. Computed.

ServicePolicy configuration details at route level.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_policy_choice": "[\"disable\"]"
}
```

## Direct properties

<a id="schema-routes--service_policy--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

## Next pages

- [routes](data-sources--route--properties--routes.md)
- [xcsh_route](../data-sources/route.md)
