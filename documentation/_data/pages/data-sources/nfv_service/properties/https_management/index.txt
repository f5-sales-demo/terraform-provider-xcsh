---
page_title: "https_management"
subcategory: ""
description: "HTTPS based configuration."
xcsh_docs: {"aliases": ["https management"], "body_bytes": 4251, "body_sha256": "sha256:3f7f523d1720a6019e351a2670fbf3acec1ab7c7a4826dbc163018525c8fcd06", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet_default_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip", "xcsh-docs:data-sources:nfv_service:properties:https_management:default_https_port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management", "parent_id": "xcsh-docs:data-sources:nfv_service:reference", "path": "documentation/data-sources/nfv_service/properties/https_management/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on internet"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_internet"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on internet default vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet_default_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_internet_default_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on sli vip"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo internet vip"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_slo_internet_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo sli"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo vip"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_slo_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management default https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:default_https_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "default_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management domain suffix"], "anchor": "schema-https_management--domain_suffix", "description": "Domain suffix will be used along with node name to form URL to access node management.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "domain_suffix"], "syntax": "attribute", "type": "string"}, {"aliases": ["https management https port"], "anchor": "schema-https_management--https_port", "description": "Exclusive with Enter TCP port number.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "https_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "HTTPS based configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
