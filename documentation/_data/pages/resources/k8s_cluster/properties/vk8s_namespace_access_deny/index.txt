---
page_title: "vk8s_namespace_access_deny"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["vk8s namespace access deny"], "body_bytes": 1686, "body_sha256": "sha256:1b5290312221bf8f8798cb21819069d10dc0d4821a8127c876e354745e6194ad", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:vk8s_namespace_access_deny", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/vk8s_namespace_access_deny/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2210013331002303-1110100111123032-0010132331001203-0120302213302320-1110231110012230-3113330311303112-3232113213313233-1023013222120221", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vk8s_namespace_access_deny"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/vk8s_namespace_access_deny/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vk8s_namespace_access_deny

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- vk8s_namespace_access_deny

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: vk8s\_namespace\_access\_deny, vk8s\_namespace\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [vk8s_namespace_access_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/vk8s_namespace_access_deny/#section)
- [vk8s_namespace_access_permit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/vk8s_namespace_access_permit/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
vk8s_namespace_access_deny = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
