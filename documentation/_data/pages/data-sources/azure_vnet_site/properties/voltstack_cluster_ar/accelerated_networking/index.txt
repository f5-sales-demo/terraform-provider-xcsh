---
page_title: "voltstack_cluster_ar.accelerated_networking"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.accelerated_networking for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2115, "body_sha256": "sha256:1dc483c087ac9157cd6d68a9422d27e5be00c5a2903d18fb3257e2962c79349e", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:accelerated_networking:disable_spec", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:accelerated_networking:enable"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:accelerated_networking", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["voltstack_cluster_ar", "accelerated_networking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.accelerated_networking for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster_ar.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/)
- voltstack_cluster_ar.accelerated_networking

<a id="section"></a>

Type: `"single"`. Computed.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/enable/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/disable_spec/)
- [voltstack_cluster_ar.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/enable/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
