---
page_title: "ingress_egress_gw.active_network_policies"
subcategory: "Infrastructure"
description: "ingress_egress_gw.active_network_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1538, "body_sha256": "sha256:181b7d0bc4106439d42f40b0764debb1bc28cf31f1f6f410879345b5e05f8bba", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.active_network_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.active_network_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
```

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

Terraform syntax:

```terraform
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_network_policies.network_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies--network_policies.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
