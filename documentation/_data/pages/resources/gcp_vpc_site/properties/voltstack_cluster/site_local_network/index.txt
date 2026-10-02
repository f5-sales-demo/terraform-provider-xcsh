---
page_title: "voltstack_cluster.site_local_network"
subcategory: "Infrastructure"
description: "This defines choice about GCP VPC network for a view."
xcsh_docs: {"aliases": ["voltstack cluster site local network"], "body_bytes": 2981, "body_sha256": "sha256:52074bbb69722d131a46066b73c7634989fd9a40f1d211a8a1391c6a1661bfda", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "site_local_network"], "schema_version": 1, "sections": [{"aliases": ["existing network"], "anchor": "section", "description": "Name of existing VPC network.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--site_local_network--existing_network--name", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network.existing_network:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "type": "requires"}], "schema_path": ["voltstack_cluster", "site_local_network", "existing_network"], "syntax": "block", "type": "object"}, {"aliases": ["new network"], "anchor": "section", "description": "Parameters to create a new GCP VPC Network.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--site_local_network--new_network--name", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network.new_network:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "type": "requires"}], "schema_path": ["voltstack_cluster", "site_local_network", "new_network"], "syntax": "block", "type": "object"}, {"aliases": ["new network autogenerate"], "anchor": "section", "description": "Create a new GCP VPC Network with autogenerated name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "site_local_network", "new_network_autogenerate"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines choice about GCP VPC network for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.site_local_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.site_local_network

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
site_local_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/existing_network/): complete subsection reference.

- [new_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network/): complete subsection reference.

- [new_network_autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network_autogenerate/): complete subsection reference.

## Next pages

- [voltstack_cluster.site_local_network.existing_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/existing_network/)
- [voltstack_cluster.site_local_network.new_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network/)
- [voltstack_cluster.site_local_network.new_network_autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network_autogenerate/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
