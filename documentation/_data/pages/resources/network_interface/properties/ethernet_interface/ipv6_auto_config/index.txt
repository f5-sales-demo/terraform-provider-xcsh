---
page_title: "ethernet_interface.ipv6_auto_config"
subcategory: ""
description: "IPV6AutoConfigType."
xcsh_docs: {"aliases": ["ethernet interface ipv6 auto config"], "body_bytes": 2184, "body_sha256": "sha256:b3d12ba2cfcacd38992fd32e2f06554dfcf9aedbcb20e5babed8fce9a61a4cf0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "documentation/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config"], "schema_version": 1, "sections": [{"aliases": ["host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "ipv6_auto_config", "host"], "syntax": "attribute", "type": "object"}, {"aliases": ["router"], "anchor": "section", "description": "IPV6AutoConfigRouterType.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ethernet_interface--ipv6_auto_config--router--network_prefix", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful", "type": "conflicts"}], "schema_path": ["ethernet_interface", "ipv6_auto_config", "router"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6AutoConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [ethernet_interface.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/host/)
- [ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
