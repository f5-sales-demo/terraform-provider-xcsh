---
page_title: "graphql_rules.graphql_settings.disable_introspection"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["graphql rules graphql settings disable introspection"], "body_bytes": 1533, "body_sha256": "sha256:dfaafb585b538b438f264be5e433dd1606b4ab47ed3b1ce6767bdf8693253863", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "path": "documentation/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules", "graphql_settings", "disable_introspection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [graphql_rules.graphql_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
