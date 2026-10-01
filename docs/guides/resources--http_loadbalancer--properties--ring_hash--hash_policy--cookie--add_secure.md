---
page_title: "ring_hash.hash_policy.cookie.add_secure"
subcategory: "Load Balancing"
description: "ring_hash.hash_policy.cookie.add_secure for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1271, "body_sha256": "sha256:6d6e758eb228751e176074d8acf0f5fe75eb31142726d4b8c8ece356831a1340", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_secure", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:add_secure", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "path": "docs/guides/resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie--add_secure.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ring_hash", "hash_policy", "cookie", "add_secure"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/add_secure/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ring_hash.hash_policy.cookie.add_secure for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy.cookie.add_secure

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [ring_hash](resources--http_loadbalancer--properties--ring_hash.md)
- [ring_hash.hash_policy](resources--http_loadbalancer--properties--ring_hash--hash_policy.md)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie.md)
- ring_hash.hash_policy.cookie.add_secure

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

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
