---
page_title: "voltstack_cluster.global_network_list"
subcategory: "Infrastructure"
description: "voltstack_cluster.global_network_list for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1813, "body_sha256": "sha256:a5f9e1299d85dba8e639e5b91d01b7519f9e6a3dd6351de12a8055498d7be3d0", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/index.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["voltstack_cluster", "global_network_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.global_network_list for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.global_network_list

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.global_network_list

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

- [global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/): complete subsection reference.

## Next pages

- [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
