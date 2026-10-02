---
page_title: "routes.simple_route.incoming_port.no_port_match"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes simple route incoming port no port match"], "body_bytes": 1654, "body_sha256": "sha256:f637ef6443d9946301ded945ba5454d898e9e1eb2571a293b629fa2a710ee363", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/no_port_match/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3112022330030001-1320010021111133-0022310012100020-0001032203312333-1120302230110301-2233301001022023-1222301132111323-0013031200120203", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "incoming_port", "no_port_match"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/no_port_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.incoming_port.no_port_match

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/)
- routes.simple_route.incoming_port.no_port_match

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
no_port_match = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
