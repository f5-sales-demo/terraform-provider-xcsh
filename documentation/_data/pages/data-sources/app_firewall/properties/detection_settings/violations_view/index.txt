---
page_title: "detection_settings.violations_view"
subcategory: "Security"
description: "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations."
xcsh_docs: {"aliases": ["detection settings violations view"], "body_bytes": 4432, "body_sha256": "sha256:bfc1dfc578b15f1d4164e2aec253f3c917adec4f53358acf375e163bac962f7d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "path": "documentation/data-sources/app_firewall/properties/detection_settings/violations_view/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3100112002213131-1121100000222220-1222121201303032-3033012303001222-2102330311310130-1021103301002112-0221013203310100-2322122000201311", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "violations_view"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-detection_settings--violations_view--description_spec", "description": "Description. Human-readable description text", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["enabled"], "anchor": "schema-detection_settings--violations_view--enabled", "description": "Enable or disable the feature", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["enabled by default"], "anchor": "schema-detection_settings--violations_view--enabled_by_default", "description": "Violations that are enabled by default by F5 are advisable to leave enabled.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "enabled_by_default"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-detection_settings--violations_view--name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["title"], "anchor": "schema-detection_settings--violations_view--title", "description": "Human-readable title for the resource", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "title"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/violations_view/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Enable or disable the feature

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Human-readable name for the resource

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Human-readable title for the resource

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
