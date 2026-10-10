---
page_title: "equinix.not_managed.node_list"
subcategory: ""
description: "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section."
xcsh_docs: {"aliases": ["equinix not managed node list"], "body_bytes": 5429, "body_sha256": "sha256:e81d87fd72e627019244751f708e6aa1d2d00bad93ea5c7c6e7a29510736b7ec", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed", "path": "documentation/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["equinix", "not_managed", "node_list"], "schema_version": 1, "sections": [{"aliases": ["equinix not managed node list hostname"], "anchor": "schema-equinix--not_managed--node_list--hostname", "description": "Hostname for this Node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["equinix not managed node list interface list"], "anchor": "section", "description": "Manage interfaces belonging to this node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:bond_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:bond_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,dhcp_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,dhcp_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ethernet_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ethernet_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ethernet_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,no_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:monitor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:monitor_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv4_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv4_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv4_address,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv4_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,no_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv6_address,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:site_to_site_connectivity_interface_disabled,site_to_site_connectivity_interface_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:site_to_site_connectivity_interface_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:site_to_site_connectivity_interface_disabled,site_to_site_connectivity_interface_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:site_to_site_connectivity_interface_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv4_address,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv6_address,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ethernet_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "type": "conflicts"}], "schema_path": ["equinix", "not_managed", "node_list", "interface_list"], "syntax": "block", "type": "object"}, {"aliases": ["equinix not managed node list public ip"], "anchor": "schema-equinix--not_managed--node_list--public_ip", "description": "Public IP for this Node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "public_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["equinix not managed node list type"], "anchor": "schema-equinix--not_managed--node_list--type", "description": "Type for this Node, can be Control or Worker.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["Control", "Worker"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# equinix.not_managed.node_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [equinix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/)
- [equinix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/)
- equinix.not_managed.node_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-equinix--not_managed--node_list--hostname"></a>

### hostname property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/): complete subsection reference.

<a id="schema-equinix--not_managed--node_list--public_ip"></a>

### public_ip property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-equinix--not_managed--node_list--type"></a>

### type property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Control","Worker"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```
