---
page_title: "aws_parameters.new_tgw"
subcategory: ""
description: "aws_parameters.new_tgw for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2102, "body_sha256": "sha256:ac6740ded06c76bae223255739c544af36082cfa9face2ef0a4389569b2415ff", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/new_tgw/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["aws_parameters", "new_tgw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_tgw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.new_tgw for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
