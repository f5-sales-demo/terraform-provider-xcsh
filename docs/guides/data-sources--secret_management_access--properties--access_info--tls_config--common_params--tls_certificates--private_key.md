---
page_title: "access_info.tls_config.common_params.tls_certificates.private_key"
subcategory: ""
description: "access_info.tls_config.common_params.tls_certificates.private_key for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2372, "body_sha256": "sha256:9946474c73130c019c3f4353a42803f5199cf72b9a0c7ea1063b57eeefca9239", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "path": "docs/guides/data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.common_params.tls_certificates.private_key for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# access_info.tls_config.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](data-sources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.tls_config.common_params](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- access_info.tls_config.common_params.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md)
- [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
