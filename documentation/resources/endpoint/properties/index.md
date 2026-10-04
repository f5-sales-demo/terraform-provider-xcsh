---
page_title: "Property reference"
subcategory: "Networking"
description: "Property reference for xcsh_endpoint."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 25750, "body_sha256": "sha256:27778f40c2b141963c7521a8c107d3c62919fe05b3851210738433eeaceb455a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:endpoint:properties:dns_name_advanced", "xcsh-docs:resources:endpoint:properties:service_info", "xcsh-docs:resources:endpoint:properties:snat_pool", "xcsh-docs:resources:endpoint:properties:timeouts", "xcsh-docs:resources:endpoint:properties:where"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:reference", "parent_id": "xcsh-docs:resources:endpoint:fundamentals", "path": "documentation/resources/endpoint/properties/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["dns name"], "anchor": "schema-dns_name", "description": "Exclusive with Endpoint's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns name advanced"], "anchor": "section", "description": "Specifies name and TTL used for DNS resolution.", "document_id": "xcsh-docs:resources:endpoint:properties:dns_name_advanced", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_name_advanced"], "syntax": "block", "type": "object"}, {"aliases": ["health check port"], "anchor": "schema-health_check_port", "description": "By default the health check port of an endpoint is the same as the endpoint’s port. This option provides an alternative health check port. Setting this with a non-zero value allows an endpoint to have different health check port.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip"], "anchor": "schema-ip", "description": "Exclusive with Endpoint is reachable at the given IPv4/IPv6 address.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-port", "description": "Endpoint service is available on this port.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port"], "syntax": "attribute", "type": "number"}, {"aliases": ["protocol"], "anchor": "schema-protocol", "description": "Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.", "document_id": "xcsh-docs:resources:endpoint:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["service info"], "anchor": "section", "description": "Specifies whether endpoint service is discovered by name or labels.", "document_id": "xcsh-docs:resources:endpoint:properties:service_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service_info--service_name", "enforcement": "provider-schema", "group": "service_info:ConflictingObjectAttributes:service_name,service_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:service_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service_info:ConflictingObjectAttributes:service_name,service_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:service_info:service_selector", "type": "conflicts"}], "schema_path": ["service_info"], "syntax": "block", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:resources:endpoint:properties:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool", "type": "conflicts"}], "schema_path": ["snat_pool"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:endpoint:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["where"], "anchor": "section", "description": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "document_id": "xcsh-docs:resources:endpoint:properties:where", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:where:virtual_site", "type": "conflicts"}], "schema_path": ["where"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-dns_name"></a>

### dns_name property

Type: `"string"`. Optional, Computed.

\[OneOf: dns\_name, dns\_name\_advanced, ip, service\_info\] Exclusive with \[dns\_name\_advanced IP
service\_info\] Endpoint's IP address is discovered using DNS name resolution. The name given here
is fully qualified domain name.

Upstream description:

Exclusive with \[dns\_name\_advanced IP service\_info\] Endpoint's IP address is discovered using
DNS name resolution. The name given here is fully qualified domain name.

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
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-dns_name)
- [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/dns_name_advanced/#section)
- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-ip)
- [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/#section)

Select alternatives according to the provider validators above.

- [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/dns_name_advanced/): complete subsection reference.

<a id="schema-health_check_port"></a>

### health_check_port property

Type: `"number"`. Optional, Computed.

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Upstream description:

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(65535),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-ip"></a>

### ip property

Type: `"string"`. Optional, Computed.

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Upstream description:

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Endpoint. Must be unique within the namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Namespace where the Endpoint is created.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"number"`. Optional, Computed.

Endpoint service is available on this port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-protocol"></a>

### protocol property

Type: `"string"`. Optional, Computed.

\[Enum: TCP|UDP\] Protocol. Endpoint protocol. Default is TCP. Both TCP and UDP protocols are
supported. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TCP",
    "UDP"),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/): complete subsection reference.

- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-disable) |
| `dns_name` | [dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-dns_name) |
| `dns_name_advanced` | [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/dns_name_advanced/#section) |
| `dns_name_advanced.name` | [dns_name_advanced.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/dns_name_advanced/#schema-dns_name_advanced--name) |
| `dns_name_advanced.refresh_interval` | [dns_name_advanced.refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/dns_name_advanced/#schema-dns_name_advanced--refresh_interval) |
| `health_check_port` | [health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-health_check_port) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-id) |
| `ip` | [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-ip) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-namespace) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-port) |
| `protocol` | [protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/#schema-protocol) |
| `service_info` | [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/#section) |
| `service_info.discovery_type` | [service_info.discovery_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/#schema-service_info--discovery_type) |
| `service_info.service_name` | [service_info.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/#schema-service_info--service_name) |
| `service_info.service_selector` | [service_info.service_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/service_selector/#section) |
| `service_info.service_selector.expressions` | [service_info.service_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/service_selector/#schema-service_info--service_selector--expressions) |
| `snat_pool` | [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/#section) |
| `snat_pool.no_snat_pool` | [snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/no_snat_pool/#section) |
| `snat_pool.snat_pool` | [snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/snat_pool/#section) |
| `snat_pool.snat_pool.prefixes` | [snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/snat_pool/#schema-snat_pool--snat_pool--prefixes) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/#schema-timeouts--update) |
| `where` | [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/#section) |
| `where.site` | [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/disable_internet_vip/#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/enable_internet_vip/#section) |
| `where.site.network_type` | [where.site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/ref/#section) |
| `where.site.ref.kind` | [where.site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/ref/#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/ref/#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/ref/#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/ref/#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/site/ref/#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/disable_internet_vip/#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/enable_internet_vip/#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/ref/#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--uid) |

## Next pages

- [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/dns_name_advanced/)
- [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/)
- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/timeouts/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
