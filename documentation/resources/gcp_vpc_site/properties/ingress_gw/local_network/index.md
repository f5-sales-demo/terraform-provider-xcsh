---
page_title: "ingress_gw.local_network"
subcategory: "Infrastructure"
description: "This defines choice about GCP VPC network for a view."
xcsh_docs: {"aliases": ["ingress gw local network"], "body_bytes": 2816, "body_sha256": "sha256:792c12c958925c871f274131ee22b14bb7b12218f8f51e221d73736cf586f9ad", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:existing_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network_autogenerate"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw", "path": "documentation/resources/gcp_vpc_site/properties/ingress_gw/local_network/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0123313310300223-3102132301021233-3032132031120333-0202013203021322-3331302001033202-2300330222001312-3312031111121222-0020001221000222", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.local_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.local_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.local_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.local_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.local_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network_autogenerate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.local_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network_autogenerate", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "local_network"], "schema_version": 1, "sections": [{"aliases": ["existing network"], "anchor": "section", "description": "Name of existing VPC network.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:existing_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_gw--local_network--existing_network--name", "enforcement": "provider-schema", "group": "ingress_gw.local_network.existing_network:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:existing_network", "type": "requires"}], "schema_path": ["ingress_gw", "local_network", "existing_network"], "syntax": "block", "type": "object"}, {"aliases": ["new network"], "anchor": "section", "description": "Parameters to create a new GCP VPC Network.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_gw--local_network--new_network--name", "enforcement": "provider-schema", "group": "ingress_gw.local_network.new_network:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network", "type": "requires"}], "schema_path": ["ingress_gw", "local_network", "new_network"], "syntax": "block", "type": "object"}, {"aliases": ["new network autogenerate"], "anchor": "section", "description": "Create a new GCP VPC Network with autogenerated name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network_autogenerate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "local_network", "new_network_autogenerate"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_gw/local_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines choice about GCP VPC network for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.local_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/)
- ingress_gw.local_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
local_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/local_network/existing_network/): complete subsection reference.

- [new_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/local_network/new_network/): complete subsection reference.

- [new_network_autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/local_network/new_network_autogenerate/): complete subsection reference.

## Next pages

- [ingress_gw.local_network.existing_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/local_network/existing_network/)
- [ingress_gw.local_network.new_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/local_network/new_network/)
- [ingress_gw.local_network.new_network_autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/local_network/new_network_autogenerate/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
