---
page_title: "admin_password"
subcategory: "Infrastructure"
description: "admin_password for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1622, "body_sha256": "sha256:8a87beb92119f6d2d0589feff483fffc9391b44c318483f0ebac16b49496c6da", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:admin_password", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:admin_password:blindfold_secret_info", "xcsh-docs:resources:aws_vpc_site:properties:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:admin_password", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--admin_password.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "admin_password for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_password

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- admin_password

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

- [blindfold_secret_info](resources--aws_vpc_site--properties--admin_password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--aws_vpc_site--properties--admin_password--clear_secret_info.md): complete subsection reference.

## Next pages

- [admin_password.blindfold_secret_info](resources--aws_vpc_site--properties--admin_password--blindfold_secret_info.md)
- [admin_password.clear_secret_info](resources--aws_vpc_site--properties--admin_password--clear_secret_info.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
