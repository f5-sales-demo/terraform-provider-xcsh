---
page_title: "tls_intercept.custom_certificate.private_key"
subcategory: ""
description: "tls_intercept.custom_certificate.private_key for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1641, "body_sha256": "sha256:6c29bb0825c55392b3c70eff4aab1ad30995aa40e115278b0c6eabdc522992cd", "canonical_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key:blindfold_secret_info", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "path": "docs/guides/data-sources--proxy--properties--tls_intercept--custom_certificate--private_key.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "custom_certificate", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/custom_certificate/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.custom_certificate.private_key for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_intercept.custom_certificate.private_key

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md)
- [tls_intercept.custom_certificate](data-sources--proxy--properties--tls_intercept--custom_certificate.md)
- tls_intercept.custom_certificate.private_key

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

- [blindfold_secret_info](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md)
- [tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md)
- [tls_intercept.custom_certificate](data-sources--proxy--properties--tls_intercept--custom_certificate.md)
- [xcsh_proxy](../data-sources/proxy.md)
