---
page_title: "custom_anonymization.anonymization_config.cookie"
subcategory: "Security"
description: "Configure anonymization for HTTP Cookies."
xcsh_docs: {"aliases": ["custom anonymization anonymization config cookie"], "body_bytes": 2599, "body_sha256": "sha256:ed2281180fbe58391907e6fd4166c98b41d7448181be6aa380747f14dd4852ed", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/resources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "resources", "registry_anchor": "canonical-2101000230003012-2330321113123112-0322310021020121-2201101001123233-2022132320210033-1201223100123233-0121120101321001-3121013323013030", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--cookie--cookie_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.cookie:RequiredObjectAttributes:cookie_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "cookie"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config cookie cookie name"], "anchor": "schema-custom_anonymization--anonymization_config--cookie--cookie_name", "description": "Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by prefixing or suffixing the cookie name with a wildcard asterisk (*), or by using only an asterisk to match any cookie name.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "cookie", "cookie_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Configure anonymization for HTTP Cookies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["app_firewallCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.cookie

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.cookie

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Cookies.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-13T13:24:52+00:00"
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
