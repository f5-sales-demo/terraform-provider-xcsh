---
page_title: "ingress_egress_gw.inside_subnet"
subcategory: "Infrastructure"
description: "This defines choice about GCP VPC network for a view."
xcsh_docs: {"aliases": ["ingress egress gw inside subnet"], "body_bytes": 2276, "body_sha256": "sha256:e96f135ba47acefc9841a77e88c1b893a737b9637e0a7be717012a43f5c1235d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:new_subnet"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:new_subnet", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_subnet"], "schema_version": 1, "sections": [{"aliases": ["existing subnet"], "anchor": "section", "description": "Name of existing GCP subnet.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--inside_subnet--existing_subnet--subnet_name", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet.existing_subnet:RequiredObjectAttributes:subnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "type": "requires"}], "schema_path": ["ingress_egress_gw", "inside_subnet", "existing_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["new subnet"], "anchor": "section", "description": "Parameters for GCP subnet.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:new_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--inside_subnet--new_subnet--primary_ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet.new_subnet:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:new_subnet", "type": "requires"}], "schema_path": ["ingress_egress_gw", "inside_subnet", "new_subnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines choice about GCP VPC network for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- ingress_egress_gw.inside_subnet

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
inside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/existing_subnet/): complete subsection reference.

- [new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/new_subnet/): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_subnet.existing_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/existing_subnet/)
- [ingress_egress_gw.inside_subnet.new_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/new_subnet/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
