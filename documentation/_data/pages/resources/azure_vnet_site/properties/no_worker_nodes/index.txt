---
page_title: "no_worker_nodes"
subcategory: "Infrastructure"
description: "no_worker_nodes for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1798, "body_sha256": "sha256:c994962b8abf9ece2132beed811e8b43fed2c3961c1a1aa29a917ac0574c1d43", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:no_worker_nodes", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/no_worker_nodes/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["no_worker_nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/no_worker_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_worker_nodes for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_worker_nodes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- no_worker_nodes

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: no\_worker\_nodes, nodes\_per\_az, total\_nodes; Default: no\_worker\_nodes\] Configuration
parameter for no worker nodes. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/no_worker_nodes/#section)
- [nodes_per_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/#schema-nodes_per_az)
- [total_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/#schema-total_nodes)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_worker_nodes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
