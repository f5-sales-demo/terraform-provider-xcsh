---
page_title: "origin_pool.use_tls.use_mtls"
subcategory: "Load Balancing"
description: "MTLS Client Certificate."
xcsh_docs: {"aliases": ["backend servers", "cert", "certificate", "existing certificates", "origin pool use tls use mtls", "origin servers", "tls certificates", "upstream servers"], "body_bytes": 1690, "body_sha256": "sha256:5ea9463bf71091f52660157bc114a318d274f6c1721e09f458a839a645e8ea0d", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls", "path": "documentation/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates"], "anchor": "section", "description": "MTLS Client Certificate.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "MTLS Client Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- origin_pool.use_tls.use_mtls

<a id="section"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

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

## Direct properties

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/): complete subsection reference.

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
