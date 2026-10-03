---
page_title: "vmware.not_managed.node_list"
subcategory: ""
description: "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section."
xcsh_docs: {"aliases": ["vmware not managed node list"], "body_bytes": 5809, "body_sha256": "sha256:294ff0eba2587d5ae8b044b02b85e024f961502fd728bc08e75a558e107d324c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed", "path": "documentation/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vmware", "not_managed", "node_list"], "schema_version": 1, "sections": [{"aliases": ["vmware not managed node list hostname"], "anchor": "schema-vmware--not_managed--node_list--hostname", "description": "Hostname for this Node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["vmware not managed node list interface list"], "anchor": "section", "description": "Manage interfaces belonging to this node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,dhcp_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,dhcp_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ethernet_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,no_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:monitor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:monitor_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv4_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,no_ipv4_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv4_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv4_address,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv4_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,no_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv6_address,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:site_to_site_connectivity_interface_disabled,site_to_site_connectivity_interface_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:site_to_site_connectivity_interface_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:site_to_site_connectivity_interface_disabled,site_to_site_connectivity_interface_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:site_to_site_connectivity_interface_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_client,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:dhcp_server,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv4_address,static_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ipv6_auto_config,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:no_ipv6_address,static_ipv6_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:bond_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:vlan_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list:ConflictingListObjectAttributes:ethernet_interface,vlan_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:vlan_interface", "type": "conflicts"}], "schema_path": ["vmware", "not_managed", "node_list", "interface_list"], "syntax": "block", "type": "object"}, {"aliases": ["vmware not managed node list public ip"], "anchor": "schema-vmware--not_managed--node_list--public_ip", "description": "Public IP for this Node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "public_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["vmware not managed node list type"], "anchor": "schema-vmware--not_managed--node_list--type", "description": "Type for this Node, can be Control or Worker.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [vmware](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/)
- [vmware.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/)
- vmware.not_managed.node_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-vmware--not_managed--node_list--hostname"></a>

### hostname property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--public_ip"></a>

### public_ip property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-vmware--not_managed--node_list--type"></a>

### type property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [vmware.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/)
- [vmware.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
