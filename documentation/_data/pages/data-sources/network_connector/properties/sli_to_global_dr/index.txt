---
page_title: "sli_to_global_dr"
subcategory: "Networking"
description: "sli_to_global_dr for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1961, "body_sha256": "sha256:fd03a6039cc1d21b84b43f8626de55bec40cb0c0d69a677fd68ebf51ac7134eb", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "documentation/data-sources/network_connector/properties/sli_to_global_dr/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_global_dr for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
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

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_global_dr/#section)
- [sli_to_slo_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_slo_snat/#section)
- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/slo_to_global_dr/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_global_dr/global_vn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
