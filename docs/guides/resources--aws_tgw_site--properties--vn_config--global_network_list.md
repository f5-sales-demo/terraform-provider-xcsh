---
page_title: "vn_config.global_network_list"
subcategory: ""
description: "vn_config.global_network_list for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1386, "body_sha256": "sha256:4214647b99b35c47bacb4ef51aadf211c6910c1c0a87526b5891a127815cc84c", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--global_network_list.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "global_network_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.global_network_list for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vn_config.global_network_list

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- vn_config.global_network_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
```

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
global_network_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
