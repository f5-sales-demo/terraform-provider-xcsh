---
page_title: "oci.not_managed.node_list.interface_list.static_ipv6_address"
subcategory: ""
description: "oci.not_managed.node_list.interface_list.static_ipv6_address for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2519, "body_sha256": "sha256:3b9bdbfef5200305fdeb909d7906abebabf2a186881b77126c68cbbf1e4e7c0c", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "path": "docs/guides/resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--static_ipv6_address.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oci.not_managed.node_list.interface_list.static_ipv6_address for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list.interface_list.static_ipv6_address

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [oci](resources--securemesh_site_v2--properties--oci.md)
- [oci.not_managed](resources--securemesh_site_v2--properties--oci--not_managed.md)
- [oci.not_managed.node_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list.md)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- oci.not_managed.node_list.interface_list.static_ipv6_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip.md): complete subsection reference.

## Next pages

- [oci.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip.md)
- [oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip.md)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
