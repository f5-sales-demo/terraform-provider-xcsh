---
page_title: "Property reference"
subcategory: "Networking"
description: "Property reference for xcsh_endpoint."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 22624, "body_sha256": "sha256:54b03cd0644f5770d2847e08ca58b576695b7ef31825654e7c0d6722af377b6b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:endpoint:properties:dns_name_advanced", "xcsh-docs:data-sources:endpoint:properties:service_info", "xcsh-docs:data-sources:endpoint:properties:snat_pool", "xcsh-docs:data-sources:endpoint:properties:where"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:reference", "parent_id": "xcsh-docs:data-sources:endpoint:fundamentals", "path": "documentation/data-sources/endpoint/properties/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223", "registry_path": "docs/guides/data-sources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns name"], "anchor": "schema-dns_name", "description": "Exclusive with Endpoint's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns name advanced"], "anchor": "section", "description": "Specifies name and TTL used for DNS resolution.", "document_id": "xcsh-docs:data-sources:endpoint:properties:dns_name_advanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_name_advanced"], "syntax": "attribute", "type": "object"}, {"aliases": ["health check port"], "anchor": "schema-health_check_port", "description": "By default the health check port of an endpoint is the same as the endpoint’s port. This option provides an alternative health check port. Setting this with a non-zero value allows an endpoint to have different health check port.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip"], "anchor": "schema-ip", "description": "Exclusive with Endpoint is reachable at the given IPv4/IPv6 address.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-port", "description": "Endpoint service is available on this port.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port"], "syntax": "attribute", "type": "number"}, {"aliases": ["protocol"], "anchor": "schema-protocol", "description": "Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.", "document_id": "xcsh-docs:data-sources:endpoint:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["service info"], "anchor": "section", "description": "Specifies whether endpoint service is discovered by name or labels.", "document_id": "xcsh-docs:data-sources:endpoint:properties:service_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:data-sources:endpoint:properties:snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["where"], "anchor": "section", "description": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites selec", "document_id": "xcsh-docs:data-sources:endpoint:properties:where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

Type: `"string"`. Computed.

Description of the Endpoint.

Upstream description:

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

<a id="schema-dns_name"></a>

### dns_name property

Type: `"string"`. Computed.

\[OneOf: dns\_name, dns\_name\_advanced, ip, service\_info\] Exclusive with \[dns\_name\_advanced IP
service\_info\] Endpoint's IP address is discovered using DNS name resolution. The name given here
is fully qualified domain name.

Upstream description:

Exclusive with \[dns\_name\_advanced IP service\_info\] Endpoint's IP address is discovered using
DNS name resolution. The name given here is fully qualified domain name.

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

- [dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-dns_name)
- [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/dns_name_advanced/#section)
- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-ip)
- [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/#section)

Select alternatives according to the provider validators above.

- [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/dns_name_advanced/): complete subsection reference.

<a id="schema-health_check_port"></a>

### health_check_port property

Type: `"number"`. Computed.

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Upstream description:

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

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

Type: `"string"`. Computed.

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Upstream description:

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

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

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

Name of the Endpoint.

Upstream description:

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

Namespace where the Endpoint exists.

Upstream description:

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

<a id="schema-port"></a>

### port property

Type: `"number"`. Computed.

Endpoint service is available on this port.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Endpoint protocol. Default is TCP. Both TCP and UDP protocols are
supported. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.

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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/): complete subsection reference.

- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-description) |
| `dns_name` | [dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-dns_name) |
| `dns_name_advanced` | [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/dns_name_advanced/#section) |
| `dns_name_advanced.name` | [dns_name_advanced.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/dns_name_advanced/#schema-dns_name_advanced--name) |
| `dns_name_advanced.refresh_interval` | [dns_name_advanced.refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/dns_name_advanced/#schema-dns_name_advanced--refresh_interval) |
| `health_check_port` | [health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-health_check_port) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-id) |
| `ip` | [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-ip) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-namespace) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-port) |
| `protocol` | [protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/#schema-protocol) |
| `service_info` | [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/#section) |
| `service_info.discovery_type` | [service_info.discovery_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/#schema-service_info--discovery_type) |
| `service_info.service_name` | [service_info.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/#schema-service_info--service_name) |
| `service_info.service_selector` | [service_info.service_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/service_selector/#section) |
| `service_info.service_selector.expressions` | [service_info.service_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/service_selector/#schema-service_info--service_selector--expressions) |
| `snat_pool` | [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/#section) |
| `snat_pool.no_snat_pool` | [snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/no_snat_pool/#section) |
| `snat_pool.snat_pool` | [snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/snat_pool/#section) |
| `snat_pool.snat_pool.prefixes` | [snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/snat_pool/#schema-snat_pool--snat_pool--prefixes) |
| `where` | [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/#section) |
| `where.site` | [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/disable_internet_vip/#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/enable_internet_vip/#section) |
| `where.site.network_type` | [where.site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/ref/#section) |
| `where.site.ref.kind` | [where.site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/ref/#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/ref/#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/ref/#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/ref/#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/site/ref/#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/ref/#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/disable_internet_vip/#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/enable_internet_vip/#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/ref/#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--uid) |

## Next pages

- [dns_name_advanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/dns_name_advanced/)
- [service_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/)
- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/where/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
