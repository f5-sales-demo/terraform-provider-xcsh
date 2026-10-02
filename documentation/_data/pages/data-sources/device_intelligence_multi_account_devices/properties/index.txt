---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_multi_account_devices."
xcsh_docs: {"aliases": ["device intelligence multi account devices"], "body_bytes": 5085, "body_sha256": "sha256:6828014eecb3ef3eaa59d59ab4b31458256e8136530a8cb2a9584407504cf41f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters", "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:fundamentals", "path": "documentation/data-sources/device_intelligence_multi_account_devices/properties/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0002211303210131-3301122203222131-3121332000231220-3031103123001111-3332222231232120-3000111203301111-1311330311113231-3322102221001210", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["end time"], "anchor": "schema-end_time", "description": "End time of the query period Format: unix_timestamp|RFC 3339 Optional: If not specified, then the end_time will be evaluated to start_time+10m If start_time is not specified, then the end_time will be evaluated to <current time>.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters"], "anchor": "section", "description": "Global Filters. Query Global Filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["multi account devices"], "anchor": "section", "description": "Distribution of devices by account range buckets.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["multi_account_devices"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. Namespace name.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start time of the query period Format: unix_timestamp|RFC 3339 Optional: If not specified, then the start_time will be evaluated to end_time-10m If end_time is not specified, then the start_time will be evaluated to <current time>-10m.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_device_intelligence_multi_account_devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
- Property reference

## Direct properties

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

End time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
end\_time will be evaluated to start\_time+10m If start\_time is not specified, then the end\_time
will be evaluated to &lt;current time&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/): complete subsection reference.

- [multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
start\_time will be evaluated to end\_time-10m If end\_time is not specified, then the start\_time
will be evaluated to &lt;current time&gt;-10m.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/#schema-end_time) |
| `filters` | [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/#section) |
| `filters.global_filters` | [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/#section) |
| `filters.global_filters.key` | [filters.global_filters.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/#schema-filters--region_filter) |
| `multi_account_devices` | [multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/#section) |
| `multi_account_devices.account_range` | [multi_account_devices.account_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/#schema-multi_account_devices--account_range) |
| `multi_account_devices.device_count` | [multi_account_devices.device_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/#schema-multi_account_devices--device_count) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/#schema-namespace) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/#schema-start_time) |

## Next pages

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/)
- [multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/)
- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
