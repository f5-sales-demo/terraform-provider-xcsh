---
page_title: "http2_options"
subcategory: ""
description: "Http2 Protocol OPTIONS for upstream connections."
xcsh_docs: {"aliases": ["http2 options"], "body_bytes": 1203, "body_sha256": "sha256:ba0e80725ebacb3b33e39c39a9c2b843676a271f345a3f030d9c8b89daf113db", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:http2_options", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/http2_options/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1030010231200022-3322131033122210-1230033020312003-0111313202211202-3322302012332230-1211221323031332-1123230123011022-3232203321032103", "registry_path": "docs/guides/data-sources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http2_options"], "schema_version": 1, "sections": [{"aliases": ["enabled"], "anchor": "schema-http2_options--enabled", "description": "Enable/disable HTTP2 Protocol for upstream connections.", "document_id": "xcsh-docs:data-sources:cluster:properties:http2_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http2_options", "enabled"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/http2_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Http2 Protocol OPTIONS for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http2_options

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- http2_options

<a id="section"></a>

Type: `"single"`. Computed.

Http2 Protocol OPTIONS for upstream connections.

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

## Direct properties

<a id="schema-http2_options--enabled"></a>

### enabled property

Type: `"bool"`. Computed.

Enable/disable HTTP2 Protocol for upstream connections.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
