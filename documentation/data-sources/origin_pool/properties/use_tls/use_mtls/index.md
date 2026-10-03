---
page_title: "use_tls.use_mtls"
subcategory: "Load Balancing"
description: "MTLS Client Certificate."
xcsh_docs: {"aliases": ["use tls use mtls"], "body_bytes": 1410, "body_sha256": "sha256:df6c7cc3869b8e9ce038037fb0bf422d5d86fabb2d8b614c6894d87e92c2f237", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "path": "documentation/data-sources/origin_pool/properties/use_tls/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_tls", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use tls use mtls tls certificates"], "anchor": "section", "description": "MTLS Client Certificate.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/use_mtls/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "MTLS Client Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/)
- use_tls.use_mtls

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

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
