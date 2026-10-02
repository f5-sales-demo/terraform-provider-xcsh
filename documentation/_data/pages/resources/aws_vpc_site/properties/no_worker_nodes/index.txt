---
page_title: "no_worker_nodes"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no worker nodes"], "body_bytes": 1699, "body_sha256": "sha256:76f96480dfee7895cef71a72134deb0beeaf01c59e22c09100960fa02d91349c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:no_worker_nodes", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/no_worker_nodes/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0323111213120210-3330110331133331-1123320133233101-0112302203111100-3323311230111302-2322033010012210-0303221110112120-1110302200110131", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_worker_nodes"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/no_worker_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_worker_nodes

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- no_worker_nodes

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_worker\_nodes, nodes\_per\_az, total\_nodes; Default: no\_worker\_nodes\] Configuration
parameter for no worker nodes.

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

- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/no_worker_nodes/#section)
- [nodes_per_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/#schema-nodes_per_az)
- [total_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/#schema-total_nodes)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_worker_nodes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
