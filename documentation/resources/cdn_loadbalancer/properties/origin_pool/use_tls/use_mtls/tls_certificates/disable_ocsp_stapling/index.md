---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["origin pool use tls use mtls tls certificates disable ocsp stapling"], "body_bytes": 1607, "body_sha256": "sha256:b694998b15a23c3c45d9ac42e30342fff273ebbed9d3be09083f869a845c74e0", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0030000012212300-1203012233022132-3211102131102220-3312301020230120-1130301312112300-2123323020031311-1101031022003322-3131131211031002", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [origin_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/)
- [origin_pool.use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/)
- origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.
