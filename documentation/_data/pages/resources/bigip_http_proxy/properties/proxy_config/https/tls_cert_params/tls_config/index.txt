---
page_title: "proxy_config.https.tls_cert_params.tls_config"
subcategory: ""
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["proxy config https tls cert params tls config"], "body_bytes": 2267, "body_sha256": "sha256:93b8bcdbdd92fc80d0a8110526a2c17235bb1ac8622b9a9080ac3c0330beffb1", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:default_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:low_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["proxy config https tls cert params tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:custom_security", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config", "custom_security"], "syntax": "block", "type": "object"}, {"aliases": ["proxy config https tls cert params tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:default_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https tls cert params tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:low_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https tls cert params tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:medium_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/)
- proxy_config.https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/medium_security/): complete subsection reference.
