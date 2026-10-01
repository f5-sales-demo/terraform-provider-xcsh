---
page_title: "pagination"
subcategory: ""
description: "pagination for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": [], "body_bytes": 1062, "body_sha256": "sha256:5620ef97ed8ee03bb88eb74329c932ba955376b217ccc30ad9923f9409fe26d2", "canonical_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "path": "docs/guides/data-sources--device_intelligence_device_history--properties--pagination.md", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["pagination"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/pagination/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "pagination for xcsh_device_intelligence_device_history.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pagination

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md)
- [Property reference](data-sources--device_intelligence_device_history--reference.md)
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

- [Property reference](data-sources--device_intelligence_device_history--reference.md)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md)
