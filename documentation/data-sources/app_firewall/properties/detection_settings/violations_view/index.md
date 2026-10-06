---
page_title: "detection_settings.violations_view"
subcategory: "Security"
description: "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations."
xcsh_docs: {"aliases": ["detection settings violations view"], "body_bytes": 3996, "body_sha256": "sha256:e65f362dc3fd1838667d93fe4260476cfc3958155ca978a115ce11060d9be414", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "path": "documentation/data-sources/app_firewall/properties/detection_settings/violations_view/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3100112002213131-1121100000222220-1222121201303032-3033012303001222-2102330311310130-1021103301002112-0221013203310100-2322122000201311", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "violations_view"], "schema_version": 1, "sections": [{"aliases": ["detection settings violations view description spec"], "anchor": "schema-detection_settings--violations_view--description_spec", "description": "Description. Human-readable description text", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings violations view enabled"], "anchor": "schema-detection_settings--violations_view--enabled", "description": "Enable or disable the feature", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["detection settings violations view enabled by default"], "anchor": "schema-detection_settings--violations_view--enabled_by_default", "description": "Violations that are enabled by default by F5 are advisable to leave enabled.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "enabled_by_default"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings violations view name"], "anchor": "schema-detection_settings--violations_view--name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings violations view title"], "anchor": "schema-detection_settings--violations_view--title", "description": "Human-readable title for the resource", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "title"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/violations_view/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.violations_view

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- detection_settings.violations_view

<a id="section"></a>

Type: `"list"`. Computed.

List of violation checks that are performed on HTTP request to ensure the requests are properly
formatted, detection of evasion techniques and other violations.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Direct properties

<a id="schema-detection_settings--violations_view--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="schema-detection_settings--violations_view--enabled"></a>

### enabled property

Type: `"bool"`. Computed.

State. Enable or disable the feature

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

<a id="schema-detection_settings--violations_view--enabled_by_default"></a>

### enabled_by_default property

Type: `"string"`. Computed.

Violations that are enabled by default by F5 are advisable to leave enabled.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-detection_settings--violations_view--name"></a>

### name property

Type: `"string"`. Computed.

Name. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-detection_settings--violations_view--title"></a>

### title property

Type: `"string"`. Computed.

Title. Human-readable title for the resource

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
