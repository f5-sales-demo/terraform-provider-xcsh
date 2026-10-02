---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["backend servers", "cert", "certificate", "existing certificates", "origin pool use tls use mtls tls certificates private key", "origin servers", "tls certificates", "upstream servers"], "body_bytes": 2826, "body_sha256": "sha256:1ba24257b5a6ddb6dbf746f00855adc8990cc5eb28af0247f4686fcf122d3717", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "path": "documentation/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1131310320310132-1322011321200001-2220003132113021-3212311312132031-2310031330033111-0333130012300332-3231311120112201-0122232033233013", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls.tls_certificates.private_key

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [origin_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/)
- [origin_pool.use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/)
- [origin_pool.use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
