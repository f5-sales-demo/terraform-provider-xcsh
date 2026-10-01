---
page_title: "custom_network_config.active_network_policies"
subcategory: ""
description: "custom_network_config.active_network_policies for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1299, "body_sha256": "sha256:2b9a764866b452a707f974867b3c74c8ed8a965416078ac51e6dfa29497280b7", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_network_policies", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_network_policies:network_policies"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_network_policies", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config--active_network_policies.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.active_network_policies for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.active_network_policies

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- custom_network_config.active_network_policies

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

## Direct properties

- [network_policies](data-sources--voltstack_site--properties--custom_network_config--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [custom_network_config.active_network_policies.network_policies](data-sources--voltstack_site--properties--custom_network_config--active_network_policies--network_policies.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
