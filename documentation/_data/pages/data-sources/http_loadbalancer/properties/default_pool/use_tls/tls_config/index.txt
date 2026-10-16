---
page_title: "default_pool.use_tls.tls_config"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["default pool use tls tls config"], "body_bytes": 2276, "body_sha256": "sha256:82f9ebe17c1a94887811e0a42540d322882ee15c525879922e2aa650f6c0755b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "use_tls", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["default pool use tls tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "custom_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool use tls tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool use tls tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool use tls tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/)
- default_pool.use_tls.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/medium_security/): complete subsection reference.
