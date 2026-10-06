---
page_title: "tls_tcp_auto_cert"
subcategory: "Load Balancing"
description: "Choice for selecting TLS over TCP proxy with automatic certificates."
xcsh_docs: {"aliases": ["tls tcp auto cert"], "body_bytes": 1291, "body_sha256": "sha256:fe6a4ea41cd154461efb119922639b46acef386d2e16359fac080d9f22b11ddd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp_auto_cert"], "schema_version": 1, "sections": [{"aliases": ["tls tcp auto cert no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp_auto_cert", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp auto cert tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp_auto_cert", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp auto cert use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp_auto_cert", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Choice for selecting TLS over TCP proxy with automatic certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
