---
page_title: "ingress_gw.accelerated_networking"
subcategory: "Infrastructure"
description: "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen."
xcsh_docs: {"aliases": ["ingress gw accelerated networking"], "body_bytes": 2403, "body_sha256": "sha256:41989b683c4a8585d1603e5ccbdeb63dd5517f64c9cabe0d611ba784eea909d4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:disable_spec", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:enable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw", "path": "documentation/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "accelerated_networking"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "accelerated_networking", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:accelerated_networking:enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "accelerated_networking", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/)
- ingress_gw.accelerated_networking

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.
Server applies default when omitted.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/enable/): complete subsection reference.

## Next pages

- [ingress_gw.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/disable_spec/)
- [ingress_gw.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/accelerated_networking/enable/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
