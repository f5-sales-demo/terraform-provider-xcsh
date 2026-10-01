---
page_title: "cookie_stickiness.ignore_httponly"
subcategory: "Load Balancing"
description: "cookie_stickiness.ignore_httponly for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1077, "body_sha256": "sha256:982581fb89b17657b1c414defea7a375c8194587b1bf0b3ae705d5255b355051", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_httponly", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness:ignore_httponly", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:cookie_stickiness", "path": "docs/guides/resources--http_loadbalancer--properties--cookie_stickiness--ignore_httponly.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cookie_stickiness", "ignore_httponly"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/cookie_stickiness/ignore_httponly/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_stickiness.ignore_httponly for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_stickiness.ignore_httponly

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [cookie_stickiness](resources--http_loadbalancer--properties--cookie_stickiness.md)
- cookie_stickiness.ignore_httponly

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cookie_stickiness](resources--http_loadbalancer--properties--cookie_stickiness.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
