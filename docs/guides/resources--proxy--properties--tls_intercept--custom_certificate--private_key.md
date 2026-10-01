---
page_title: "tls_intercept.custom_certificate.private_key"
subcategory: ""
description: "tls_intercept.custom_certificate.private_key for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2018, "body_sha256": "sha256:107b2dae6acf4a4e6d5f5279e3fb5c19ce3bb9163ad977093f6e10a20f1520d5", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key:blindfold_secret_info", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "path": "docs/guides/resources--proxy--properties--tls_intercept--custom_certificate--private_key.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "custom_certificate", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/custom_certificate/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.custom_certificate.private_key for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.custom_certificate.private_key

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- [tls_intercept.custom_certificate](resources--proxy--properties--tls_intercept--custom_certificate.md)
- tls_intercept.custom_certificate.private_key

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

- [blindfold_secret_info](resources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md)
- [tls_intercept.custom_certificate.private_key.clear_secret_info](resources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md)
- [tls_intercept.custom_certificate](resources--proxy--properties--tls_intercept--custom_certificate.md)
- [xcsh_proxy](../resources/proxy.md)
