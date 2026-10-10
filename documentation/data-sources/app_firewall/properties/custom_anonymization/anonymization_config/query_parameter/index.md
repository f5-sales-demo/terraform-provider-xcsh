---
page_title: "custom_anonymization.anonymization_config.query_parameter"
subcategory: "Security"
description: "Configure anonymization for HTTP Parameters."
xcsh_docs: {"aliases": ["custom anonymization anonymization config query parameter"], "body_bytes": 2225, "body_sha256": "sha256:e6a4996c89937ae2d2be78f91db39409777496fe6f8c5342d299dfab4aa51936", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3223321311230011-3302030021021133-1210012223323311-2112000322330230-1331223300031322-1120130002012321-3203200323123110-3222120002113302", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config query parameter query param name"], "anchor": "schema-custom_anonymization--anonymization_config--query_parameter--query_param_name", "description": "Masks the query parameter value. The setting does not mask the query parameter name. Wildcard matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk (*), or by using only an asterisk to match any query parameter name.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter", "query_param_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configure anonymization for HTTP Parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.query_parameter

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.query_parameter

<a id="section"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Parameters.

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

<a id="schema-custom_anonymization--anonymization_config--query_parameter--query_param_name"></a>

### query_param_name property

Type: `"string"`. Computed.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
