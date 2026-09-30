---
page_title: "pagination"
subcategory: ""
description: "pagination for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": [], "body_bytes": 921, "body_sha256": "sha256:0afab41b0a9e3d821fbb01bbdfe191c8c12a71d785e88368ade5393c8caef939", "canonical_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:pagination", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:properties:pagination", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "path": "docs/guides/data-sources--device_intelligence_devices--properties--pagination.md", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["pagination"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/pagination/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "pagination for xcsh_device_intelligence_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# pagination

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md)
- [Property reference](data-sources--device_intelligence_devices--reference.md)
- pagination

<a id="section"></a>

Type: `"single"`. Optional.

Pagination for Request with number and size.

## Direct properties

<a id="schema-pagination--page_number"></a>

### page_number property

Type: `"number"`. Optional.

Configuration parameter for page number.

<a id="schema-pagination--page_size"></a>

### page_size property

Type: `"number"`. Optional.

Page Size. Size or capacity specification

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

## Next pages

- [Property reference](data-sources--device_intelligence_devices--reference.md)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md)
