---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": ["dns lb health check"], "body_bytes": 17831, "body_sha256": "sha256:1449619fa671b35cffcf3f29b85f2da2306e71ccbf41c91b085064aa41facb34", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "xcsh-docs:resources:dns_lb_health_check:properties:icmp_health_check", "xcsh-docs:resources:dns_lb_health_check:properties:tcp_health_check", "xcsh-docs:resources:dns_lb_health_check:properties:tcp_hex_health_check", "xcsh-docs:resources:dns_lb_health_check:properties:timeouts", "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:reference", "parent_id": "xcsh-docs:resources:dns_lb_health_check:fundamentals", "path": "documentation/resources/dns_lb_health_check/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100", "registry_path": "docs/guides/resources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["http health check"], "anchor": "section", "description": "Configuration parameter for http health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_health_check--virtual_host", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "type": "conflicts"}, {"anchor": "schema-http_health_check--virtual_host", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:inherit_load_balancer_fqdn,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,inherit_load_balancer_fqdn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,inherit_load_balancer_fqdn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:inherit_load_balancer_fqdn,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "type": "conflicts"}, {"anchor": "schema-http_health_check--health_check_port", "enforcement": "provider-schema", "group": "http_health_check:RequiredObjectAttributes:health_check_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "type": "requires"}], "schema_path": ["http_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["https health check"], "anchor": "section", "description": "Configuration parameter for https health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-https_health_check--virtual_host", "enforcement": "provider-schema", "group": "https_health_check:ConflictingObjectAttributes:disable_virtual_host,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "type": "conflicts"}, {"anchor": "schema-https_health_check--virtual_host", "enforcement": "provider-schema", "group": "https_health_check:ConflictingObjectAttributes:inherit_load_balancer_fqdn,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_health_check:ConflictingObjectAttributes:disable_virtual_host,inherit_load_balancer_fqdn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:disable_virtual_host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_health_check:ConflictingObjectAttributes:disable_virtual_host,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:disable_virtual_host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_health_check:ConflictingObjectAttributes:disable_virtual_host,inherit_load_balancer_fqdn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:inherit_load_balancer_fqdn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_health_check:ConflictingObjectAttributes:inherit_load_balancer_fqdn,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:inherit_load_balancer_fqdn", "type": "conflicts"}, {"anchor": "schema-https_health_check--health_check_port", "enforcement": "provider-schema", "group": "https_health_check:RequiredObjectAttributes:health_check_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "type": "requires"}], "schema_path": ["https_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["icmp health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:icmp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["icmp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:dns_lb_health_check:reference", "enum_extraction_complete": false, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["system"], "version": 1}], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tcp health check"], "anchor": "section", "description": "Configuration parameter for tcp health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:tcp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tcp_health_check--health_check_port", "enforcement": "provider-schema", "group": "tcp_health_check:RequiredObjectAttributes:health_check_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:tcp_health_check", "type": "requires"}], "schema_path": ["tcp_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["tcp hex health check"], "anchor": "section", "description": "Configuration parameter for tcp hex health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:tcp_hex_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tcp_hex_health_check--health_check_port", "enforcement": "provider-schema", "group": "tcp_hex_health_check:RequiredObjectAttributes:health_check_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:tcp_hex_health_check", "type": "requires"}], "schema_path": ["tcp_hex_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["udp health check"], "anchor": "section", "description": "Configuration parameter for udp health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-udp_health_check--health_check_port", "enforcement": "provider-schema", "group": "udp_health_check:RequiredObjectAttributes:health_check_port,receive,send", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "type": "requires"}, {"anchor": "schema-udp_health_check--receive", "enforcement": "provider-schema", "group": "udp_health_check:RequiredObjectAttributes:health_check_port,receive,send", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "type": "requires"}, {"anchor": "schema-udp_health_check--send", "enforcement": "provider-schema", "group": "udp_health_check:RequiredObjectAttributes:health_check_port,receive,send", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "type": "requires"}], "schema_path": ["udp_health_check"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_dns_lb_health_check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Additional upstream details:

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/): complete subsection reference.

- [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/): complete subsection reference.

- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/icmp_health_check/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the DNS LB Health Check. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Namespace for the DNS LB Health Check. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/): complete subsection reference.

- [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/timeouts/): complete subsection reference.

- [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-disable) |
| `http_health_check` | [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#section) |
| `http_health_check.disable_virtual_host` | [http_health_check.disable_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/disable_virtual_host/#section) |
| `http_health_check.health_check_port` | [http_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--health_check_port) |
| `http_health_check.health_check_secondary_port` | [http_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--health_check_secondary_port) |
| `http_health_check.inherit_load_balancer_fqdn` | [http_health_check.inherit_load_balancer_fqdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/#section) |
| `http_health_check.receive` | [http_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--receive) |
| `http_health_check.send` | [http_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--send) |
| `http_health_check.virtual_host` | [http_health_check.virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#schema-http_health_check--virtual_host) |
| `https_health_check` | [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#section) |
| `https_health_check.disable_virtual_host` | [https_health_check.disable_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/disable_virtual_host/#section) |
| `https_health_check.health_check_port` | [https_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--health_check_port) |
| `https_health_check.health_check_secondary_port` | [https_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--health_check_secondary_port) |
| `https_health_check.inherit_load_balancer_fqdn` | [https_health_check.inherit_load_balancer_fqdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/inherit_load_balancer_fqdn/#section) |
| `https_health_check.receive` | [https_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--receive) |
| `https_health_check.send` | [https_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--send) |
| `https_health_check.virtual_host` | [https_health_check.virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#schema-https_health_check--virtual_host) |
| `icmp_health_check` | [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/icmp_health_check/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/#schema-namespace) |
| `tcp_health_check` | [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/#section) |
| `tcp_health_check.health_check_port` | [tcp_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--health_check_port) |
| `tcp_health_check.health_check_secondary_port` | [tcp_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--health_check_secondary_port) |
| `tcp_health_check.receive` | [tcp_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--receive) |
| `tcp_health_check.send` | [tcp_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/#schema-tcp_health_check--send) |
| `tcp_hex_health_check` | [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/#section) |
| `tcp_hex_health_check.health_check_port` | [tcp_hex_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--health_check_port) |
| `tcp_hex_health_check.health_check_secondary_port` | [tcp_hex_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--health_check_secondary_port) |
| `tcp_hex_health_check.receive` | [tcp_hex_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--receive) |
| `tcp_hex_health_check.send` | [tcp_hex_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/#schema-tcp_hex_health_check--send) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/timeouts/#schema-timeouts--update) |
| `udp_health_check` | [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/#section) |
| `udp_health_check.health_check_port` | [udp_health_check.health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--health_check_port) |
| `udp_health_check.health_check_secondary_port` | [udp_health_check.health_check_secondary_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--health_check_secondary_port) |
| `udp_health_check.receive` | [udp_health_check.receive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--receive) |
| `udp_health_check.send` | [udp_health_check.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/#schema-udp_health_check--send) |
