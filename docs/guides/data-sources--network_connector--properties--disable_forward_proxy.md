---
page_title: "disable_forward_proxy"
subcategory: "Networking"
description: "disable_forward_proxy for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1216, "body_sha256": "sha256:d203e6a1618abd158dc44ec3950d83e10dadc9e00d4e35d24703cfacacdd6bf3", "canonical_id": "xcsh-docs:data-sources:network_connector:properties:disable_forward_proxy", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:disable_forward_proxy", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "docs/guides/data-sources--network_connector--properties--disable_forward_proxy.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_forward_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/disable_forward_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_forward_proxy for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_forward_proxy

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md)
- [Property reference](data-sources--network_connector--reference.md)
- disable_forward_proxy

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_forward\_proxy, enable\_forward\_proxy; Default: disable\_forward\_proxy\]
Configuration parameter for disable forward proxy.

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

- [disable_forward_proxy](data-sources--network_connector--properties--disable_forward_proxy.md#section)
- [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--network_connector--reference.md)
- [xcsh_network_connector](../data-sources/network_connector.md)
