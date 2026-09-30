---
page_title: "azure_pfx_certificate.password"
subcategory: "Infrastructure"
description: "azure_pfx_certificate.password for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1542, "body_sha256": "sha256:c8d09239cd2c391dbb2de73a8d9b332116cedf9cd59c719ce39f0a45d2e50047", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "path": "docs/guides/data-sources--cloud_credentials--properties--azure_pfx_certificate--password.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_pfx_certificate", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_pfx_certificate/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_pfx_certificate.password for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_pfx_certificate.password

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
- [azure_pfx_certificate](data-sources--cloud_credentials--properties--azure_pfx_certificate.md)
- azure_pfx_certificate.password

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

- [blindfold_secret_info](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [azure_pfx_certificate.password.blindfold_secret_info](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md)
- [azure_pfx_certificate.password.clear_secret_info](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md)
- [azure_pfx_certificate](data-sources--cloud_credentials--properties--azure_pfx_certificate.md)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
