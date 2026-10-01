---
page_title: "disable_proxy_protocol"
subcategory: ""
description: "disable_proxy_protocol for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1702, "body_sha256": "sha256:d12b853603a11c4af8434c7b1d6596cd944f57bdd30213467e3cbdfb6793c5eb", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:disable_proxy_protocol", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/disable_proxy_protocol/index.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["disable_proxy_protocol"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/disable_proxy_protocol/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_proxy_protocol for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_proxy_protocol

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- disable_proxy_protocol

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/disable_proxy_protocol/#section)
- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/proxy_protocol_v1/#section)
- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/proxy_protocol_v2/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
