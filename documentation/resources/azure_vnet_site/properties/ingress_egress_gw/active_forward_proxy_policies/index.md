---
page_title: "ingress_egress_gw.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["ingress egress gw active forward proxy policies"], "body_bytes": 1898, "body_sha256": "sha256:a7e490e7190e17c717b75a33eb09d7ba746dde980c2f1cab13441f765b0ff74b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0113300232113103-1212023322102110-0220312231322132-1301033012010110-3132033003332200-2330233211311212-0033131002331110-1121012132232110", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.active_forward_proxy_policies:RequiredObjectAttributes:forward_proxy_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--name", "enforcement": "provider-schema", "group": "ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "schema_path": ["ingress_egress_gw", "active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
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

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
