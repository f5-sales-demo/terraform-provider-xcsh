---
page_title: "endpoint.outside_endpoints"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["endpoint outside endpoints"], "body_bytes": 1274, "body_sha256": "sha256:19d1ca16453dd319510599d37497cf985c00689445c5559e464b994b1abf23f8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:endpoint:outside_endpoints", "parent_id": "xcsh-docs:resources:network_policy_view:properties:endpoint", "path": "documentation/resources/network_policy_view/properties/endpoint/outside_endpoints/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2112000022120120-0200101111220011-3013033011113300-0332020332303221-3030013232221120-1130322330031322-3210003023003333-0001030220223321", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint", "outside_endpoints"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/endpoint/outside_endpoints/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint.outside_endpoints

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/endpoint/)
- endpoint.outside_endpoints

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
outside_endpoints = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/endpoint/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
