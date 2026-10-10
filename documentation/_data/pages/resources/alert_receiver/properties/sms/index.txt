---
page_title: "sms"
subcategory: ""
description: "SMS Configuration."
xcsh_docs: {"aliases": ["sms"], "body_bytes": 1530, "body_sha256": "sha256:a7cdb85f5efe6e9d56a55bc1adf6be880dd31fd8f18e3d98ba02eb5dd2331ac8", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:sms", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "documentation/resources/alert_receiver/properties/sms/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2321322330122322-3323101322120202-0332000331221112-1333103013101001-3330200333023112-3123201202122121-1121033313113332-2312331010332202", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sms"], "schema_version": 1, "sections": [{"aliases": ["sms contact number"], "anchor": "schema-sms--contact_number", "description": "Contact number of the user in ITU E.164 format", "document_id": "xcsh-docs:resources:alert_receiver:properties:sms", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sms", "contact_number"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/sms/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SMS Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sms

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- sms

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
sms {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-sms--contact_number"></a>

### contact_number property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
