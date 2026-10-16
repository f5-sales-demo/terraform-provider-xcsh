---
page_title: "endpoint.inside_endpoints"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["endpoint inside endpoints"], "body_bytes": 1002, "body_sha256": "sha256:a1172e8d8adb506729e7aae0736449aedc462126feef362de7986b54018a735f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:endpoint:inside_endpoints", "parent_id": "xcsh-docs:resources:network_policy_view:properties:endpoint", "path": "documentation/resources/network_policy_view/properties/endpoint/inside_endpoints/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0212031231300103-0131103033232231-3123111111003223-3110320333310210-0323303213000233-1320011322331022-0330333130232000-1010202322012210", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint", "inside_endpoints"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/endpoint/inside_endpoints/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint.inside_endpoints

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/endpoint/)
- endpoint.inside_endpoints

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.
