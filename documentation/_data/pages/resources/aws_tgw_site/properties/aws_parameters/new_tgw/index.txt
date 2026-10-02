---
page_title: "aws_parameters.new_tgw"
subcategory: ""
description: "TGWParamsType."
xcsh_docs: {"aliases": ["aws parameters new tgw"], "body_bytes": 2102, "body_sha256": "sha256:ac6740ded06c76bae223255739c544af36082cfa9face2ef0a4389569b2415ff", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/new_tgw/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.new_tgw:ConflictingObjectAttributes:system_generated,user_assigned", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.new_tgw:ConflictingObjectAttributes:system_generated,user_assigned", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "new_tgw"], "schema_version": 1, "sections": [{"aliases": ["system generated"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "system_generated"], "syntax": "attribute", "type": "object"}, {"aliases": ["user assigned"], "anchor": "section", "description": "Information needed when ASNs are assigned by the user.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_tgw/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TGWParamsType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_tgw

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.new_tgw

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TGWParamsType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("system_generated",
    "user_assigned")}
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
  "x-ves-oneof-field-asn_choice": "[\"system_generated\",\"user_assigned\"]"
}
```

Terraform syntax:

```terraform
new_tgw {
  # Configure direct properties listed below.
}
```

## Direct properties

- [system_generated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/system_generated/): complete subsection reference.

- [user_assigned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/): complete subsection reference.

## Next pages

- [aws_parameters.new_tgw.system_generated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/system_generated/)
- [aws_parameters.new_tgw.user_assigned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
