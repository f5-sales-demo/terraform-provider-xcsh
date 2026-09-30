---
page_title: "ingress_egress_gw.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "ingress_egress_gw.active_forward_proxy_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1444, "body_sha256": "sha256:0ac74920b1ba1feb7770b9916e633db5f695237612714e01629229c384c943bb", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.active_forward_proxy_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.active_forward_proxy_policies

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

- [forward_proxy_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
