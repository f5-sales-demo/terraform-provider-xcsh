---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": [], "body_bytes": 16130, "body_sha256": "sha256:3c7c7459dfb345551b311259c8e4addb8deec4878b48c6ee7d7fc57beb055c98", "child_ids": ["xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "xcsh-docs:data-sources:dns_lb_health_check:properties:https_health_check", "xcsh-docs:data-sources:dns_lb_health_check:properties:icmp_health_check", "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_health_check", "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_hex_health_check", "xcsh-docs:data-sources:dns_lb_health_check:properties:udp_health_check"], "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_health_check:reference", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:fundamentals", "path": "documentation/data-sources/dns_lb_health_check/properties/index.md", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_lb_health_check.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
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

Description of the DNSLBHealthCheck.

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/): complete subsection reference.

- [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/): complete subsection reference.

- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/icmp_health_check/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

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

Name of the DNSLBHealthCheck.

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

Type: `"string"`. Optional, Computed.

Namespace where the DNSLBHealthCheck exists.

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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/): complete subsection reference.

- [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/): complete subsection reference.

- [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/#schema-description) |
| `http_health_check` | [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#section) |
| `http_health_check.disable_virtual_host` | [http_health_check.disable_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/disable_virtual_host/#section) |
| `http_health_check.health_check_port` | [http_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--health_check_port) |
| `http_health_check.health_check_secondary_port` | [http_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--health_check_secondary_port) |
| `http_health_check.inherit_load_balancer_fqdn` | [http_health_check.inherit_load_balancer_fqdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/#section) |
| `http_health_check.receive` | [http_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--receive) |
| `http_health_check.send` | [http_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--send) |
| `http_health_check.virtual_host` | [http_health_check.virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--virtual_host) |
| `https_health_check` | [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#section) |
| `https_health_check.disable_virtual_host` | [https_health_check.disable_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/disable_virtual_host/#section) |
| `https_health_check.health_check_port` | [https_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--health_check_port) |
| `https_health_check.health_check_secondary_port` | [https_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--health_check_secondary_port) |
| `https_health_check.inherit_load_balancer_fqdn` | [https_health_check.inherit_load_balancer_fqdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/inherit_load_balancer_fqdn/#section) |
| `https_health_check.receive` | [https_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--receive) |
| `https_health_check.send` | [https_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--send) |
| `https_health_check.virtual_host` | [https_health_check.virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--virtual_host) |
| `icmp_health_check` | [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/icmp_health_check/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/#schema-namespace) |
| `tcp_health_check` | [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/#section) |
| `tcp_health_check.health_check_port` | [tcp_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--health_check_port) |
| `tcp_health_check.health_check_secondary_port` | [tcp_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--health_check_secondary_port) |
| `tcp_health_check.receive` | [tcp_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--receive) |
| `tcp_health_check.send` | [tcp_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--send) |
| `tcp_hex_health_check` | [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/#section) |
| `tcp_hex_health_check.health_check_port` | [tcp_hex_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--health_check_port) |
| `tcp_hex_health_check.health_check_secondary_port` | [tcp_hex_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--health_check_secondary_port) |
| `tcp_hex_health_check.receive` | [tcp_hex_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--receive) |
| `tcp_hex_health_check.send` | [tcp_hex_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--send) |
| `udp_health_check` | [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/#section) |
| `udp_health_check.health_check_port` | [udp_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--health_check_port) |
| `udp_health_check.health_check_secondary_port` | [udp_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--health_check_secondary_port) |
| `udp_health_check.receive` | [udp_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--receive) |
| `udp_health_check.send` | [udp_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--send) |

## Next pages

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/)
- [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/)
- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/icmp_health_check/)
- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/)
- [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/)
- [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/)
- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
