---
page_title: "virtual_server.auto_last_hop"
subcategory: ""
description: "When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a"
xcsh_docs: {"aliases": ["virtual server auto last hop"], "body_bytes": 4013, "body_sha256": "sha256:f35f0e17f3ebe99f99bc8df6fb59982b3af92fbc3cafce61473e920e0b7d2604", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/auto_last_hop/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1213231221312103-2101111203001231-0201213300231211-0313012112021033-1012200101220121-2101220302131103-3212333133311002-3200300331321033", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "auto_last_hop"], "schema_version": 1, "sections": [{"aliases": ["virtual server auto last hop auto last hop default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server auto last hop auto last hop disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server auto last hop auto last hop enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/auto_last_hop/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.auto_last_hop

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
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

- [auto_last_hop_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/): complete subsection reference.

- [auto_last_hop_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_disable/): complete subsection reference.

- [auto_last_hop_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_enable/): complete subsection reference.

## Next pages

- [virtual_server.auto_last_hop.auto_last_hop_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/)
- [virtual_server.auto_last_hop.auto_last_hop_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_disable/)
- [virtual_server.auto_last_hop.auto_last_hop_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_enable/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
