---
page_title: "ring_hash.hash_policy.cookie.ignore_samesite"
subcategory: "Load Balancing"
description: "ring_hash.hash_policy.cookie.ignore_samesite for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1286, "body_sha256": "sha256:9a6d1e63b50df441991f92fb720a8282311bc1b7b496c21fbcd63b89ce243233", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_samesite", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "path": "docs/guides/resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie--ignore_samesite.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ring_hash", "hash_policy", "cookie", "ignore_samesite"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_samesite/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ring_hash.hash_policy.cookie.ignore_samesite for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy.cookie.ignore_samesite

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [ring_hash](resources--http_loadbalancer--properties--ring_hash.md)
- [ring_hash.hash_policy](resources--http_loadbalancer--properties--ring_hash--hash_policy.md)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie.md)
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

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
