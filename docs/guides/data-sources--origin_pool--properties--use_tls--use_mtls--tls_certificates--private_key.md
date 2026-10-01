---
page_title: "use_tls.use_mtls.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "use_tls.use_mtls.tls_certificates.private_key for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1900, "body_sha256": "sha256:15c949db95278e12c4e01f5cfcaa68306dd552e849d5ba4e3817cc43fd44642a", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "path": "docs/guides/data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls.tls_certificates.private_key for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls.tls_certificates.private_key

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- [use_tls.use_mtls](data-sources--origin_pool--properties--use_tls--use_mtls.md)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- use_tls.use_mtls.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md)
- [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
