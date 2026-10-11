---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_advertise_policy."
xcsh_docs: {"aliases": ["advertise policy"], "body_bytes": 33598, "body_sha256": "sha256:cc4ef7d822ea26c688ef79451a674ebdd413107d078ec57744a4684e9ad1ce94", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:dualstack", "xcsh-docs:data-sources:advertise_policy:properties:ipv4", "xcsh-docs:data-sources:advertise_policy:properties:ipv6", "xcsh-docs:data-sources:advertise_policy:properties:public_ip", "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters", "xcsh-docs:data-sources:advertise_policy:properties:where"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:reference", "parent_id": "xcsh-docs:data-sources:advertise_policy:fundamentals", "path": "documentation/data-sources/advertise_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111", "registry_path": "docs/guides/data-sources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["address"], "anchor": "schema-address", "description": "Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where contains a site or virtual site of type REGIONAL_EDGE or public network If not specified and \"where\" is specified with site or virtual site option, inside_vip or outside_vip specified in the site object will be used", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address"], "syntax": "attribute", "type": "string"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["dualstack"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:dualstack", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dualstack"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv4"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:ipv6", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv6"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-port", "description": "Exclusive with Port to advertise.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port"], "syntax": "attribute", "type": "number"}, {"aliases": ["port ranges"], "anchor": "schema-port_ranges", "description": "Exclusive with A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["protocol"], "anchor": "schema-protocol", "description": "Protocol to advertise.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["public ip"], "anchor": "section", "description": "Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["skip xff append"], "anchor": "schema-skip_xff_append", "description": "If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.", "document_id": "xcsh-docs:data-sources:advertise_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["skip_xff_append"], "syntax": "attribute", "type": "bool"}, {"aliases": ["tls parameters"], "anchor": "section", "description": "TLS configuration for downstream connections.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["where"], "anchor": "section", "description": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:where", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_advertise_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Computed.

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
'where' is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site..

Additional upstream details:

This VIP can be either V4/V6 address You can not specify this if where contains a site or virtual
site of type REGIONAL\_EDGE or public network If not specified and "where" is specified with site or
virtual site option, inside\_vip or outside\_vip specified in the site object will be used based on
the network type. If inside\_vip/outside\_vip is not configured in the site object, system use
interface IP in the respected networks.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

Type: `"string"`. Computed.

Description of the AdvertisePolicy.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [dualstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/dualstack/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv6/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AdvertisePolicy.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Namespace where the AdvertisePolicy exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-port"></a>

### port property

Type: `"number"`. Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-port)
- [port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-port_ranges)

Select alternatives according to the provider validators above.

<a id="schema-port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="schema-protocol"></a>

### protocol property

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/): complete subsection reference.

<a id="schema-skip_xff_append"></a>

### skip_xff_append property

Type: `"bool"`. Computed.

If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.

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

- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/): complete subsection reference.

- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-address) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-description) |
| `dualstack` | [dualstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/dualstack/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-id) |
| `ipv4` | [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv4/#section) |
| `ipv6` | [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv6/#section) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-namespace) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-port) |
| `port_ranges` | [port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-port_ranges) |
| `protocol` | [protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-protocol) |
| `public_ip` | [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/#section) |
| `public_ip.kind` | [public_ip.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/#schema-public_ip--kind) |
| `public_ip.name` | [public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/#schema-public_ip--name) |
| `public_ip.namespace` | [public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/#schema-public_ip--namespace) |
| `public_ip.tenant` | [public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/#schema-public_ip--tenant) |
| `public_ip.uid` | [public_ip.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/public_ip/#schema-public_ip--uid) |
| `skip_xff_append` | [skip_xff_append](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/#schema-skip_xff_append) |
| `tls_parameters` | [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/#section) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/client_certificate_optional/#section) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/client_certificate_required/#section) |
| `tls_parameters.common_params` | [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/#section) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--cipher_suites) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--maximum_protocol_version) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--minimum_protocol_version) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/#section) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/#schema-tls_parameters--common_params--tls_certificates--certificate_url) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/#section) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/#schema-tls_parameters--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/#schema-tls_parameters--common_params--tls_certificates--description_spec) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/#section) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/#section) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/#section) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--skip_hostname_verification) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--trusted_ca_url) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--verify_subject_alt_names) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/no_client_certificate/#section) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/#schema-tls_parameters--xfcc_header_elements) |
| `where` | [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/#section) |
| `where.site` | [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/disable_internet_vip/#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/enable_internet_vip/#section) |
| `where.site.network_type` | [where.site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/ref/#section) |
| `where.site.ref.kind` | [where.site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/ref/#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/ref/#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/ref/#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/ref/#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/ref/#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/disable_internet_vip/#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/enable_internet_vip/#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/ref/#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--uid) |
