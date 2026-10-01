---
page_title: "gcp_cred_file.credential_file"
subcategory: "Infrastructure"
description: "gcp_cred_file.credential_file for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1886, "body_sha256": "sha256:ef9e60f06f254278cbd6794412421f1432d44f123c748cb9b5bc3c4ea1df901c", "canonical_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info"], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "path": "docs/guides/resources--cloud_credentials--properties--gcp_cred_file--credential_file.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp_cred_file", "credential_file"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_cred_file.credential_file for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file.credential_file

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
- [Property reference](resources--cloud_credentials--reference.md)
- [gcp_cred_file](resources--cloud_credentials--properties--gcp_cred_file.md)
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

- [blindfold_secret_info](resources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md): complete subsection reference.

## Next pages

- [gcp_cred_file.credential_file.blindfold_secret_info](resources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md)
- [gcp_cred_file.credential_file.clear_secret_info](resources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md)
- [gcp_cred_file](resources--cloud_credentials--properties--gcp_cred_file.md)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
