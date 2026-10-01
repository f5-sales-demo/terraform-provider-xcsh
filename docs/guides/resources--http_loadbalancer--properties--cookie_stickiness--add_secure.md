---
page_title: "cookie_stickiness.add_secure"
subcategory: "Load Balancing"
description: "cookie_stickiness.add_secure for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1036, "body_sha256": "sha256:5a8c543a8ce095f971274bac76a6e5384adcace5b15079c346a302a38f1a72b9", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_secure", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:add_secure", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "path": "docs/guides/resources--http_loadbalancer--properties--cookie_stickiness--add_secure.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cookie_stickiness", "add_secure"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/cookie_stickiness/add_secure/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_stickiness.add_secure for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_stickiness.add_secure

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [cookie_stickiness](resources--http_loadbalancer--properties--cookie_stickiness.md)
- cookie_stickiness.add_secure

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
add_secure = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cookie_stickiness](resources--http_loadbalancer--properties--cookie_stickiness.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
