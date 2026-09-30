---
page_title: "ingress_gw_ar.accelerated_networking"
subcategory: "Infrastructure"
description: "ingress_gw_ar.accelerated_networking for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1578, "body_sha256": "sha256:19faf1b6f75f43c60e2716ff4078a6e13be5b1c1675274eccd7205381c5b23b6", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:disable_spec", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:enable"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw_ar", "accelerated_networking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw_ar.accelerated_networking for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_gw_ar.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_gw_ar](data-sources--azure_vnet_site--properties--ingress_gw_ar.md)
- ingress_gw_ar.accelerated_networking

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

- [disable_spec](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking--disable_spec.md): complete subsection reference.

- [enable](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking--enable.md): complete subsection reference.

## Next pages

- [ingress_gw_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking--disable_spec.md)
- [ingress_gw_ar.accelerated_networking.enable](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking--enable.md)
- [ingress_gw_ar](data-sources--azure_vnet_site--properties--ingress_gw_ar.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
