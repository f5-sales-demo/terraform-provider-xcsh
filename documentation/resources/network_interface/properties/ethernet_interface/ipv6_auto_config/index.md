---
page_title: "ethernet_interface.ipv6_auto_config"
subcategory: ""
description: "IPV6AutoConfigType."
xcsh_docs: {"aliases": ["ethernet interface ipv6 auto config"], "body_bytes": 1576, "body_sha256": "sha256:13aae18071f08da26645065880592a9212c56f01193e26d340b9a1089beba559", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "documentation/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config"], "schema_version": 1, "sections": [{"aliases": ["ethernet interface ipv6 auto config host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "ipv6_auto_config", "host"], "syntax": "attribute", "type": "object"}, {"aliases": ["ethernet interface ipv6 auto config router"], "anchor": "section", "description": "IPV6AutoConfigRouterType.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ethernet_interface--ipv6_auto_config--router--network_prefix", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful", "type": "conflicts"}], "schema_path": ["ethernet_interface", "ipv6_auto_config", "router"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "IPV6AutoConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- ethernet_interface.ipv6_auto_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/host/): complete subsection reference.

- [router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/): complete subsection reference.
