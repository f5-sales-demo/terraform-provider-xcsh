---
page_title: "automatic_port"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["automatic port"], "body_bytes": 1346, "body_sha256": "sha256:7c1210c8330e42435e4db2847b9b6bfc035fa745610d14b3a2db716032976b26", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:automatic_port", "parent_id": "xcsh-docs:resources:origin_pool:reference", "path": "documentation/resources/origin_pool/properties/automatic_port/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1112022203223032-3323221233013132-1121203303031132-1021300130303111-2312333020233311-1132310003313023-1330012320112102-3213232312123011", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["automatic_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/automatic_port/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["origin_poolCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# automatic_port

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- automatic_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: automatic\_port, lb\_port, port\] Enable this option

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

OneOf alternatives in this subsection:

- [automatic_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/automatic_port/#section)
- [lb_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/lb_port/#section)
- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-port)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
automatic_port = {}
```

This is an empty object or choice marker. It has no direct properties.
