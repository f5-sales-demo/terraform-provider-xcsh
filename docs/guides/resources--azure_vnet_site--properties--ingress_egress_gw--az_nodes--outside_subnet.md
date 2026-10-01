---
page_title: "ingress_egress_gw.az_nodes.outside_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw.az_nodes.outside_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1991, "body_sha256": "sha256:6f2938e3f44c6b56be13121cc3f2bc5f4ff6b5992819905b4d248bd7b96cda1a", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet_param"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.az_nodes.outside_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.outside_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.az_nodes](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md)
- ingress_egress_gw.az_nodes.outside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

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
outside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [subnet](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet.md): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.outside_subnet.subnet](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet.md)
- [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet_param.md)
- [ingress_egress_gw.az_nodes](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
