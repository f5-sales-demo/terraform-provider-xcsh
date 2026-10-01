---
page_title: "ring_hash.hash_policy.cookie.ignore_samesite"
subcategory: "Load Balancing"
description: "ring_hash.hash_policy.cookie.ignore_samesite for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1640, "body_sha256": "sha256:2eebdbaa62c6316befe939a4abaf09f479c3adf5c595c4c8fc1218827dc09bae", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "path": "documentation/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_samesite/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["ring_hash", "hash_policy", "cookie", "ignore_samesite"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_samesite/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ring_hash.hash_policy.cookie.ignore_samesite for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy.cookie.ignore_samesite

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [ring_hash](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/)
- [ring_hash.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/)
- [ring_hash.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/)
- ring_hash.hash_policy.cookie.ignore_samesite

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
ignore_samesite = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ring_hash.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
