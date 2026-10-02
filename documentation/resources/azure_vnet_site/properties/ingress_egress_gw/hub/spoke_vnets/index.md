---
page_title: "ingress_egress_gw.hub.spoke_vnets"
subcategory: "Infrastructure"
description: "Spoke VNet Peering."
xcsh_docs: {"aliases": ["ingress egress gw hub spoke vnets"], "body_bytes": 3592, "body_sha256": "sha256:5a1e321c48e69750f13e6eef18980305419feb2b01443a48644944eddd12b235", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:auto", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:labels", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3320303032312013-0231112302020313-2131133002322101-1102323321102310-3100131012320001-3133332123120201-2021301311023031-1010112000032303", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets:ConflictingListObjectAttributes:auto,manual", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets:ConflictingListObjectAttributes:auto,manual", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets"], "schema_version": 1, "sections": [{"aliases": ["auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:auto", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall policy These labels used must be from known key and label defined in shared namespace.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:labels", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "labels"], "syntax": "block", "type": "object"}, {"aliases": ["manual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "manual"], "syntax": "attribute", "type": "object"}, {"aliases": ["vnet"], "anchor": "section", "description": "Resource group and name of existing Azure VNet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:ConflictingObjectAttributes:f5_orchestrated_routing,manual_routing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:f5_orchestrated_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:ConflictingObjectAttributes:f5_orchestrated_routing,manual_routing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet:manual_routing", "type": "conflicts"}, {"anchor": "schema-ingress_egress_gw--hub--spoke_vnets--vnet--resource_group", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:RequiredObjectAttributes:resource_group,vnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "type": "requires"}, {"anchor": "schema-ingress_egress_gw--hub--spoke_vnets--vnet--vnet_name", "enforcement": "provider-schema", "group": "ingress_egress_gw.hub.spoke_vnets.vnet:RequiredObjectAttributes:resource_group,vnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "type": "requires"}], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Spoke VNet Peering.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.spoke_vnets

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- ingress_egress_gw.hub.spoke_vnets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("auto",
    "manual")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
spoke_vnets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/auto/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/labels/): complete subsection reference.

- [manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/manual/): complete subsection reference.

- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/): complete subsection reference.

## Next pages

- [ingress_egress_gw.hub.spoke_vnets.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/auto/)
- [ingress_egress_gw.hub.spoke_vnets.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/labels/)
- [ingress_egress_gw.hub.spoke_vnets.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/manual/)
- [ingress_egress_gw.hub.spoke_vnets.vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
