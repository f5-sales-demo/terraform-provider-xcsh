---
page_title: "lb_algorithm.round_robin"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["lb algorithm round robin"], "body_bytes": 1294, "body_sha256": "sha256:f59315f31d0fd3acea99dd84812c6c1e7a781c879c3ba57e9265e0c2c57615e1", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm:round_robin", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm", "path": "documentation/resources/bigip_http_proxy/properties/lb_algorithm/round_robin/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1032322331030110-3003003122311201-2203010003020020-0033132123302200-2213122211132021-0133023330113310-3102121331113002-0221302122230131", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["lb_algorithm", "round_robin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/lb_algorithm/round_robin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# lb_algorithm.round_robin

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/lb_algorithm/)
- lb_algorithm.round_robin

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

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
round_robin {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/lb_algorithm/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
