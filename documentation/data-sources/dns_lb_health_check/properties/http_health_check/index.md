---
page_title: "http_health_check"
subcategory: ""
description: "Configuration parameter for http health check."
xcsh_docs: {"aliases": ["http health check"], "body_bytes": 7097, "body_sha256": "sha256:d8f2f2d20e764fc5b0b37eb7f1c31c4907dd14212026283d1686c1d2c69b896a", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:reference", "path": "documentation/data-sources/dns_lb_health_check/properties/http_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012", "registry_path": "docs/guides/data-sources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check"], "schema_version": 1, "sections": [{"aliases": ["http health check disable virtual host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "disable_virtual_host"], "syntax": "attribute", "type": "object"}, {"aliases": ["http health check health check port"], "anchor": "schema-http_health_check--health_check_port", "description": "Port used for performing health check.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check health check secondary port"], "anchor": "schema-http_health_check--health_check_secondary_port", "description": "Secondary port used for performing health check. If included, both ports must be healthy for the health check to pass.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "health_check_secondary_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check inherit load balancer fqdn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "inherit_load_balancer_fqdn"], "syntax": "attribute", "type": "object"}, {"aliases": ["http health check receive", "succeeded", "success", "successful"], "anchor": "schema-http_health_check--receive", "description": "Regular expression used to match against the response to the health check's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "receive"], "syntax": "attribute", "type": "string"}, {"aliases": ["http health check send"], "anchor": "schema-http_health_check--send", "description": "HTTP payload to send to the target.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "send"], "syntax": "attribute", "type": "string"}, {"aliases": ["http health check virtual host"], "anchor": "schema-http_health_check--virtual_host", "description": "Exclusive with Name of the virtual host to use for SNI.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "virtual_host"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/properties/http_health_check/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration parameter for http health check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/)
- http_health_check

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: http\_health\_check, https\_health\_check, icmp\_health\_check, tcp\_health\_check,
tcp\_hex\_health\_check, udp\_health\_check\] Configuration parameter for http health check.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/#section)
- [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/https_health_check/#section)
- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/icmp_health_check/#section)
- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_health_check/#section)
- [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/#section)
- [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/udp_health_check/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/disable_virtual_host/): complete subsection reference.

<a id="schema-http_health_check--health_check_port"></a>

### health_check_port property

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-http_health_check--health_check_secondary_port"></a>

### health_check_secondary_port property

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/): complete subsection reference.

<a id="schema-http_health_check--receive"></a>

### receive property

Type: `"string"`. Computed.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-http_health_check--send"></a>

### send property

Type: `"string"`. Computed.

Send String. HTTP payload to send to the target.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="schema-http_health_check--virtual_host"></a>

### virtual_host property

Type: `"string"`. Computed.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```
