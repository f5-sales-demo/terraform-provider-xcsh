---
page_title: "data_guard_rules"
subcategory: "Load Balancing"
description: "Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature."
xcsh_docs: {"aliases": ["data guard rules"], "body_bytes": 4726, "body_sha256": "sha256:91317918c72a1e1c3a3dcc82564b949afc2cdae873c7fc105cc182bc40fe0a2e", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:apply_data_guard", "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:path", "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:skip_data_guard"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/data_guard_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data_guard_rules"], "schema_version": 1, "sections": [{"aliases": ["data guard rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["data guard rules apply data guard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:apply_data_guard", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "apply_data_guard"], "syntax": "attribute", "type": "object"}, {"aliases": ["data guard rules exact value"], "anchor": "schema-data_guard_rules--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["data guard rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data_guard_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["data guard rules path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data_guard_rules", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["data guard rules skip data guard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules:skip_data_guard", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "skip_data_guard"], "syntax": "attribute", "type": "object"}, {"aliases": ["data guard rules suffix value"], "anchor": "schema-data_guard_rules--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:data_guard_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/data_guard_rules/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_guard_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- data_guard_rules

<a id="section"></a>

Type: `"list"`. Computed.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/data_guard_rules/any_domain/): complete subsection reference.

- [apply_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/data_guard_rules/apply_data_guard/): complete subsection reference.

<a id="schema-data_guard_rules--exact_value"></a>

### exact_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/data_guard_rules/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/data_guard_rules/path/): complete subsection reference.

- [skip_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/data_guard_rules/skip_data_guard/): complete subsection reference.

<a id="schema-data_guard_rules--suffix_value"></a>

### suffix_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
