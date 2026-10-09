---
page_title: "tls_tcp.tls_cert_params"
subcategory: "Load Balancing"
description: "Select TLS Parameters and Certificates."
xcsh_docs: {"aliases": ["tls tcp tls cert params"], "body_bytes": 1673, "body_sha256": "sha256:3109e9edfded321389b1e4e2f7ba476d801d5bde77698ca12d88c223e013e0b3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:certificates", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_cert_params"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls cert params certificates"], "anchor": "section", "description": "Select one or more certificates with any domain names.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tls_tcp", "tls_cert_params", "certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp tls cert params no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:no_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_cert_params", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp tls cert params tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:tls_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp", "tls_cert_params", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp tls cert params use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp", "tls_cert_params", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Select TLS Parameters and Certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_cert_params

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/)
- tls_tcp.tls_cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

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

- [certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/certificates/): complete subsection reference.

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/use_mtls/): complete subsection reference.
