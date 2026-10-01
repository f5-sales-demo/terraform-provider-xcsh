---
page_title: "access_info.tls_config.common_params.tls_certificates.private_key"
subcategory: ""
description: "access_info.tls_config.common_params.tls_certificates.private_key for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2743, "body_sha256": "sha256:923d570b6fb778c9f147d6a059ed4fab90f9bcc3fd55ced3d0cee3f0fa313982", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "path": "docs/guides/resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.common_params.tls_certificates.private_key for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](resources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.tls_config.common_params](resources--secret_management_access--properties--access_info--tls_config--common_params.md)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- access_info.tls_config.common_params.tls_certificates.private_key

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
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md)
- [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
