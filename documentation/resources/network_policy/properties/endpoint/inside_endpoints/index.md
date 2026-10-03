---
page_title: "endpoint.inside_endpoints"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["endpoint inside endpoints"], "body_bytes": 1236, "body_sha256": "sha256:8aad0e7f4fd03e3fc24cdb9f10e910f7b15c3755d6898e3ffab4ca343567cbf4", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "parent_id": "xcsh-docs:resources:network_policy:properties:endpoint", "path": "documentation/resources/network_policy/properties/endpoint/inside_endpoints/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0232222033301310-1011011021331010-0122120033333302-3300213020031332-2302231000210000-1231102022100310-3232223322113101-3313131200300031", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint", "inside_endpoints"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/inside_endpoints/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint.inside_endpoints

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/)
- endpoint.inside_endpoints

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
inside_endpoints = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
