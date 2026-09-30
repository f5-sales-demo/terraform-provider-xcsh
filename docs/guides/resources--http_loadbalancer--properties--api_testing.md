---
page_title: "api_testing"
subcategory: "Load Balancing"
description: "api_testing for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3347, "body_sha256": "sha256:bb6bbb2d731b0fd0c607bb9af9706e29fd6b0b9c7ace1fbc093773a7e3083499", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--api_testing.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_testing

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- api_testing

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_testing, disable\_api\_testing; Default: disable\_api\_testing\] API Testing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("every_day",
    "every_month"),
  validators.ConflictingObjectAttributes("every_day",
    "every_week"),
  validators.ConflictingObjectAttributes("every_month",
    "every_week")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-frequency_choice": "[\"every_day\",\"every_month\",\"every_week\"]"
}
```

OneOf alternatives in this subsection:

- [api_testing](resources--http_loadbalancer--properties--api_testing.md#section)
- [disable_api_testing](resources--http_loadbalancer--properties--disable_api_testing.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_testing {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_testing--custom_header_value"></a>

### custom_header_value property

Type: `"string"`. Optional.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [domains](resources--http_loadbalancer--properties--api_testing--domains.md): complete subsection reference.

- [every_day](resources--http_loadbalancer--properties--api_testing--every_day.md): complete subsection reference.

- [every_month](resources--http_loadbalancer--properties--api_testing--every_month.md): complete subsection reference.

- [every_week](resources--http_loadbalancer--properties--api_testing--every_week.md): complete subsection reference.

## Next pages

- [api_testing.domains](resources--http_loadbalancer--properties--api_testing--domains.md)
- [api_testing.every_day](resources--http_loadbalancer--properties--api_testing--every_day.md)
- [api_testing.every_month](resources--http_loadbalancer--properties--api_testing--every_month.md)
- [api_testing.every_week](resources--http_loadbalancer--properties--api_testing--every_week.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
