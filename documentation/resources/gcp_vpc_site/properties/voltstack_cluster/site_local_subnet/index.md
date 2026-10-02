---
page_title: "voltstack_cluster.site_local_subnet"
subcategory: "Infrastructure"
description: "This defines choice about GCP VPC network for a view."
xcsh_docs: {"aliases": ["voltstack cluster site local subnet"], "body_bytes": 2312, "body_sha256": "sha256:7a90eda8aed7afe4e0d1d083eb01bab6cd681fba8a565270c9ff31649f9d010a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:new_subnet"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:new_subnet", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "site_local_subnet"], "schema_version": 1, "sections": [{"aliases": ["existing subnet"], "anchor": "section", "description": "Name of existing GCP subnet.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--site_local_subnet--existing_subnet--subnet_name", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_subnet.existing_subnet:RequiredObjectAttributes:subnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "type": "requires"}], "schema_path": ["voltstack_cluster", "site_local_subnet", "existing_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["new subnet"], "anchor": "section", "description": "Parameters for GCP subnet.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:new_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--site_local_subnet--new_subnet--primary_ipv4", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_subnet.new_subnet:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:new_subnet", "type": "requires"}], "schema_path": ["voltstack_cluster", "site_local_subnet", "new_subnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines choice about GCP VPC network for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.site_local_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.site_local_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
site_local_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/existing_subnet/): complete subsection reference.

- [new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/new_subnet/): complete subsection reference.

## Next pages

- [voltstack_cluster.site_local_subnet.existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/existing_subnet/)
- [voltstack_cluster.site_local_subnet.new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/new_subnet/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
