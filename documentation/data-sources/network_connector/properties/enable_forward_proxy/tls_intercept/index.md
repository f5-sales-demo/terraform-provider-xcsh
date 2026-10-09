---
page_title: "enable_forward_proxy.tls_intercept"
subcategory: "Networking"
description: "Configuration to enable TLS interception."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept"], "body_bytes": 3236, "body_sha256": "sha256:c14b29b048aa5613511016196b7872bf09c7f87875c18ccb9cf9bc444c283e9b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:enable_for_all_domains", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_certificate", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept", "parent_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "path": "documentation/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept"], "schema_version": 1, "sections": [{"aliases": ["enable forward proxy tls intercept custom certificate"], "anchor": "section", "description": "Handle to fetch certificate and key.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept enable for all domains"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:enable_for_all_domains", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "enable_for_all_domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept policy"], "anchor": "section", "description": "Policy to enable or disable TLS interception.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept trusted ca url"], "anchor": "schema-enable_forward_proxy--tls_intercept--trusted_ca_url", "description": "Exclusive with Custom Root CA Certificate for validating upstream server certificate.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept volterra certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "volterra_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable forward proxy tls intercept volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration to enable TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/)
- enable_forward_proxy.tls_intercept

<a id="section"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

## Direct properties

- [custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/): complete subsection reference.

- [enable_for_all_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/enable_for_all_domains/): complete subsection reference.

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/): complete subsection reference.

<a id="schema-enable_forward_proxy--tls_intercept--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_certificate/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_trusted_ca/): complete subsection reference.
