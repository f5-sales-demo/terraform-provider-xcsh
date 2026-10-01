---
page_title: "ring_hash"
subcategory: "Load Balancing"
description: "ring_hash for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 988, "body_sha256": "sha256:fc8caf97b95f70c0de189f7ad34cd3829fa1677906715b90f36e0fa99b4a8eab", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--ring_hash.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ring_hash"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/ring_hash/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ring_hash for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- ring_hash

<a id="section"></a>

Type: `"single"`. Computed.

Hash Policy List. List of hash policy rules.

Upstream description:

List of hash policy rules.

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

- [hash_policy](data-sources--http_loadbalancer--properties--ring_hash--hash_policy.md): complete subsection reference.

## Next pages

- [ring_hash.hash_policy](data-sources--http_loadbalancer--properties--ring_hash--hash_policy.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
