---
page_title: "sms"
subcategory: ""
description: "SMS Configuration."
xcsh_docs: {"aliases": ["sms"], "body_bytes": 1424, "body_sha256": "sha256:1984fd60184d9282b76f10e32c56b60204237e8e98488054137d6e33bcaf340a", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:sms", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "documentation/data-sources/alert_receiver/properties/sms/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1102020323302231-3330031312223101-2332200311130100-0221003321130033-0310033330311110-2222220310130220-0311300330010121-3122100310111130", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sms"], "schema_version": 1, "sections": [{"aliases": ["sms contact number"], "anchor": "schema-sms--contact_number", "description": "Contact number of the user in ITU E.164 format", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:sms", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sms", "contact_number"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/sms/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "SMS Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sms

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- sms

<a id="section"></a>

Type: `"single"`. Computed.

SMS Configuration.

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

<a id="schema-sms--contact_number"></a>

### contact_number property

Type: `"string"`. Computed.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```
