---
page_title: "dedicated_interface.cluster"
subcategory: ""
description: "Applies dedicated interface configuration cluster-wide across the site rather than to a specific node."
xcsh_docs: {"aliases": ["cluster scope", "cluster wide interface", "dedicated interface cluster"], "body_bytes": 1296, "body_sha256": "sha256:b140a304bedd45987f3fb56c54b4a7709f883385bc50f276c42d6deb91ed2682", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule", "reviewed-summary"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:cluster", "parent_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "path": "documentation/resources/network_interface/properties/dedicated_interface/cluster/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0213100202031002-3003330112113330-2022302121333132-3233311103010020-2020133223123011-2303121203030301-3210330331100101-0203022231111013", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_interface", "cluster"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_interface/cluster/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Applies dedicated interface configuration cluster-wide across the site rather than to a specific node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_interface.cluster

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/)
- dedicated_interface.cluster

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
cluster = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
