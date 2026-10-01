---
page_title: "dedicated_interface.monitor"
subcategory: ""
description: "dedicated_interface.monitor for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1001, "body_sha256": "sha256:3646202b7b614b961851ed397681248e22c4e71629f3a79ac39e8bbae74c8c2f", "canonical_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "child_ids": [], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "parent_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "path": "docs/guides/resources--network_interface--properties--dedicated_interface--monitor.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dedicated_interface", "monitor"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_interface/monitor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dedicated_interface.monitor for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_interface.monitor

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [dedicated_interface](resources--network_interface--properties--dedicated_interface.md)
- dedicated_interface.monitor

<a id="section"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dedicated_interface](resources--network_interface--properties--dedicated_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
