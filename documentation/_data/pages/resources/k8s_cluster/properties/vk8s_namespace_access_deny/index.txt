---
page_title: "vk8s_namespace_access_deny"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["vk8s namespace access deny"], "body_bytes": 1440, "body_sha256": "sha256:a396043b7c459a3e7f26f1822c4cfb7e14c4efea47c009b6d9e85a6ef7fc30ec", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:vk8s_namespace_access_deny", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/vk8s_namespace_access_deny/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2210013331002303-1110100111123032-0010132331001203-0120302213302320-1110231110012230-3113330311303112-3232113213313233-1023013222120221", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vk8s_namespace_access_deny"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/vk8s_namespace_access_deny/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

- [vk8s_namespace_access_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/vk8s_namespace_access_deny/#section)
- [vk8s_namespace_access_permit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/vk8s_namespace_access_permit/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
vk8s_namespace_access_deny = {}
```

This is an empty object or choice marker. It has no direct properties.
