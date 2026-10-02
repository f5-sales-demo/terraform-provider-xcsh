---
page_title: "features"
subcategory: ""
description: "List of various advanced security features enabled."
xcsh_docs: {"aliases": ["features"], "body_bytes": 3141, "body_sha256": "sha256:e128866084d236d82057b8858bbb074c34f6d82518bbd7a285b3dc0f9b5c9b61", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_type:properties:features", "parent_id": "xcsh-docs:resources:app_type:reference", "path": "documentation/resources/app_type/properties/features/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1311121001333100-0000000111012131-2121330230031311-1213111133111000-3123303122012022-0203021310311320-3322202132100311-1222302220321220", "registry_path": "docs/guides/resources--app_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["features"], "schema_version": 1, "sections": [{"aliases": ["type"], "anchor": "schema-features--type", "description": "Enumeration for advanced security features supported API Discovery enables generation of model for various API interactions between services of App type. Enable analysis of timeseries for various metric collected like requests, errors, latency etc. Enable anomaly detection per API request, i.e. The probability density", "document_id": "xcsh-docs:resources:app_type:properties:features", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["features", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/properties/features/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of various advanced security features enabled.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# features

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/)
- features

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
features {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-features--type"></a>

### type property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"),
}
```

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/)
- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
