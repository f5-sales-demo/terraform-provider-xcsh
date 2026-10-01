---
page_title: "bigip_virtual_server"
subcategory: ""
description: "bigip_virtual_server for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1573, "body_sha256": "sha256:e57770604842b184a8674bd0bd5f96316d3804e697aeed5ac98e128d22cf57cf", "canonical_id": "xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server"], "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server", "parent_id": "xcsh-docs:data-sources:app_api_group:reference", "path": "docs/guides/data-sources--app_api_group--properties--bigip_virtual_server.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bigip_virtual_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/bigip_virtual_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bigip_virtual_server for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bigip_virtual_server

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md)
- [Property reference](data-sources--app_api_group--reference.md)
- bigip_virtual_server

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

Upstream description:

Set the scope of the API Group to a specific BIG-IP Virtual Server.

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

- [bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server.md#section)
- [cdn_loadbalancer](data-sources--app_api_group--properties--cdn_loadbalancer.md#section)
- [http_loadbalancer](data-sources--app_api_group--properties--http_loadbalancer.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server--bigip_virtual_server.md): complete subsection reference.

## Next pages

- [bigip_virtual_server.bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server--bigip_virtual_server.md)
- [Property reference](data-sources--app_api_group--reference.md)
- [xcsh_app_api_group](../data-sources/app_api_group.md)
