---
page_title: "bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4127, "body_sha256": "sha256:3eb946cf75b01026bc050a031014814e08374edd13c734a335d16e02374fa490", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:rules:metadata", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:rules:metadata", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--rules--metadata.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insertion_rules", "rules", "metadata"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/rules/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules.md)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--rules.md)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

<a id="section"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

## Direct properties

<a id="schema-bot_defense_advanced_protection--web_only--js_insertion_rules--rules--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="schema-bot_defense_advanced_protection--web_only--js_insertion_rules--rules--metadata--name"></a>

### name property

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

## Next pages

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
