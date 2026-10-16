---
page_title: "graphql_rules.graphql_settings.disable_introspection"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["graphql rules graphql settings disable introspection"], "body_bytes": 1226, "body_sha256": "sha256:53bd8185865dde0e495bc7d9d99f8da4bb1c1e06c53c19fa7fdb44189b2f6705", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "path": "documentation/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules", "graphql_settings", "disable_introspection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.graphql_settings.disable_introspection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/)
- [graphql_rules.graphql_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/)
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
