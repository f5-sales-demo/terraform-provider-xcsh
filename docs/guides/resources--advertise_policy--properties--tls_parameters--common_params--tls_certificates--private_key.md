---
page_title: "tls_parameters.common_params.tls_certificates.private_key"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates.private_key for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2422, "body_sha256": "sha256:5877ef8f252ca3ed363f077f4f8ad862e70c6d9b1083e48f4f5e8753bca00a80", "canonical_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params:tls_certificates", "path": "docs/guides/resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates.private_key for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md)
- [Property reference](resources--advertise_policy--reference.md)
- [tls_parameters](resources--advertise_policy--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--advertise_policy--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md)
- tls_parameters.common_params.tls_certificates.private_key

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

- [blindfold_secret_info](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md)
- [xcsh_advertise_policy](../resources/advertise_policy.md)
