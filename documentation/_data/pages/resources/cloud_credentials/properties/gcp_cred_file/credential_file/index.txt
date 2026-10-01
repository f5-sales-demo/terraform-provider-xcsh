---
page_title: "gcp_cred_file.credential_file"
subcategory: "Infrastructure"
description: "gcp_cred_file.credential_file for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 2339, "body_sha256": "sha256:13efdf2f002c0128472a3583f599f543f9f024e67ab887c1f1ec7d343b848bb0", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info"], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "path": "documentation/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["gcp_cred_file", "credential_file"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_cred_file.credential_file for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file.credential_file

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/)
- gcp_cred_file.credential_file

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
credential_file {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/): complete subsection reference.

## Next pages

- [gcp_cred_file.credential_file.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/)
- [gcp_cred_file.credential_file.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
