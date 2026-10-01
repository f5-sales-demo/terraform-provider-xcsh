---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "discovery_consul.publish_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1450, "body_sha256": "sha256:6ae70b402e0618640895c6a20af417e73c9a11e53dccb319f39385ea4c81aa5c", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:publish"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "path": "docs/guides/data-sources--discovery--properties--discovery_consul--publish_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.publish_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.publish_info

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_consul](data-sources--discovery--properties--discovery_consul.md)
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

- [disable_spec](data-sources--discovery--properties--discovery_consul--publish_info--disable_spec.md): complete subsection reference.

- [publish](data-sources--discovery--properties--discovery_consul--publish_info--publish.md): complete subsection reference.

## Next pages

- [discovery_consul.publish_info.disable_spec](data-sources--discovery--properties--discovery_consul--publish_info--disable_spec.md)
- [discovery_consul.publish_info.publish](data-sources--discovery--properties--discovery_consul--publish_info--publish.md)
- [discovery_consul](data-sources--discovery--properties--discovery_consul.md)
- [xcsh_discovery](../data-sources/discovery.md)
