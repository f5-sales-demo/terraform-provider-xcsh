---
page_title: "aws_secret_key.secret_key"
subcategory: "Infrastructure"
description: "aws_secret_key.secret_key for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1573, "body_sha256": "sha256:61502c67554140e1310f4c8b4f9184bef2d793a09e3df77f0ba719a6d9780fc2", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key", "path": "docs/guides/data-sources--cloud_credentials--properties--aws_secret_key--secret_key.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_secret_key", "secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_secret_key.secret_key for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_secret_key.secret_key

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
- [aws_secret_key](data-sources--cloud_credentials--properties--aws_secret_key.md)
- aws_secret_key.secret_key

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

- [blindfold_secret_info](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--blindfold_secret_info.md)
- [aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--clear_secret_info.md)
- [aws_secret_key](data-sources--cloud_credentials--properties--aws_secret_key.md)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
