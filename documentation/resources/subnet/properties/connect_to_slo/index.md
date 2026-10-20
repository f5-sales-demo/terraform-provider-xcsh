---
page_title: "connect_to_slo"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["connect to slo"], "body_bytes": 843, "body_sha256": "sha256:6de0840e09a8619bbae58e1481bc81ac288d0c64fa04eb2384b970a04e0603b2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:connect_to_slo", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "documentation/resources/subnet/properties/connect_to_slo/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2221112002303323-1312003022203230-1110230122111201-1302202132320223-1132333021002002-1133133132221231-0220310313110131-0031223211313300", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connect_to_slo"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/connect_to_slo/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["subnetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connect_to_slo

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- connect_to_slo

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for connect to slo.

Additional upstream details:

This can be used for messages where no values are needed.

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
connect_to_slo = {}
```

This is an empty object or choice marker. It has no direct properties.
