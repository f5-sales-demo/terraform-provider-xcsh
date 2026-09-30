---
page_title: "oci.not_managed.node_list.interface_list.ipv6_auto_config"
subcategory: ""
description: "oci.not_managed.node_list.interface_list.ipv6_auto_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2173, "body_sha256": "sha256:c95c9cb85b342013a48d829316b9a8d22142b249e7ab773a770d452664a6db7b", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:host", "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "path": "docs/guides/resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oci.not_managed.node_list.interface_list.ipv6_auto_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# oci.not_managed.node_list.interface_list.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [oci](resources--securemesh_site_v2--properties--oci.md)
- [oci.not_managed](resources--securemesh_site_v2--properties--oci--not_managed.md)
- [oci.not_managed.node_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list.md)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- oci.not_managed.node_list.interface_list.ipv6_auto_config

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

- [host](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--host.md): complete subsection reference.

- [router](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router.md): complete subsection reference.

## Next pages

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--host.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
