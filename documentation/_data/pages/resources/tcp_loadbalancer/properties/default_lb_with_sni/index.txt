---
page_title: "default_lb_with_sni"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default lb with sni"], "body_bytes": 1727, "body_sha256": "sha256:43d5773c8494087a259b6ede13b7f79e8c82c5bf920e377ce6731a17554fa365", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:default_lb_with_sni", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "documentation/resources/tcp_loadbalancer/properties/default_lb_with_sni/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1233300113221313-3110203321303320-1132100212321211-3233111001220210-2202120330231303-1320322300203123-3132233222313123-1030123011100010", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_lb_with_sni"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/default_lb_with_sni/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_lb_with_sni

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- default_lb_with_sni

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

OneOf alternatives in this subsection:

- [default_lb_with_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/default_lb_with_sni/#section)
- [no_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/no_sni/#section)
- [sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/sni/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_lb_with_sni = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
