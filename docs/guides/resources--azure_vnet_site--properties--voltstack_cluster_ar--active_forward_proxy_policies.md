---
page_title: "voltstack_cluster_ar.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.active_forward_proxy_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1471, "body_sha256": "sha256:c4da0ad99453afc9e0c7ff884f5677140f4631e9c816ea1d14cfa9381417236d", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_forward_proxy_policies", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.active_forward_proxy_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster_ar.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- voltstack_cluster_ar.active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](resources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](resources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
