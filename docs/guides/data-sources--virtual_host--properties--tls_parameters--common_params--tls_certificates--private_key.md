---
page_title: "tls_parameters.common_params.tls_certificates.private_key"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates.private_key for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1996, "body_sha256": "sha256:79afc12f23809e1c22d25b35c3d615d515587b2f2af9074141818fa0330685db", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "path": "docs/guides/data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates.private_key for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [tls_parameters](data-sources--virtual_host--properties--tls_parameters.md)
- [tls_parameters.common_params](data-sources--virtual_host--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates.md)
- tls_parameters.common_params.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--properties--tls_parameters--common_params--tls_certificates.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
