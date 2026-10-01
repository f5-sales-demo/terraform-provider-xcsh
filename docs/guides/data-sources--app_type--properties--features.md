---
page_title: "features"
subcategory: ""
description: "features for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 2578, "body_sha256": "sha256:831b44595277ccc44bab38dd1b746cd15bac51628bc70b0019e9c3c3263ba8e7", "canonical_id": "xcsh-docs:data-sources:app_type:properties:features", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_type:properties:features", "parent_id": "xcsh-docs:data-sources:app_type:reference", "path": "docs/guides/data-sources--app_type--properties--features.md", "provider_name": "app_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["features"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/properties/features/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "features for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# features

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md)
- [Property reference](data-sources--app_type--reference.md)
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Enumeration for advanced security features supported

API Discovery enables generation of model for various API interactions between services of App type.
Enable analysis of timeseries for various metric collected like requests, errors, latency etc.
Enable anomaly detection per API request, i.e. The probability density function (PDF) charts
generation for API endpoints Enable user behavior analysis.

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

## Next pages

- [Property reference](data-sources--app_type--reference.md)
- [xcsh_app_type](../data-sources/app_type.md)
