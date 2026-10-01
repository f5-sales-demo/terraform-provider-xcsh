---
page_title: "azure_pfx_certificate.password"
subcategory: "Infrastructure"
description: "azure_pfx_certificate.password for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1641, "body_sha256": "sha256:f0a7be5a2b407ae08d06814e4403ade122611f060a5d4a8c6d3bbef499b32234", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "path": "docs/guides/data-sources--cloud_credentials--properties--azure_pfx_certificate--password.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_pfx_certificate", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_pfx_certificate/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_pfx_certificate.password for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
