---
page_title: "custom_anonymization.anonymization_config"
subcategory: "Security"
description: "List of HTTP headers, cookies and query parameters whose values will be masked."
xcsh_docs: {"aliases": ["custom anonymization anonymization config"], "body_bytes": 2167, "body_sha256": "sha256:82d7aaadb7500f14d62d13dac7c91d98c37176f79ee9947f2ed50103e7fa1855", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "path": "documentation/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3313103023110012-3112123312210312-2333322032100303-0002031333003321-0330010220033330-2232312011032230-0002312223300213-3010110200220233", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config cookie"], "anchor": "section", "description": "Configure anonymization for HTTP Cookies.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "cookie"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom anonymization anonymization config http header"], "anchor": "section", "description": "Configure anonymization for HTTP Headers.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "http_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom anonymization anonymization config query parameter"], "anchor": "section", "description": "Configure anonymization for HTTP Parameters.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of HTTP headers, cookies and query parameters whose values will be masked.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/)
- custom_anonymization.anonymization_config

<a id="section"></a>

Type: `"list"`. Computed.

List of HTTP headers, cookies and query parameters whose values will be masked.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/): complete subsection reference.

- [http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/): complete subsection reference.

- [query_parameter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/): complete subsection reference.
