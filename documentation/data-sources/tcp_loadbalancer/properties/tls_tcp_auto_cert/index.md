---
page_title: "tls_tcp_auto_cert"
subcategory: "Load Balancing"
description: "Choice for selecting TLS over TCP proxy with automatic certificates."
xcsh_docs: {"aliases": ["tls tcp auto cert"], "body_bytes": 2016, "body_sha256": "sha256:ff6b0ed63a0766e211ad3208885a147a20da0220933bfdaed39fdd52d8c481a1", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp_auto_cert"], "schema_version": 1, "sections": [{"aliases": ["tls tcp auto cert no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp_auto_cert", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp auto cert tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp_auto_cert", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp auto cert use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp_auto_cert", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Choice for selecting TLS over TCP proxy with automatic certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- tls_tcp_auto_cert

<a id="section"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with automatic certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

## Direct properties

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/): complete subsection reference.

## Next pages

- [tls_tcp_auto_cert.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/no_mtls/)
- [tls_tcp_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/)
- [tls_tcp_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
