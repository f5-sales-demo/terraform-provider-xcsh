---
page_title: "openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip"
subcategory: ""
description: "Configure Static IP parameters for a node."
xcsh_docs: {"aliases": ["openstack not managed node list interface list static ipv6 address node static ip"], "body_bytes": 4580, "body_sha256": "sha256:5afacb327f15e6b3bc7808d495317a559bd7eb2c8fb2d2c97c7b8c073a5bbe41", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address", "path": "documentation/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3023020130101312-2030330013132000-0020030020300102-3131211312001203-2222322302130311-2222233121220012-3003330123213000-0031033102330231", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [{"anchor": "schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip"], "schema_version": 1, "sections": [{"aliases": ["openstack not managed node list interface list static ipv6 address node static ip default gw"], "anchor": "schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--default_gw", "description": "IP address of the default gateway.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip", "default_gw"], "syntax": "attribute", "type": "string"}, {"aliases": ["openstack not managed node list interface list static ipv6 address node static ip dns server"], "anchor": "schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--dns_server", "description": "DNS server address for the static interface configuration.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip", "dns_server"], "syntax": "attribute", "type": "string"}, {"aliases": ["openstack not managed node list interface list static ipv6 address node static ip ip address"], "anchor": "schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "description": "IP address of the interface and prefix length.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip", "ip_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Configure Static IP parameters for a node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/)
- [openstack.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/)
- [openstack.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/)
- [openstack.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/static_ipv6_address/)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--default_gw"></a>

### default_gw property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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

<a id="schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--dns_server"></a>

### dns_server property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="schema-openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address"></a>

### ip_address property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
