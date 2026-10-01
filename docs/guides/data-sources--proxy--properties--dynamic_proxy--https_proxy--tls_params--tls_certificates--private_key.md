---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2215, "body_sha256": "sha256:800b33bf0b28be167b432275c7b83a12338d9053a73d04b692d254465e1277a5", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- [xcsh_proxy](../data-sources/proxy.md)
