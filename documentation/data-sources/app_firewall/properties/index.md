---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_app_firewall."
xcsh_docs: {"aliases": ["app firewall"], "body_bytes": 29247, "body_sha256": "sha256:a2ef3e735edbfc1d3d123fa1a94e00ebc54be4304e6e88493dc959b79d077b9d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:allow_all_response_codes", "xcsh-docs:data-sources:app_firewall:properties:allowed_response_codes", "xcsh-docs:data-sources:app_firewall:properties:blocking", "xcsh-docs:data-sources:app_firewall:properties:blocking_page", "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "xcsh-docs:data-sources:app_firewall:properties:default_anonymization", "xcsh-docs:data-sources:app_firewall:properties:default_bot_setting", "xcsh-docs:data-sources:app_firewall:properties:default_detection_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "xcsh-docs:data-sources:app_firewall:properties:disable_ai_enhancements", "xcsh-docs:data-sources:app_firewall:properties:disable_anonymization", "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements", "xcsh-docs:data-sources:app_firewall:properties:monitoring", "xcsh-docs:data-sources:app_firewall:properties:use_default_blocking_page"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:reference", "parent_id": "xcsh-docs:data-sources:app_firewall:fundamentals", "path": "documentation/data-sources/app_firewall/properties/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allow all response codes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:allow_all_response_codes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_all_response_codes"], "syntax": "attribute", "type": "object"}, {"aliases": ["allowed response codes"], "anchor": "section", "description": "List of HTTP response status codes that are allowed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:allowed_response_codes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowed_response_codes"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:app_firewall:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["blocking"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:blocking", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocking"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocking page"], "anchor": "section", "description": "Custom blocking response page body.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:blocking_page", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["blocking_page"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot protection setting"], "anchor": "section", "description": "Configuration of WAF Bot Protection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_protection_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom anonymization"], "anchor": "section", "description": "Anonymization settings which is a list of HTTP headers, parameters and cookies.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_anonymization"], "syntax": "attribute", "type": "object"}, {"aliases": ["default anonymization"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:default_anonymization", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_anonymization"], "syntax": "attribute", "type": "object"}, {"aliases": ["default bot setting"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:default_bot_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_bot_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["default detection settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:default_detection_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_detection_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:app_firewall:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings"], "anchor": "section", "description": "Specifies detection settings to be used by WAF.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable ai enhancements"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:disable_ai_enhancements", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_ai_enhancements"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable anonymization"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:disable_anonymization", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_anonymization"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable ai enhancements"], "anchor": "section", "description": "Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_ai_enhancements"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:app_firewall:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:app_firewall:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["monitoring"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:monitoring", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["monitoring"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:app_firewall:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:app_firewall:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["use default blocking page"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:use_default_blocking_page", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_default_blocking_page"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- Property reference

## Direct properties

- [allow_all_response_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allow_all_response_codes/): complete subsection reference.

- [allowed_response_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allowed_response_codes/): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking/): complete subsection reference.

- [blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking_page/): complete subsection reference.

- [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/): complete subsection reference.

- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/): complete subsection reference.

- [default_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_anonymization/): complete subsection reference.

- [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_bot_setting/): complete subsection reference.

- [default_detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_detection_settings/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AppFirewall.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/): complete subsection reference.

- [disable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_ai_enhancements/): complete subsection reference.

- [disable_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_anonymization/): complete subsection reference.

- [enable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [monitoring](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/monitoring/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AppFirewall.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the AppFirewall exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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

- [use_default_blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/use_default_blocking_page/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_response_codes` | [allow_all_response_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allow_all_response_codes/#section) |
| `allowed_response_codes` | [allowed_response_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allowed_response_codes/#section) |
| `allowed_response_codes.response_code` | [allowed_response_codes.response_code](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allowed_response_codes/#schema-allowed_response_codes--response_code) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/#schema-annotations) |
| `blocking` | [blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking/#section) |
| `blocking_page` | [blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking_page/#section) |
| `blocking_page.blocking_page` | [blocking_page.blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking_page/#schema-blocking_page--blocking_page) |
| `blocking_page.response_code` | [blocking_page.response_code](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking_page/#schema-blocking_page--response_code) |
| `bot_protection_setting` | [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/#section) |
| `bot_protection_setting.good_bot_action` | [bot_protection_setting.good_bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/#schema-bot_protection_setting--good_bot_action) |
| `bot_protection_setting.malicious_bot_action` | [bot_protection_setting.malicious_bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/#schema-bot_protection_setting--malicious_bot_action) |
| `bot_protection_setting.suspicious_bot_action` | [bot_protection_setting.suspicious_bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/#schema-bot_protection_setting--suspicious_bot_action) |
| `custom_anonymization` | [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/#section) |
| `custom_anonymization.anonymization_config` | [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/#section) |
| `custom_anonymization.anonymization_config.cookie` | [custom_anonymization.anonymization_config.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/#section) |
| `custom_anonymization.anonymization_config.cookie.cookie_name` | [custom_anonymization.anonymization_config.cookie.cookie_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/#schema-custom_anonymization--anonymization_config--cookie--cookie_name) |
| `custom_anonymization.anonymization_config.http_header` | [custom_anonymization.anonymization_config.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/#section) |
| `custom_anonymization.anonymization_config.http_header.header_name` | [custom_anonymization.anonymization_config.http_header.header_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/#schema-custom_anonymization--anonymization_config--http_header--header_name) |
| `custom_anonymization.anonymization_config.query_parameter` | [custom_anonymization.anonymization_config.query_parameter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/#section) |
| `custom_anonymization.anonymization_config.query_parameter.query_param_name` | [custom_anonymization.anonymization_config.query_parameter.query_param_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/#schema-custom_anonymization--anonymization_config--query_parameter--query_param_name) |
| `default_anonymization` | [default_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_anonymization/#section) |
| `default_bot_setting` | [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_bot_setting/#section) |
| `default_detection_settings` | [default_detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_detection_settings/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/#schema-description) |
| `detection_settings` | [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/#section) |
| `detection_settings.bot_protection_setting` | [detection_settings.bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/#section) |
| `detection_settings.bot_protection_setting.good_bot_action` | [detection_settings.bot_protection_setting.good_bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/#schema-detection_settings--bot_protection_setting--good_bot_action) |
| `detection_settings.bot_protection_setting.malicious_bot_action` | [detection_settings.bot_protection_setting.malicious_bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/#schema-detection_settings--bot_protection_setting--malicious_bot_action) |
| `detection_settings.bot_protection_setting.suspicious_bot_action` | [detection_settings.bot_protection_setting.suspicious_bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/#schema-detection_settings--bot_protection_setting--suspicious_bot_action) |
| `detection_settings.default_bot_setting` | [detection_settings.default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/default_bot_setting/#section) |
| `detection_settings.default_violation_settings` | [detection_settings.default_violation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/default_violation_settings/#section) |
| `detection_settings.disable_staging` | [detection_settings.disable_staging](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_staging/#section) |
| `detection_settings.disable_suppression` | [detection_settings.disable_suppression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_suppression/#section) |
| `detection_settings.disable_threat_campaigns` | [detection_settings.disable_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_threat_campaigns/#section) |
| `detection_settings.enable_suppression` | [detection_settings.enable_suppression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/enable_suppression/#section) |
| `detection_settings.enable_threat_campaigns` | [detection_settings.enable_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/enable_threat_campaigns/#section) |
| `detection_settings.signature_selection_setting` | [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/#section) |
| `detection_settings.signature_selection_setting.attack_type_settings` | [detection_settings.signature_selection_setting.attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/#section) |
| `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` | [detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/#schema-detection_settings--signature_selection_setting--attack_type_settings--disabled_attack_types) |
| `detection_settings.signature_selection_setting.default_attack_type_settings` | [detection_settings.signature_selection_setting.default_attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_attack_type_settings/#section) |
| `detection_settings.signature_selection_setting.default_signature_setting` | [detection_settings.signature_selection_setting.default_signature_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/#section) |
| `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_accuracy_signatures/#section) |
| `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_low_accuracy_signatures/#section) |
| `detection_settings.signature_selection_setting.only_high_accuracy_signatures` | [detection_settings.signature_selection_setting.only_high_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/only_high_accuracy_signatures/#section) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/#section) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/#schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--high_accuracy_action) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/#schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--low_accuracy_action) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/#schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--medium_accuracy_action) |
| `detection_settings.stage_new_and_updated_signatures` | [detection_settings.stage_new_and_updated_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_and_updated_signatures/#section) |
| `detection_settings.stage_new_and_updated_signatures.staging_period` | [detection_settings.stage_new_and_updated_signatures.staging_period](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_and_updated_signatures/#schema-detection_settings--stage_new_and_updated_signatures--staging_period) |
| `detection_settings.stage_new_signatures` | [detection_settings.stage_new_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_signatures/#section) |
| `detection_settings.stage_new_signatures.staging_period` | [detection_settings.stage_new_signatures.staging_period](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_signatures/#schema-detection_settings--stage_new_signatures--staging_period) |
| `detection_settings.violation_settings` | [detection_settings.violation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violation_settings/#section) |
| `detection_settings.violation_settings.disabled_violation_types` | [detection_settings.violation_settings.disabled_violation_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violation_settings/#schema-detection_settings--violation_settings--disabled_violation_types) |
| `detection_settings.violations_view` | [detection_settings.violations_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/#section) |
| `detection_settings.violations_view.description_spec` | [detection_settings.violations_view.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/#schema-detection_settings--violations_view--description_spec) |
| `detection_settings.violations_view.enabled` | [detection_settings.violations_view.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/#schema-detection_settings--violations_view--enabled) |
| `detection_settings.violations_view.enabled_by_default` | [detection_settings.violations_view.enabled_by_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/#schema-detection_settings--violations_view--enabled_by_default) |
| `detection_settings.violations_view.name` | [detection_settings.violations_view.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/#schema-detection_settings--violations_view--name) |
| `detection_settings.violations_view.title` | [detection_settings.violations_view.title](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/#schema-detection_settings--violations_view--title) |
| `disable_ai_enhancements` | [disable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_ai_enhancements/#section) |
| `disable_anonymization` | [disable_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_anonymization/#section) |
| `enable_ai_enhancements` | [enable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/#section) |
| `enable_ai_enhancements.mitigate_high_medium_risk_action` | [enable_ai_enhancements.mitigate_high_medium_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/mitigate_high_medium_risk_action/#section) |
| `enable_ai_enhancements.mitigate_high_risk_action` | [enable_ai_enhancements.mitigate_high_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/mitigate_high_risk_action/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/#schema-labels) |
| `monitoring` | [monitoring](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/monitoring/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/#schema-namespace) |
| `use_default_blocking_page` | [use_default_blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/use_default_blocking_page/#section) |

## Next pages

- [allow_all_response_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allow_all_response_codes/)
- [allowed_response_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/allowed_response_codes/)
- [blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking/)
- [blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/blocking_page/)
- [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/)
- [default_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_anonymization/)
- [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_bot_setting/)
- [default_detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_detection_settings/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [disable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_ai_enhancements/)
- [disable_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_anonymization/)
- [enable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/)
- [monitoring](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/monitoring/)
- [use_default_blocking_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/use_default_blocking_page/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
