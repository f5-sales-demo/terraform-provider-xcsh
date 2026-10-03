---
page_title: "enable_forward_proxy.no_interception"
subcategory: "Networking"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable forward proxy no interception"], "body_bytes": 1352, "body_sha256": "sha256:7054cda35f8fb555708c69c065673ec71615f66bf57e0c01e8b114d8517e1263", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:no_interception", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/no_interception/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1220011120312221-0033121031212311-1130212312100101-3302021232121031-1231021000210332-2311003113002230-3002313213221102-2110133322032333", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "no_interception"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/no_interception/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["network_connectorCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.no_interception

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- enable_forward_proxy.no_interception

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no interception.

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
no_interception = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
