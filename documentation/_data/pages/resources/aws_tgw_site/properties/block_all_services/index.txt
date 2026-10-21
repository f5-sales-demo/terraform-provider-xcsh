---
page_title: "block_all_services"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["block all services"], "body_bytes": 1504, "body_sha256": "sha256:c944b653f003954f79bd2fb2b70dd61c6bac1a909e961340acf97e2e15853e76", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:block_all_services", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/block_all_services/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3303203233101111-1231211011103232-0133220301212023-0131322023033222-3032123223330012-1020121030201121-2102031230301211-3211000022222033", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["block_all_services"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/block_all_services/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# block_all_services

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- block_all_services

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/block_all_services/#section)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/#section)
- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/default_blocked_services/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

This is an empty object or choice marker. It has no direct properties.
