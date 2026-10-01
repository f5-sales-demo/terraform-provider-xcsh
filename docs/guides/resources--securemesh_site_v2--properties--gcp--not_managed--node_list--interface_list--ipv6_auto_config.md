---
page_title: "gcp.not_managed.node_list.interface_list.ipv6_auto_config"
subcategory: ""
description: "gcp.not_managed.node_list.interface_list.ipv6_auto_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2272, "body_sha256": "sha256:1f555a32cddffaed67f860b17d308b62e8a254c76b9ec60e13ab6774f2d82c09", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:host", "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list", "path": "docs/guides/resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--ipv6_auto_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp.not_managed.node_list.interface_list.ipv6_auto_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.not_managed.node_list.interface_list.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [gcp](resources--securemesh_site_v2--properties--gcp.md)
- [gcp.not_managed](resources--securemesh_site_v2--properties--gcp--not_managed.md)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--properties--gcp--not_managed--node_list.md)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list.md)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [host](resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--ipv6_auto_config--host.md): complete subsection reference.

- [router](resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--ipv6_auto_config--router.md): complete subsection reference.

## Next pages

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--ipv6_auto_config--host.md)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
