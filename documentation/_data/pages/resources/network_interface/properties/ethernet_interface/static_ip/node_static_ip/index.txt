---
page_title: "ethernet_interface.static_ip.node_static_ip"
subcategory: ""
description: "Configure Static IP parameters for a node."
xcsh_docs: {"aliases": ["ethernet interface static ip node static ip"], "body_bytes": 4067, "body_sha256": "sha256:179d7c37a054b61f22ebcc53c18f40d684ebb909f4a9be652daeb7ba21d77f7f", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip", "path": "documentation/resources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2303101021211120-3010222011201312-2113001133320111-0301010000312330-1320120233212331-0302201211232023-1303310100102320-3023001113030131", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "schema-ethernet_interface--static_ip--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "ethernet_interface.static_ip.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "static_ip", "node_static_ip"], "schema_version": 1, "sections": [{"aliases": ["default gw"], "anchor": "schema-ethernet_interface--static_ip--node_static_ip--default_gw", "description": "IP address of the default gateway.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "static_ip", "node_static_ip", "default_gw"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns server"], "anchor": "schema-ethernet_interface--static_ip--node_static_ip--dns_server", "description": "DNS server address for the static interface configuration.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "static_ip", "node_static_ip", "dns_server"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip address"], "anchor": "schema-ethernet_interface--static_ip--node_static_ip--ip_address", "description": "IP address of the interface and prefix length.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "static_ip", "node_static_ip", "ip_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure Static IP parameters for a node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.static_ip.node_static_ip

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/static_ip/)
- ethernet_interface.static_ip.node_static_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
node_static_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ethernet_interface--static_ip--node_static_ip--default_gw"></a>

### default_gw property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-ethernet_interface--static_ip--node_static_ip--dns_server"></a>

### dns_server property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="schema-ethernet_interface--static_ip--node_static_ip--ip_address"></a>

### ip_address property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

## Next pages

- [ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/static_ip/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
