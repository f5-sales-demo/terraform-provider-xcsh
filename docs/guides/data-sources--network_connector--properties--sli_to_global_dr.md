---
page_title: "sli_to_global_dr"
subcategory: "Networking"
description: "sli_to_global_dr for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1500, "body_sha256": "sha256:e2b1efdca98ea0240672e2d987c54ab8d17153ebf767a68ff65d2f7e0e5ba049", "canonical_id": "xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "docs/guides/data-sources--network_connector--properties--sli_to_global_dr.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_global_dr for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md)
- [Property reference](data-sources--network_connector--reference.md)
- sli_to_global_dr

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: sli\_to\_global\_dr, sli\_to\_slo\_snat, slo\_to\_global\_dr\] Global network reference for
direct connection.

Upstream description:

Global network reference for direct connection.

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

OneOf alternatives in this subsection:

- [sli_to_global_dr](data-sources--network_connector--properties--sli_to_global_dr.md#section)
- [sli_to_slo_snat](data-sources--network_connector--properties--sli_to_slo_snat.md#section)
- [slo_to_global_dr](data-sources--network_connector--properties--slo_to_global_dr.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [global_vn](data-sources--network_connector--properties--sli_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [sli_to_global_dr.global_vn](data-sources--network_connector--properties--sli_to_global_dr--global_vn.md)
- [Property reference](data-sources--network_connector--reference.md)
- [xcsh_network_connector](../data-sources/network_connector.md)
