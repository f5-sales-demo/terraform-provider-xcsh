---
page_title: "jwt_validation.token_location"
subcategory: "Load Balancing"
description: "jwt_validation.token_location for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1601, "body_sha256": "sha256:867b1490c27a09e3bb1b80373eee578c5a1c773345ee47a113a3333a0bbcc235", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:token_location:bearer_token"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:token_location", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/cdn_loadbalancer/properties/jwt_validation/token_location/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["jwt_validation", "token_location"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/token_location/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.token_location for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.token_location

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- jwt_validation.token_location

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

## Direct properties

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/token_location/bearer_token/): complete subsection reference.

## Next pages

- [jwt_validation.token_location.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/token_location/bearer_token/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
