---
page_title: "gcp_cred_file.credential_file"
subcategory: "Infrastructure"
description: "gcp_cred_file.credential_file for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1601, "body_sha256": "sha256:9e86111f5dfdf288561d0258b5db8aa41a915e80d04b14270412b6c3d19b1af2", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file", "path": "docs/guides/data-sources--cloud_credentials--properties--gcp_cred_file--credential_file.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp_cred_file", "credential_file"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/gcp_cred_file/credential_file/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_cred_file.credential_file for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file.credential_file

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
- [gcp_cred_file](data-sources--cloud_credentials--properties--gcp_cred_file.md)
- gcp_cred_file.credential_file

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md): complete subsection reference.

## Next pages

- [gcp_cred_file.credential_file.blindfold_secret_info](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md)
- [gcp_cred_file.credential_file.clear_secret_info](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md)
- [gcp_cred_file](data-sources--cloud_credentials--properties--gcp_cred_file.md)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
