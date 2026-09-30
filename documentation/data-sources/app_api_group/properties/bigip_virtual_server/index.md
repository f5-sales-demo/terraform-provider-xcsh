---
page_title: "bigip_virtual_server"
subcategory: ""
description: "bigip_virtual_server for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1935, "body_sha256": "sha256:b5546ddb1c1f81877067578e69e10648bdc1b493458e06364a47bd875bc21d61", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server"], "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server", "parent_id": "xcsh-docs:data-sources:app_api_group:reference", "path": "documentation/data-sources/app_api_group/properties/bigip_virtual_server/index.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bigip_virtual_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/bigip_virtual_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bigip_virtual_server for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bigip_virtual_server

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
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

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/bigip_virtual_server/#section)
- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/cdn_loadbalancer/#section)
- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/http_loadbalancer/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/): complete subsection reference.

## Next pages

- [bigip_virtual_server.bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
