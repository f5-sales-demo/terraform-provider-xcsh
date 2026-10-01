---
page_title: "virtual_server.nat64"
subcategory: ""
description: "virtual_server.nat64 for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 3222, "body_sha256": "sha256:2791745f81e3c637c344dffd4677537a70c51411cc19cdd7099f4aaf416fb9e7", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:nat64:nat64_disable", "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64:nat64_enable"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--nat64.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "nat64"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/nat64/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.nat64 for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.nat64

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.nat64

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("nat64_disable",
    "nat64_enable")}
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
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

Terraform syntax:

```terraform
nat64 {
  # Configure direct properties listed below.
}
```

## Direct properties

- [nat64_disable](resources--application_profiles--properties--virtual_server--nat64--nat64_disable.md): complete subsection reference.

- [nat64_enable](resources--application_profiles--properties--virtual_server--nat64--nat64_enable.md): complete subsection reference.

## Next pages

- [virtual_server.nat64.nat64_disable](resources--application_profiles--properties--virtual_server--nat64--nat64_disable.md)
- [virtual_server.nat64.nat64_enable](resources--application_profiles--properties--virtual_server--nat64--nat64_enable.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
