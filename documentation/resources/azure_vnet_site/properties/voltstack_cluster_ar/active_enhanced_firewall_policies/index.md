---
page_title: "voltstack_cluster_ar.active_enhanced_firewall_policies"
subcategory: "Infrastructure"
description: "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion."
xcsh_docs: {"aliases": ["voltstack cluster ar active enhanced firewall policies"], "body_bytes": 2308, "body_sha256": "sha256:dea67091a0fa67e3b3a9ab719af2e5024b74eb77863973d26c87d094a8f88efc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_enhanced_firewall_policies:enhanced_firewall_policies"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1301330232303323-2200110003013023-2121133213033011-0220001212130002-2332201223211310-1212130332121021-0223323003002111-1310131211131231", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.active_enhanced_firewall_policies:RequiredObjectAttributes:enhanced_firewall_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "active_enhanced_firewall_policies"], "schema_version": 1, "sections": [{"aliases": ["enhanced firewall policies"], "anchor": "section", "description": "Ordered List of Enhanced Firewall Policies active.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_enhanced_firewall_policies:enhanced_firewall_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--name", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "schema_path": ["voltstack_cluster_ar", "active_enhanced_firewall_policies", "enhanced_firewall_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- voltstack_cluster_ar.active_enhanced_firewall_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
