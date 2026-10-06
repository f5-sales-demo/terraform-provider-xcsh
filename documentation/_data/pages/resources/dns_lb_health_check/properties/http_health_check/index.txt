---
page_title: "http_health_check"
subcategory: ""
description: "Configuration parameter for http health check."
xcsh_docs: {"aliases": ["http health check"], "body_bytes": 8316, "body_sha256": "sha256:e636dda5f5d2172c724bab3db4c65880b21b929a42d4adb2f101049ee813da70", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "parent_id": "xcsh-docs:resources:dns_lb_health_check:reference", "path": "documentation/resources/dns_lb_health_check/properties/http_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3123301133033020-3001312112033103-2023010020133212-1212321113313010-2303023100323100-0031323300231202-0032201133132201-2022313223021312", "registry_path": "docs/guides/resources--dns_lb_health_check--reference--group-001.md", "relationships": [{"anchor": "schema-http_health_check--virtual_host", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "type": "conflicts"}, {"anchor": "schema-http_health_check--virtual_host", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:inherit_load_balancer_fqdn,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,inherit_load_balancer_fqdn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:disable_virtual_host,inherit_load_balancer_fqdn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:inherit_load_balancer_fqdn,virtual_host", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "type": "conflicts"}, {"anchor": "schema-http_health_check--health_check_port", "enforcement": "provider-schema", "group": "http_health_check:RequiredObjectAttributes:health_check_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check"], "schema_version": 1, "sections": [{"aliases": ["http health check disable virtual host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "disable_virtual_host"], "syntax": "attribute", "type": "object"}, {"aliases": ["http health check health check port"], "anchor": "schema-http_health_check--health_check_port", "description": "Port used for performing health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check health check secondary port"], "anchor": "schema-http_health_check--health_check_secondary_port", "description": "Secondary port used for performing health check. If included, both ports must be healthy for the health check to pass.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "health_check_secondary_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check inherit load balancer fqdn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:inherit_load_balancer_fqdn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "inherit_load_balancer_fqdn"], "syntax": "attribute", "type": "object"}, {"aliases": ["http health check receive", "succeeded", "success", "successful"], "anchor": "schema-http_health_check--receive", "description": "Regular expression used to match against the response to the health check's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "receive"], "syntax": "attribute", "type": "string"}, {"aliases": ["http health check send"], "anchor": "schema-http_health_check--send", "description": "HTTP payload to send to the target.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "send"], "syntax": "attribute", "type": "string"}, {"aliases": ["http health check virtual host"], "anchor": "schema-http_health_check--virtual_host", "description": "Exclusive with Name of the virtual host to use for SNI.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_health_check", "virtual_host"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/http_health_check/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for http health check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- http_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http\_health\_check, https\_health\_check, icmp\_health\_check, tcp\_health\_check,
tcp\_hex\_health\_check, udp\_health\_check\] Configuration parameter for http health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "inherit_load_balancer_fqdn"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "virtual_host"),
  validators.ConflictingObjectAttributes("inherit_load_balancer_fqdn",
    "virtual_host")}
```

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/#section)
- [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/#section)
- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/icmp_health_check/#section)
- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_health_check/#section)
- [tcp_hex_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/tcp_hex_health_check/#section)
- [udp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/udp_health_check/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/disable_virtual_host/): complete subsection reference.

<a id="schema-http_health_check--health_check_port"></a>

### health_check_port property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

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

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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

- [inherit_load_balancer_fqdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/inherit_load_balancer_fqdn/): complete subsection reference.

<a id="schema-http_health_check--receive"></a>

### receive property

Type: `"string"`. Optional.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

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

Type: `"string"`. Optional.

Send String. HTTP payload to send to the target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="schema-http_health_check--virtual_host"></a>

### virtual_host property

Type: `"string"`. Optional.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```
