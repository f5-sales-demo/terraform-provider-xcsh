---
page_title: "filters"
subcategory: ""
description: "filters for xcsh_device_intelligence_device_summary."
xcsh_docs: {"aliases": [], "body_bytes": 1301, "body_sha256": "sha256:aa1220e30c33bb13ff8f0df83c3f2f4c182d00414eb1a64df41f9e9cd29ec0e8", "canonical_id": "xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters:global_filters"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "path": "docs/guides/data-sources--device_intelligence_device_summary--properties--filters.md", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/properties/filters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filters for xcsh_device_intelligence_device_summary.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# filters

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md)
- [Property reference](data-sources--device_intelligence_device_summary--reference.md)
- filters

<a id="section"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

## Direct properties

- [global_filters](data-sources--device_intelligence_device_summary--properties--filters--global_filters.md): complete subsection reference.

<a id="schema-filters--region_filter"></a>

### region_filter property

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

## Next pages

- [filters.global_filters](data-sources--device_intelligence_device_summary--properties--filters--global_filters.md)
- [Property reference](data-sources--device_intelligence_device_summary--reference.md)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md)
