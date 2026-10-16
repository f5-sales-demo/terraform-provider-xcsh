---
page_title: "rules.client_city"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules client city"], "body_bytes": 975, "body_sha256": "sha256:f184ac1dfd5aad1028dab585c7873eba16eba7c08f8f599b408a5d34da2b4718", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "documentation/resources/user_identification/properties/rules/client_city/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3311130321010331-3323021212302320-2030123030310033-1121201303323103-0032230333132000-3221110011303120-0001220231033131-0110311031223310", "registry_path": "docs/guides/resources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "client_city"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/client_city/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["user_identificationCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.client_city

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/)
- rules.client_city

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
client_city = {}
```

This is an empty object or choice marker. It has no direct properties.
