---
page_title: "snat_pool"
subcategory: "Networking"
description: "snat_pool for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1178, "body_sha256": "sha256:f2640691be71bf24fd918198eb8211477943c640ee565651a39a5af36e0c51c3", "canonical_id": "xcsh-docs:data-sources:endpoint:properties:snat_pool", "child_ids": ["xcsh-docs:data-sources:endpoint:properties:snat_pool:no_snat_pool", "xcsh-docs:data-sources:endpoint:properties:snat_pool:snat_pool"], "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:properties:snat_pool", "parent_id": "xcsh-docs:data-sources:endpoint:reference", "path": "docs/guides/data-sources--endpoint--properties--snat_pool.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "snat_pool for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# snat_pool

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md)
- [Property reference](data-sources--endpoint--reference.md)
- snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

## Direct properties

- [no_snat_pool](data-sources--endpoint--properties--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](data-sources--endpoint--properties--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [snat_pool.no_snat_pool](data-sources--endpoint--properties--snat_pool--no_snat_pool.md)
- [snat_pool.snat_pool](data-sources--endpoint--properties--snat_pool--snat_pool.md)
- [Property reference](data-sources--endpoint--reference.md)
- [xcsh_endpoint](../data-sources/endpoint.md)
