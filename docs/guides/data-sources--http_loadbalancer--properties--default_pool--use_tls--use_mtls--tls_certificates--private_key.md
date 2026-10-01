---
page_title: "default_pool.use_tls.use_mtls.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "default_pool.use_tls.use_mtls.tls_certificates.private_key for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2273, "body_sha256": "sha256:3a7d24315c2b4ad911844edc8c902defe1cffb6f23ac865f014de12e05cd51d5", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates--private_key.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "use_tls", "use_mtls", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.use_tls.use_mtls.tls_certificates.private_key for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.use_mtls.tls_certificates.private_key

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.use_tls](data-sources--http_loadbalancer--properties--default_pool--use_tls.md)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls.md)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates.md)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
