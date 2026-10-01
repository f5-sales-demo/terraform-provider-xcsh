---
page_title: "ingress_egress_gw.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "ingress_egress_gw.active_forward_proxy_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1265, "body_sha256": "sha256:44048792037ff102636c46064e67989b5c426703b5cbcbdacec2c1a5005dd642", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.active_forward_proxy_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.active_forward_proxy_policies

<a id="section"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

## Direct properties

- [forward_proxy_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
