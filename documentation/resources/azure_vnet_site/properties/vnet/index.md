---
page_title: "vnet"
subcategory: "Infrastructure"
description: "This defines choice about Azure VNet for a view."
xcsh_docs: {"aliases": ["vnet"], "body_bytes": 1898, "body_sha256": "sha256:54f9bf88bb5e91555599fb72278a044f0c0d9462d3617bfff22f1d1053e0d85a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:vnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/vnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2030003203102003-3211120202023310-3332230033311222-2030101330030212-1121003230120332-2223020303131011-1031203132311202-2111303213312103", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vnet:ConflictingObjectAttributes:existing_vnet,new_vnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vnet:ConflictingObjectAttributes:existing_vnet,new_vnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vnet"], "schema_version": 1, "sections": [{"aliases": ["existing vnet"], "anchor": "section", "description": "Resource group and name of existing Azure VNet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vnet.existing_vnet:ConflictingObjectAttributes:f5_orchestrated_routing,manual_routing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet:f5_orchestrated_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vnet.existing_vnet:ConflictingObjectAttributes:f5_orchestrated_routing,manual_routing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet:manual_routing", "type": "conflicts"}, {"anchor": "schema-vnet--existing_vnet--resource_group", "enforcement": "provider-schema", "group": "vnet.existing_vnet:RequiredObjectAttributes:resource_group,vnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "type": "requires"}, {"anchor": "schema-vnet--existing_vnet--vnet_name", "enforcement": "provider-schema", "group": "vnet.existing_vnet:RequiredObjectAttributes:resource_group,vnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "type": "requires"}], "schema_path": ["vnet", "existing_vnet"], "syntax": "block", "type": "object"}, {"aliases": ["new vnet"], "anchor": "section", "description": "Parameters to create a new Azure VNet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-vnet--new_vnet--name", "enforcement": "provider-schema", "group": "vnet.new_vnet:ConflictingObjectAttributes:autogenerate,name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vnet.new_vnet:ConflictingObjectAttributes:autogenerate,name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet:autogenerate", "type": "conflicts"}, {"anchor": "schema-vnet--new_vnet--primary_ipv4", "enforcement": "provider-schema", "group": "vnet.new_vnet:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet", "type": "requires"}], "schema_path": ["vnet", "new_vnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/vnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines choice about Azure VNet for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- vnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about Azure VNet for a view.

Upstream description:

This defines choice about Azure VNet for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_vnet",
    "new_vnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_vnet\",\"new_vnet\"]"
}
```

Terraform syntax:

```terraform
vnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/): complete subsection reference.

- [new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/new_vnet/): complete subsection reference.

## Next pages

- [vnet.existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/)
- [vnet.new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/new_vnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
