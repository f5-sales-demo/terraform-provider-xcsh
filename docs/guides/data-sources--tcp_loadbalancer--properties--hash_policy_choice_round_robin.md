---
page_title: "hash_policy_choice_round_robin"
subcategory: "Load Balancing"
description: "hash_policy_choice_round_robin for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 991, "body_sha256": "sha256:18916e0dcfbb16babc6d31b9914650beffd5ecdcadd4af0432e6f2413e0fd85f", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:hash_policy_choice_round_robin", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:hash_policy_choice_round_robin", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--hash_policy_choice_round_robin.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["hash_policy_choice_round_robin"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/hash_policy_choice_round_robin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "hash_policy_choice_round_robin for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hash_policy_choice_round_robin

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- hash_policy_choice_round_robin

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
