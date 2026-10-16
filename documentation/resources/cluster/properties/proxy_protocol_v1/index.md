---
page_title: "proxy_protocol_v1"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["proxy protocol v1"], "body_bytes": 858, "body_sha256": "sha256:53c3a50ec73383416dfa3d31e6a45db9b6fc73f228549c2f8e349644c6afc452", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:proxy_protocol_v1", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/proxy_protocol_v1/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0033223222002132-3130113312332000-3110121121032222-0311212300220331-2111030221111102-1103202003220133-3230012120130011-3222201230122200", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_protocol_v1"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/proxy_protocol_v1/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["clusterCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_protocol_v1

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- proxy_protocol_v1

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v1.

Additional upstream details:

This can be used for messages where no values are needed.

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
proxy_protocol_v1 = {}
```

This is an empty object or choice marker. It has no direct properties.
