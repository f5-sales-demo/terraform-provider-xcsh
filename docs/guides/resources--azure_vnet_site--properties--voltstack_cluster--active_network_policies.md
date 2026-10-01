---
page_title: "voltstack_cluster.active_network_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster.active_network_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1538, "body_sha256": "sha256:3cda3e3389cfdf1f7bc5eb50636e362bd5fa985c0d94fdfe81d067b43966437b", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster--active_network_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.active_network_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.active_network_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster](resources--azure_vnet_site--properties--voltstack_cluster.md)
- voltstack_cluster.active_network_policies

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

- [network_policies](resources--azure_vnet_site--properties--voltstack_cluster--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_network_policies.network_policies](resources--azure_vnet_site--properties--voltstack_cluster--active_network_policies--network_policies.md)
- [voltstack_cluster](resources--azure_vnet_site--properties--voltstack_cluster.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
