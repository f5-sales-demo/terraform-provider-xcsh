---
page_title: "aws_parameters.admin_password"
subcategory: ""
description: "aws_parameters.admin_password for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1735, "body_sha256": "sha256:def18afc6acfa64e9f21fc05f8703f51e47631b10b6629d0a48631726e3eadc3", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:blindfold_secret_info", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "docs/guides/resources--aws_tgw_site--properties--aws_parameters--admin_password.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_parameters", "admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.admin_password for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws_parameters.admin_password

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
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

- [blindfold_secret_info](resources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md): complete subsection reference.

## Next pages

- [aws_parameters.admin_password.blindfold_secret_info](resources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md)
- [aws_parameters.admin_password.clear_secret_info](resources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md)
- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
