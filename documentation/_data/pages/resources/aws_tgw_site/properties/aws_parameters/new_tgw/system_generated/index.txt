---
page_title: "aws_parameters.new_tgw.system_generated"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws parameters new tgw system generated"], "body_bytes": 1183, "body_sha256": "sha256:781a1dd9d897f02c0181ab9a5ffa5c88b284ada067eea21a67c65593d7dcbd62", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/new_tgw/system_generated/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0103221113020220-1023032300110200-1122320003021100-3301223110103000-1232122010031200-2123003310013320-0101220201312303-1332330322113022", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "new_tgw", "system_generated"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_tgw/system_generated/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_tgw.system_generated

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [aws_parameters.new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/)
- aws_parameters.new_tgw.system_generated

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system generated.

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
system_generated = {}
```

This is an empty object or choice marker. It has no direct properties.
