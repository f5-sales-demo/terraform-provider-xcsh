---
page_title: "virtual_server.auto_last_hop"
subcategory: ""
description: "virtual_server.auto_last_hop for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 3462, "body_sha256": "sha256:278328e94b69841ab108e7ae47097da9a80d095a3693e87947d33b7599b6256c", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server--auto_last_hop.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "auto_last_hop"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/auto_last_hop/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.auto_last_hop for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.auto_last_hop

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- virtual_server.auto_last_hop

<a id="section"></a>

Type: `"single"`. Computed.

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auto_last_hop_choice": "[\"auto_last_hop_default\",\"auto_last_hop_disable\",\"auto_last_hop_enable\"]"
}
```

## Direct properties

- [auto_last_hop_default](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_default.md): complete subsection reference.

- [auto_last_hop_disable](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_disable.md): complete subsection reference.

- [auto_last_hop_enable](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_enable.md): complete subsection reference.

## Next pages

- [virtual_server.auto_last_hop.auto_last_hop_default](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_default.md)
- [virtual_server.auto_last_hop.auto_last_hop_disable](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_disable.md)
- [virtual_server.auto_last_hop.auto_last_hop_enable](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_enable.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
