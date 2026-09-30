---
page_title: "ethernet_interface.ipv6_auto_config"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1632, "body_sha256": "sha256:208501025c9a91935e9c0f0123bdbec3faabc5835ea88bd38c10101dfa9f08d6", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--ipv6_auto_config.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ethernet_interface.ipv6_auto_config

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- ethernet_interface.ipv6_auto_config

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

- [host](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--host.md): complete subsection reference.

- [router](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.host](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--host.md)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
