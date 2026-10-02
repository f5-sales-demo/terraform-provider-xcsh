---
page_title: "http2_options"
subcategory: ""
description: "Http2 Protocol OPTIONS for upstream connections."
xcsh_docs: {"aliases": ["http2 options"], "body_bytes": 1313, "body_sha256": "sha256:3702971a6446d81ff4ea64dc45ad2d41ef0d6b611850e7e105914d5f3e081056", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http2_options", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/http2_options/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1122100310003003-3113331120023312-1013121113333323-1300201130311202-1110110033320221-1200012120120123-3131112312310130-0313133200131313", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http2_options"], "schema_version": 1, "sections": [{"aliases": ["enabled"], "anchor": "schema-http2_options--enabled", "description": "Enable/disable HTTP2 Protocol for upstream connections.", "document_id": "xcsh-docs:resources:cluster:properties:http2_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http2_options", "enabled"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http2_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Http2 Protocol OPTIONS for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http2_options

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- http2_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http2_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-http2_options--enabled"></a>

### enabled property

Type: `"bool"`. Optional.

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
