---
page_title: "admin_user_credentials.admin_password"
subcategory: ""
description: "admin_user_credentials.admin_password for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2449, "body_sha256": "sha256:8e739354371e7c03955f0286b6f901e3f8af5ecdb49bb96147108e8aef8d73d8", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:blindfold_secret_info", "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials", "path": "documentation/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["admin_user_credentials", "admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "admin_user_credentials.admin_password for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_user_credentials.admin_password

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [admin_user_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/)
- admin_user_credentials.admin_password

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/clear_secret_info/): complete subsection reference.

## Next pages

- [admin_user_credentials.admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/blindfold_secret_info/)
- [admin_user_credentials.admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/clear_secret_info/)
- [admin_user_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
