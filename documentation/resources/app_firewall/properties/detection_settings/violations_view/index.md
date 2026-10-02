---
page_title: "detection_settings.violations_view"
subcategory: "Security"
description: "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations."
xcsh_docs: {"aliases": ["detection settings violations view"], "body_bytes": 4774, "body_sha256": "sha256:26396578f689a8f94ffa57276bb7036843a79d9fbc36da6eda63a7bce1c46170", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "path": "documentation/resources/app_firewall/properties/detection_settings/violations_view/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2223210232020100-1101120331032303-2012301231233223-1012101301133313-3303013111120223-3032101020232203-1021012202102300-2132131102313021", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "violations_view"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-detection_settings--violations_view--description_spec", "description": "Description. Human-readable description text", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["enabled"], "anchor": "schema-detection_settings--violations_view--enabled", "description": "Enable or disable the feature", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["enabled by default"], "anchor": "schema-detection_settings--violations_view--enabled_by_default", "description": "Violations that are enabled by default by F5 are advisable to leave enabled.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "enabled_by_default"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-detection_settings--violations_view--name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["title"], "anchor": "schema-detection_settings--violations_view--title", "description": "Human-readable title for the resource", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "violations_view", "title"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/violations_view/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.violations_view

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/)
- detection_settings.violations_view

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
violations_view {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-detection_settings--violations_view--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human-readable description text

<a id="schema-detection_settings--violations_view--enabled"></a>

### enabled property

Type: `"bool"`. Optional.

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

Type: `"string"`. Optional.

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

Type: `"string"`. Optional.

Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

Type: `"string"`. Optional.

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

- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
