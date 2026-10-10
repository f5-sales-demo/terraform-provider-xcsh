---
page_title: "bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata"
subcategory: "Load Balancing"
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insertion rules exclude list metadata"], "body_bytes": 3884, "body_sha256": "sha256:439d48eaf1c283bc40e8eb8bc5d80e3a1000a442f3272567ebeced1c0ce80651", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:exclude_list:metadata", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:exclude_list", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/exclude_list/metadata/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1011020020300202-2332203021211013-0023222120212201-0101123013013310-1311310301020303-3021200102103310-1103312210021013-3102330112113101", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insertion_rules", "exclude_list", "metadata"], "schema_version": 1, "sections": [{"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insertion rules exclude list metadata description spec"], "anchor": "schema-bot_defense_advanced_protection--web_only--js_insertion_rules--exclude_list--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:exclude_list:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insertion_rules", "exclude_list", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insertion rules exclude list metadata name"], "anchor": "schema-bot_defense_advanced_protection--web_only--js_insertion_rules--exclude_list--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:exclude_list:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insertion_rules", "exclude_list", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/exclude_list/metadata/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/exclude_list/)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bot_defense_advanced_protection--web_only--js_insertion_rules--exclude_list--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="schema-bot_defense_advanced_protection--web_only--js_insertion_rules--exclude_list--metadata--name"></a>

### name property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```
