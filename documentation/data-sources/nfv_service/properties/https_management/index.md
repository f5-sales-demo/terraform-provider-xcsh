---
page_title: "https_management"
subcategory: ""
description: "https_management for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 5828, "body_sha256": "sha256:e636734be985358dc3dcc3a5bb0a6c2b99e75a5e56ee5d8d986ec960cfada47e", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet_default_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:default_https_port"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management", "parent_id": "xcsh-docs:data-sources:nfv_service:reference", "path": "documentation/data-sources/nfv_service/properties/https_management/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["https_management"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- https_management

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for https management.

Upstream description:

HTTPS based configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_internet\",\"advertise_on_internet_default_vip\",\"advertise_on_sli_vip\",\"advertise_on_slo_internet_vip\",\"advertise_on_slo_sli\",\"advertise_on_slo_vip\"]",
  "x-ves-oneof-field-internet_choice": "[]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"https_port\"]"
}
```

## Direct properties

- [advertise_on_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_internet/): complete subsection reference.

- [advertise_on_internet_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_internet_default_vip/): complete subsection reference.

- [advertise_on_sli_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/): complete subsection reference.

- [advertise_on_slo_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/): complete subsection reference.

- [advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/): complete subsection reference.

- [advertise_on_slo_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/): complete subsection reference.

- [default_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/default_https_port/): complete subsection reference.

<a id="schema-https_management--domain_suffix"></a>

### domain_suffix property

Type: `"string"`. Computed.

Domain suffix will be used along with node name to form URL to access node management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-https_management--https_port"></a>

### https_port property

Type: `"number"`. Computed.

Exclusive with \[default\_https\_port\] Enter TCP port number.

Upstream description:

Exclusive with \[default\_https\_port\] Enter TCP port number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [https_management.advertise_on_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_internet/)
- [https_management.advertise_on_internet_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_internet_default_vip/)
- [https_management.advertise_on_sli_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/)
- [https_management.advertise_on_slo_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/)
- [https_management.advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/)
- [https_management.advertise_on_slo_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/)
- [https_management.default_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/default_https_port/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
