---
page_title: "tgw_security.no_network_policy"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tgw security no network policy"], "body_bytes": 1013, "body_sha256": "sha256:f4b87ee28468a3b00afe3d9af198ecc16f8d339161713adac87101e5b8304274", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_network_policy", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "path": "documentation/resources/aws_tgw_site/properties/tgw_security/no_network_policy/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1203031221233023-1200122111122211-2020320201211200-1312131320023211-3002302120132300-3203030331100323-3021220313112001-0230023220113030", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tgw_security", "no_network_policy"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/no_network_policy/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.no_network_policy

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/)
- tgw_security.no_network_policy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

This is an empty object or choice marker. It has no direct properties.
