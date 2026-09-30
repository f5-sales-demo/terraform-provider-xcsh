---
page_title: "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers"
subcategory: "Load Balancing"
description: "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5350, "body_sha256": "sha256:360dbafa976c84e9c641d0e060a41363affaab339de31aaab79def700137be13", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_not_present", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_present", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:item"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier", "headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier.md)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--check_not_present.md): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--check_present.md): complete subsection reference.

- [item](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--item.md): complete subsection reference.

<a id="schema-bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--name"></a>

### name property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--check_not_present.md)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--check_present.md)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--item.md)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
