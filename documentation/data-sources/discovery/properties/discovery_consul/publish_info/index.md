---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "Consul Configuration to publish VIPs."
xcsh_docs: {"aliases": ["discovery consul publish info"], "body_bytes": 1903, "body_sha256": "sha256:47458f4d872ddee5d8892cbc2e08c17af20355c1da5011fccdb4eab8d0373726", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:publish"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "path": "documentation/data-sources/discovery/properties/discovery_consul/publish_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3022323202130331-1300101212130033-0001310133112233-0230110021230003-0022213333130131-1230010030020111-1232120213113231-3303010023321313", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "sections": [{"aliases": ["discovery consul publish info disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "publish_info", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery consul publish info publish"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:publish", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "publish_info", "publish"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Consul Configuration to publish VIPs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.publish_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- discovery_consul.publish_info

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Upstream description:

Consul Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/disable_spec/): complete subsection reference.

- [publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/publish/): complete subsection reference.

## Next pages

- [discovery_consul.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/disable_spec/)
- [discovery_consul.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/publish/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
