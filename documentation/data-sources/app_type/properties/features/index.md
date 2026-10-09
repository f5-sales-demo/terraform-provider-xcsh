---
page_title: "features"
subcategory: ""
description: "List of various advanced security features enabled."
xcsh_docs: {"aliases": ["features"], "body_bytes": 2423, "body_sha256": "sha256:c9ac85ef4b91ec33135e22509f3188953d4e13d0f90b2d2b95b9a68b2845aded", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_type:properties:features", "parent_id": "xcsh-docs:data-sources:app_type:reference", "path": "documentation/data-sources/app_type/properties/features/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3022321230021323-1100011333312000-3101300310230021-2220221313113330-0300330220330011-1133023330331010-1100131101013300-3031001121200112", "registry_path": "docs/guides/data-sources--app_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["features"], "schema_version": 1, "sections": [{"aliases": ["features type"], "anchor": "schema-features--type", "description": "Enumeration for advanced security features supported API Discovery enables generation of model for various API interactions between services of App type. Enable analysis of timeseries for various metric collected like requests, errors, latency etc. Enable anomaly detection per API request, i.e. The probability density", "document_id": "xcsh-docs:data-sources:app_type:properties:features", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["features", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/properties/features/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of various advanced security features enabled.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_typeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# features

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/)
- features

<a id="section"></a>

Type: `"list"`. Computed.

List of various advanced security features enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-features--type"></a>

### type property

Type: `"string"`. Computed.

\[Enum:
BUSINESS\_LOGIC\_MARKUP|TIMESERIES\_ANOMALY\_DETECTION|PER\_REQ\_ANOMALY\_DETECTION|USER\_BEHAVIOR\_ANALYSIS\]
Enumeration for advanced security features supported API Discovery enables generation of model for
various API interactions between services of App type. Enable analysis of timeseries for various
metric collected like requests, errors, latency etc. Enable anomaly detection per API request, i.e.
Possible values are \`BUSINESS\_LOGIC\_MARKUP\`, \`TIMESERIES\_ANOMALY\_DETECTION\`,
\`PER\_REQ\_ANOMALY\_DETECTION\`, \`USER\_BEHAVIOR\_ANALYSIS\`. Defaults to
\`BUSINESS\_LOGIC\_MARKUP\`.

Additional upstream details:

Enumeration for advanced security features supported

API Discovery enables generation of model for various API interactions between services of App type.
The probability density function (PDF) charts generation for API endpoints Enable user behavior
analysis.

Receipt-pinned upstream constraints:

```json
{
  "default": "BUSINESS_LOGIC_MARKUP",
  "enum": [
    "BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
