---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tunnel."
xcsh_docs: {"aliases": ["tunnel"], "body_bytes": 20683, "body_sha256": "sha256:468be3eaf8b1569055025d4d74d7a66660d80c4a94e9e547ed32ef90a9480df0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip", "xcsh-docs:resources:tunnel:properties:params", "xcsh-docs:resources:tunnel:properties:remote_ip", "xcsh-docs:resources:tunnel:properties:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:reference", "parent_id": "xcsh-docs:resources:tunnel:fundamentals", "path": "documentation/resources/tunnel/properties/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["local ip"], "anchor": "section", "description": "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip:ConflictingObjectAttributes:intf,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:intf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip:ConflictingObjectAttributes:intf,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "type": "conflicts"}], "schema_path": ["local_ip"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["params"], "anchor": "section", "description": "Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which PSK can be configured.", "document_id": "xcsh-docs:resources:tunnel:properties:params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["params"], "syntax": "block", "type": "object"}, {"aliases": ["remote ip"], "anchor": "section", "description": "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip:ConflictingObjectAttributes:endpoints,ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip:ConflictingObjectAttributes:endpoints,ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "type": "conflicts"}], "schema_path": ["remote_ip"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:tunnel:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["tunnel type"], "anchor": "schema-tunnel_type", "description": "Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel type.", "document_id": "xcsh-docs:resources:tunnel:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tunnel_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Tunnel. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Tunnel is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/): complete subsection reference.

- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/): complete subsection reference.

<a id="schema-tunnel_type"></a>

### tunnel_type property

Type: `"string"`. Optional, Computed.

\[Enum: IPSEC\_PSK|GRE\] Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel
type. Possible values are \`IPSEC\_PSK\`, \`GRE\`. Defaults to \`IPSEC\_PSK\`.

Upstream description:

Supported tunnel types are IPsec

IPsec tunnel type with PSK GRE tunnel type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IPSEC_PSK",
    "GRE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "IPSEC_PSK",
  "enum": [
    "IPSEC_PSK",
    "GRE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-labels) |
| `local_ip` | [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/#section) |
| `local_ip.intf` | [local_ip.intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/#section) |
| `local_ip.intf.local_intf` | [local_ip.intf.local_intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/#section) |
| `local_ip.intf.local_intf.kind` | [local_ip.intf.local_intf.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/#schema-local_ip--intf--local_intf--kind) |
| `local_ip.intf.local_intf.name` | [local_ip.intf.local_intf.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/#schema-local_ip--intf--local_intf--name) |
| `local_ip.intf.local_intf.namespace` | [local_ip.intf.local_intf.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/#schema-local_ip--intf--local_intf--namespace) |
| `local_ip.intf.local_intf.tenant` | [local_ip.intf.local_intf.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/#schema-local_ip--intf--local_intf--tenant) |
| `local_ip.intf.local_intf.uid` | [local_ip.intf.local_intf.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/local_intf/#schema-local_ip--intf--local_intf--uid) |
| `local_ip.ip_address` | [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/#section) |
| `local_ip.ip_address.auto` | [local_ip.ip_address.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/auto/#section) |
| `local_ip.ip_address.ip_address` | [local_ip.ip_address.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/#section) |
| `local_ip.ip_address.ip_address.dual_stack` | [local_ip.ip_address.ip_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/#section) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4` | [local_ip.ip_address.ip_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv4/#section) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv4/#schema-local_ip--ip_address--ip_address--dual_stack--ipv4--addr) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6` | [local_ip.ip_address.ip_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv6/#section) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv6/#schema-local_ip--ip_address--ip_address--dual_stack--ipv6--addr) |
| `local_ip.ip_address.ip_address.ipv4` | [local_ip.ip_address.ip_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/ipv4/#section) |
| `local_ip.ip_address.ip_address.ipv4.addr` | [local_ip.ip_address.ip_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/ipv4/#schema-local_ip--ip_address--ip_address--ipv4--addr) |
| `local_ip.ip_address.ip_address.ipv6` | [local_ip.ip_address.ip_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/ipv6/#section) |
| `local_ip.ip_address.ip_address.ipv6.addr` | [local_ip.ip_address.ip_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/ipv6/#schema-local_ip--ip_address--ip_address--ipv6--addr) |
| `local_ip.ip_address.virtual_network_type` | [local_ip.ip_address.virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/#section) |
| `local_ip.ip_address.virtual_network_type.public` | [local_ip.ip_address.virtual_network_type.public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/#section) |
| `local_ip.ip_address.virtual_network_type.site_local` | [local_ip.ip_address.virtual_network_type.site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/#section) |
| `local_ip.ip_address.virtual_network_type.site_local_inside` | [local_ip.ip_address.virtual_network_type.site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-namespace) |
| `params` | [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/#section) |
| `params.ipsec` | [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/#section) |
| `params.ipsec.ipsec_psk` | [params.ipsec.ipsec_psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/#section) |
| `params.ipsec.ipsec_psk.blindfold_secret_info` | [params.ipsec.ipsec_psk.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/#section) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/#schema-params--ipsec--ipsec_psk--blindfold_secret_info--decryption_provider) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.location` | [params.ipsec.ipsec_psk.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/#schema-params--ipsec--ipsec_psk--blindfold_secret_info--location) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/#schema-params--ipsec--ipsec_psk--blindfold_secret_info--store_provider) |
| `params.ipsec.ipsec_psk.clear_secret_info` | [params.ipsec.ipsec_psk.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/#section) |
| `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` | [params.ipsec.ipsec_psk.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/#schema-params--ipsec--ipsec_psk--clear_secret_info--provider_ref) |
| `params.ipsec.ipsec_psk.clear_secret_info.url` | [params.ipsec.ipsec_psk.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/#schema-params--ipsec--ipsec_psk--clear_secret_info--url) |
| `remote_ip` | [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/#section) |
| `remote_ip.endpoints` | [remote_ip.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/endpoints/#section) |
| `remote_ip.endpoints.endpoints` | [remote_ip.endpoints.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/endpoints/endpoints/#section) |
| `remote_ip.ip` | [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/#section) |
| `remote_ip.ip.dual_stack` | [remote_ip.ip.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/#section) |
| `remote_ip.ip.dual_stack.ipv4` | [remote_ip.ip.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/#section) |
| `remote_ip.ip.dual_stack.ipv4.addr` | [remote_ip.ip.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/#schema-remote_ip--ip--dual_stack--ipv4--addr) |
| `remote_ip.ip.dual_stack.ipv6` | [remote_ip.ip.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv6/#section) |
| `remote_ip.ip.dual_stack.ipv6.addr` | [remote_ip.ip.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv6/#schema-remote_ip--ip--dual_stack--ipv6--addr) |
| `remote_ip.ip.ipv4` | [remote_ip.ip.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv4/#section) |
| `remote_ip.ip.ipv4.addr` | [remote_ip.ip.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv4/#schema-remote_ip--ip--ipv4--addr) |
| `remote_ip.ip.ipv6` | [remote_ip.ip.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv6/#section) |
| `remote_ip.ip.ipv6.addr` | [remote_ip.ip.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv6/#schema-remote_ip--ip--ipv6--addr) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/#schema-timeouts--update) |
| `tunnel_type` | [tunnel_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/#schema-tunnel_type) |

## Next pages

- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/timeouts/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
