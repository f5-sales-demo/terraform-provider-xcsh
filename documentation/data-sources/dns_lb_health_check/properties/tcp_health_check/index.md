---
page_title: "tcp_health_check"
subcategory: ""
description: "Configuration parameter for tcp health check."
xcsh_docs: {"aliases": ["tcp health check"], "body_bytes": 4936, "body_sha256": "sha256:c30215db812738e22273c7ea3ef017345e851f28ad3831ef2616711b262804e4", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_health_check", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:reference", "path": "documentation/data-sources/dns_lb_health_check/properties/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3312322330200030-1232321021303111-3101103331113211-2113200131132022-1122310212123331-1332031232103330-0012301023331021-3103320231031323", "registry_path": "docs/guides/data-sources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["health check port"], "anchor": "schema-tcp_health_check--health_check_port", "description": "Port used for performing health check.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["health check secondary port"], "anchor": "schema-tcp_health_check--health_check_secondary_port", "description": "Secondary port used for performing health check. If included, both ports must be healthy for the health check to pass.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "health_check_secondary_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["receive", "succeeded", "success", "successful"], "anchor": "schema-tcp_health_check--receive", "description": "Regular expression used to match against the response to the monitor's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "receive"], "syntax": "attribute", "type": "string"}, {"aliases": ["send"], "anchor": "schema-tcp_health_check--send", "description": "Send this string to target (default empty. When send and receive are both empty, monitor just tests 3WHS)", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "send"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/properties/tcp_health_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for tcp health check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tcp_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/)
- tcp_health_check

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp health check.

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

## Direct properties

<a id="schema-tcp_health_check--health_check_port"></a>

### health_check_port property

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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

<a id="schema-tcp_health_check--health_check_secondary_port"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-tcp_health_check--receive"></a>

### receive property

Type: `"string"`. Computed.

Regular expression used to match against the response to the monitor's request. Mark node up upon
receipt of a successful regular expression match. Uses re2 regular expression syntax.

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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-tcp_health_check--send"></a>

### send property

Type: `"string"`. Computed.

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS).

Upstream description:

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS)

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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/)
- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
