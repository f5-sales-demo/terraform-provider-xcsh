---
page_title: "custom_anonymization.anonymization_config.cookie"
subcategory: "Security"
description: "custom_anonymization.anonymization_config.cookie for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 2782, "body_sha256": "sha256:0112b7edfc277fdf5a890244e943e8ba2a09478db62d14f3a0364f0334c525f8", "canonical_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "docs/guides/resources--app_firewall--properties--custom_anonymization--anonymization_config--cookie.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "cookie"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_anonymization.anonymization_config.cookie for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.cookie

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- [custom_anonymization](resources--app_firewall--properties--custom_anonymization.md)
- [custom_anonymization.anonymization_config](resources--app_firewall--properties--custom_anonymization--anonymization_config.md)
- custom_anonymization.anonymization_config.cookie

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Cookies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_name")}
```

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
cookie {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_anonymization--anonymization_config--cookie--cookie_name"></a>

### cookie_name property

Type: `"string"`. Optional.

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Upstream description:

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [custom_anonymization.anonymization_config](resources--app_firewall--properties--custom_anonymization--anonymization_config.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
