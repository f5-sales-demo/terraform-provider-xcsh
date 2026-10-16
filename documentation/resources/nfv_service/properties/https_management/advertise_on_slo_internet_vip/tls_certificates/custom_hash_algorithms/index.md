---
page_title: "https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms"
subcategory: ""
description: "Specifies the hash algorithms to be used."
xcsh_docs: {"aliases": ["https management advertise on slo internet vip tls certificates custom hash algorithms"], "body_bytes": 2863, "body_sha256": "sha256:2f0529526010819858bedeeb9ac59e51c8be09f5fc0188df84c59c88cdb5fa4d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "path": "documentation/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/custom_hash_algorithms/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3233132331203230-2011001123021320-3230231000120130-0111230000101322-3001330103210111-1101113230233232-1212002311301212-3300120122233132", "registry_path": "docs/guides/resources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "custom_hash_algorithms"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on slo internet vip tls certificates custom hash algorithms hash algorithms"], "anchor": "schema-https_management--advertise_on_slo_internet_vip--tls_certificates--custom_hash_algorithms--hash_algorithms", "description": "Ordered list of hash algorithms to be used.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "custom_hash_algorithms", "hash_algorithms"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Specifies the hash algorithms to be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/)
- https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-https_management--advertise_on_slo_internet_vip--tls_certificates--custom_hash_algorithms--hash_algorithms"></a>

### hash_algorithms property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
