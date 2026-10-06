---
page_title: "virtual_server.auto_last_hop"
subcategory: ""
description: "When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a"
xcsh_docs: {"aliases": ["virtual server auto last hop"], "body_bytes": 3312, "body_sha256": "sha256:0227d92a2d5b0bda661535487c2cbb8de598e5d1e2901d89b9f5bb1f4629a147", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/auto_last_hop/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.auto_last_hop:ConflictingObjectAttributes:auto_last_hop_default,auto_last_hop_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.auto_last_hop:ConflictingObjectAttributes:auto_last_hop_default,auto_last_hop_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.auto_last_hop:ConflictingObjectAttributes:auto_last_hop_default,auto_last_hop_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.auto_last_hop:ConflictingObjectAttributes:auto_last_hop_disable,auto_last_hop_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.auto_last_hop:ConflictingObjectAttributes:auto_last_hop_default,auto_last_hop_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.auto_last_hop:ConflictingObjectAttributes:auto_last_hop_disable,auto_last_hop_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "auto_last_hop"], "schema_version": 1, "sections": [{"aliases": ["virtual server auto last hop auto last hop default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server auto last hop auto last hop disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_disable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server auto last hop auto last hop enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/auto_last_hop/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.auto_last_hop

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.auto_last_hop

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_last_hop_default",
    "auto_last_hop_disable"),
  validators.ConflictingObjectAttributes("auto_last_hop_default",
    "auto_last_hop_enable"),
  validators.ConflictingObjectAttributes("auto_last_hop_disable",
    "auto_last_hop_enable")}
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
  "x-ves-oneof-field-auto_last_hop_choice": "[\"auto_last_hop_default\",\"auto_last_hop_disable\",\"auto_last_hop_enable\"]"
}
```

Terraform syntax:

```terraform
auto_last_hop {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto_last_hop_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/): complete subsection reference.

- [auto_last_hop_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_disable/): complete subsection reference.

- [auto_last_hop_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_enable/): complete subsection reference.
