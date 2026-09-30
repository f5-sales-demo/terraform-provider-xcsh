---
page_title: "ingress_gw_ar.node.local_subnet"
subcategory: "Infrastructure"
description: "ingress_gw_ar.node.local_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1509, "body_sha256": "sha256:96fb5f10ac1c4073e5e3563cefae8d176bfc4fe5c596b5aadcd8299c3ece1429", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet:subnet_param"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw_ar", "node", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw_ar.node.local_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_gw_ar.node.local_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_gw_ar](data-sources--azure_vnet_site--properties--ingress_gw_ar.md)
- [ingress_gw_ar.node](data-sources--azure_vnet_site--properties--ingress_gw_ar--node.md)
- ingress_gw_ar.node.local_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

## Direct properties

- [subnet](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet.md): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [ingress_gw_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet.md)
- [ingress_gw_ar.node.local_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet_param.md)
- [ingress_gw_ar.node](data-sources--azure_vnet_site--properties--ingress_gw_ar--node.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
