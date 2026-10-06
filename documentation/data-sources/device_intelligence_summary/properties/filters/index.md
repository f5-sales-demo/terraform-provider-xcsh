---
page_title: "filters"
subcategory: ""
description: "Global Filters. Query Global Filters."
xcsh_docs: {"aliases": ["filters"], "body_bytes": 1465, "body_sha256": "sha256:fb9cd62e184c3dd4ef250c761b21bacd73e66e77e5ab1264a3dba29c6c0bf537", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_summary:properties:filters:global_filters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "path": "documentation/data-sources/device_intelligence_summary/properties/filters/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_summary", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0310002002331232-2312333120230003-0000332312002113-1212213332001111-1322320003331030-0033011301031002-0011120310231111-1300301022011032", "registry_path": "docs/guides/data-sources--device_intelligence_summary--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filters"], "schema_version": 1, "sections": [{"aliases": ["filters global filters"], "anchor": "section", "description": "Global Filters. List of global filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters:global_filters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["filters", "global_filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["filters region filter"], "anchor": "schema-filters--region_filter", "description": "Defines a selection for Bot Defense region - US: US United States of America - EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are `US`, `EU`, `ASIA`, `CA`. Defaults to `US`.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ASIA", "CA", "EU", "US"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "region_filter"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_summary/properties/filters/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Global Filters. Query Global Filters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filters

Breadcrumbs:

- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/)
- filters

<a id="section"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

## Direct properties

- [global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/): complete subsection reference.

<a id="schema-filters--region_filter"></a>

### region_filter property

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASIA","CA","EU","US"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```
