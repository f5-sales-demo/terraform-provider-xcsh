---
page_title: "where.virtual_network"
subcategory: ""
description: "This specifies a direct reference to a network configuration object."
xcsh_docs: {"aliases": ["where virtual network"], "body_bytes": 1109, "body_sha256": "sha256:1100629c2bf6329e845d9b1214a75a2be3cc39144b5a409eb5d735c25019ff86", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "parent_id": "xcsh-docs:resources:advertise_policy:properties:where", "path": "documentation/resources/advertise_policy/properties/where/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["where virtual network ref"], "anchor": "section", "description": "A virtual network direct reference.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["where", "virtual_network", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This specifies a direct reference to a network configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

This specifies a direct reference to a network configuration object.

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
