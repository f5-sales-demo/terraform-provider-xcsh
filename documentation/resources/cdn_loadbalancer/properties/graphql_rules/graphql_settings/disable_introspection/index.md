---
page_title: "graphql_rules.graphql_settings.disable_introspection"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["graphql rules graphql settings disable introspection"], "body_bytes": 1221, "body_sha256": "sha256:527c3b54b878e5ff9ffd614ed11fe8212709597bee83b0a4b011649304c138d3", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "path": "documentation/resources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3011202003003032-0001310103133331-0111311023220230-3202133313001031-2011303203032233-1303320103110121-1232112230112301-2301230212301233", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules", "graphql_settings", "disable_introspection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.graphql_settings.disable_introspection

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/)
- [graphql_rules.graphql_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/)
- graphql_rules.graphql_settings.disable_introspection

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
disable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.
