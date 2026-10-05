---
page_title: "local_access_config.default_port"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["local access config default port"], "body_bytes": 1269, "body_sha256": "sha256:67978e6187f69fa4624cfcba7460668581e0475892e0c35b6d9c708eb0b8697e", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config:default_port", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "path": "documentation/resources/k8s_cluster/properties/local_access_config/default_port/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2203333201103021-0122002320130220-3131113321302033-3312332231312001-0111130130301101-3301101203103331-1120231312033310-2301323113210212", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_access_config", "default_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/local_access_config/default_port/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_access_config.default_port

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [local_access_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/local_access_config/)
- local_access_config.default_port

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
default_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_access_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/local_access_config/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
