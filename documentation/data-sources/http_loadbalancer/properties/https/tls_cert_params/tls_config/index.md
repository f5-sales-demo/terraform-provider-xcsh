---
page_title: "https.tls_cert_params.tls_config"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["https tls cert params tls config"], "body_bytes": 1925, "body_sha256": "sha256:2720f664477a6b137d6758510aa58e92420ab1118dd7a24845c77d3abea593f4", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:default_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:low_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params", "path": "documentation/data-sources/http_loadbalancer/properties/https/tls_cert_params/tls_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_cert_params", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["https tls cert params tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:custom_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https", "tls_cert_params", "tls_config", "custom_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https tls cert params tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:default_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_cert_params", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https tls cert params tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:low_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_cert_params", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https tls cert params tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:tls_config:medium_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_cert_params", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/)
- [https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_cert_params/)
- https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_cert_params/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_cert_params/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_cert_params/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_cert_params/tls_config/medium_security/): complete subsection reference.
