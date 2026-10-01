---
page_title: "custom_network_config.active_network_policies"
subcategory: ""
description: "custom_network_config.active_network_policies for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1308, "body_sha256": "sha256:0490227810da0eaf9f7423e49b62ddaef5ffe18c854479df9e53ee217d74ec21", "canonical_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_network_policies", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_network_policies:network_policies"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_network_policies", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "path": "docs/guides/data-sources--securemesh_site--properties--custom_network_config--active_network_policies.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.active_network_policies for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.active_network_policies

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
- [Property reference](data-sources--securemesh_site--reference.md)
- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md)
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

- [network_policies](data-sources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [custom_network_config.active_network_policies.network_policies](data-sources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md)
- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
