---
page_title: "aws_parameters.admin_password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["aws parameters admin password"], "body_bytes": 2287, "body_sha256": "sha256:705547a0341e0f2e6da430561316165ea58cf043b7bd22a2f33dde93f3d6a3ff", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:blindfold_secret_info", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/admin_password/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2020332001100331-1200322312312112-1111030131330011-1232231201300330-2111132210322012-3131233231123113-0310232032230311-0112212210101102", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "admin_password"], "schema_version": 1, "sections": [{"aliases": ["aws parameters admin password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--admin_password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "aws_parameters.admin_password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:blindfold_secret_info", "type": "requires"}], "schema_path": ["aws_parameters", "admin_password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters admin password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--admin_password--clear_secret_info--url", "enforcement": "provider-schema", "group": "aws_parameters.admin_password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:clear_secret_info", "type": "requires"}], "schema_path": ["aws_parameters", "admin_password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/admin_password/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.admin_password

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.admin_password

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/clear_secret_info/): complete subsection reference.

## Next pages

- [aws_parameters.admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/blindfold_secret_info/)
- [aws_parameters.admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/clear_secret_info/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
