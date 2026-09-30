---
page_title: "voltstack_cluster_ar.node.local_subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.node.local_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2364, "body_sha256": "sha256:a1cc3d4e8d97679114e9ca4d9ca535f2fa1d6982d80641cf62bb764cc4476d60", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet", "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet_param"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["voltstack_cluster_ar", "node", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.node.local_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster_ar.node.local_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [voltstack_cluster_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/)
- voltstack_cluster_ar.node.local_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
```

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

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet_param/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.node.local_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/)
- [voltstack_cluster_ar.node.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet_param/)
- [voltstack_cluster_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
