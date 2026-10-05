---
page_title: "where.virtual_network"
subcategory: ""
description: "This specifies a direct reference to a network configuration object."
xcsh_docs: {"aliases": ["where virtual network"], "body_bytes": 1730, "body_sha256": "sha256:b1078271a3320dadb234cb12e5ff9ca6c8b38c327b2a7caa9840c045caff71ec", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "parent_id": "xcsh-docs:resources:advertise_policy:properties:where", "path": "documentation/resources/advertise_policy/properties/where/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_network:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["where virtual network ref"], "anchor": "section", "description": "A virtual network direct reference.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["where", "virtual_network", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This specifies a direct reference to a network configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/)
- where.virtual_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/virtual_network/ref/): complete subsection reference.

## Next pages

- [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/virtual_network/ref/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
