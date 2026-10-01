---
page_title: "sli_to_global_dr"
subcategory: "Networking"
description: "sli_to_global_dr for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1598, "body_sha256": "sha256:bcaf57928ea7fd1088247b541b93c01e5af9c3cf12b7d25327124a676e10a139", "canonical_id": "xcsh-docs:resources:network_connector:properties:sli_to_global_dr", "child_ids": ["xcsh-docs:resources:network_connector:properties:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_global_dr", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "docs/guides/resources--network_connector--properties--sli_to_global_dr.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_global_dr for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- sli_to_global_dr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [sli_to_global_dr](resources--network_connector--properties--sli_to_global_dr.md#section)
- [sli_to_slo_snat](resources--network_connector--properties--sli_to_slo_snat.md#section)
- [slo_to_global_dr](resources--network_connector--properties--slo_to_global_dr.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](resources--network_connector--properties--sli_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [sli_to_global_dr.global_vn](resources--network_connector--properties--sli_to_global_dr--global_vn.md)
- [Property reference](resources--network_connector--reference.md)
- [xcsh_network_connector](../resources/network_connector.md)
