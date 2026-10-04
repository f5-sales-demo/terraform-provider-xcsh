---
page_title: "tls_tcp.tls_parameters"
subcategory: "Load Balancing"
description: "Inline TLS parameters."
xcsh_docs: {"aliases": ["tls tcp tls parameters"], "body_bytes": 2585, "body_sha256": "sha256:9863adae8a89cd1f0f920af8b6c9d001fe0150e077a41d368c9dd0d419025b76", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:use_mtls"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["tls tcp tls parameters no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls parameters tls certificates"], "anchor": "section", "description": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp tls parameters tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp tls parameters use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Inline TLS parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/)
- tls_tcp.tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

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

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/no_mtls/): complete subsection reference.

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/use_mtls/): complete subsection reference.

## Next pages

- [tls_tcp.tls_parameters.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/no_mtls/)
- [tls_tcp.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/)
- [tls_tcp.tls_parameters.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_config/)
- [tls_tcp.tls_parameters.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/use_mtls/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
