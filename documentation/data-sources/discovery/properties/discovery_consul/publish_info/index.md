---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "discovery_consul.publish_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1903, "body_sha256": "sha256:47458f4d872ddee5d8892cbc2e08c17af20355c1da5011fccdb4eab8d0373726", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:publish"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "path": "documentation/data-sources/discovery/properties/discovery_consul/publish_info/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.publish_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
