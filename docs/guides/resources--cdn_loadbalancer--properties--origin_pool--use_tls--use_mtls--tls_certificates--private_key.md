---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.use_mtls.tls_certificates.private_key for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2414, "body_sha256": "sha256:2806d85debfbe1239c54d17b6793477ea0e2900cf4f2cbd9af909349c189f953", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.use_mtls.tls_certificates.private_key for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pool.use_tls.use_mtls.tls_certificates.private_key

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

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

- [blindfold_secret_info](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
