---
page_title: "voltstack_cluster_ar.active_network_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.active_network_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1466, "body_sha256": "sha256:b0bc848656b1ff55c5f212bca65f10298365dda179914b1e95b9c3d76d70aa5d", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.active_network_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster_ar.active_network_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- voltstack_cluster_ar.active_network_policies

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

- [network_policies](resources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.active_network_policies.network_policies](resources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
