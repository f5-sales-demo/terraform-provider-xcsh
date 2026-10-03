---
page_title: "origin_servers.health_checks.health_check"
subcategory: ""
description: "List of Health Checks."
xcsh_docs: {"aliases": ["origin servers health checks health check"], "body_bytes": 3817, "body_sha256": "sha256:c5d5f261cea869544729d7423b24d66adfae8172edf72f72597f0ec8957f32d8", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "path": "documentation/resources/dns_proxy/properties/origin_servers/health_checks/health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0011102230222312-1022032120311221-3030020231301220-1033133112220102-1212330123002211-3011313200110103-2111001133202211-3000032121011230", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check:ConflictingListObjectAttributes:dns_health_check,icmp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check:ConflictingListObjectAttributes:dns_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check:ConflictingListObjectAttributes:dns_health_check,icmp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check:ConflictingListObjectAttributes:icmp_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check:ConflictingListObjectAttributes:dns_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check:ConflictingListObjectAttributes:icmp_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check"], "schema_version": 1, "sections": [{"aliases": ["origin servers health checks health check dns health check", "succeeded", "success", "successful"], "anchor": "section", "description": "DNS health check reports healthy if DNS query is successful and response header and answer matches the given value.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_response", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.dns_health_check:RequiredObjectAttributes:expected_response,query_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--query_name", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.dns_health_check:RequiredObjectAttributes:expected_response,query_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "type": "requires"}], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers health checks health check icmp health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "icmp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers health checks health check tcp health check", "succeeded", "success", "successful"], "anchor": "section", "description": "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.tcp_health_check:RequiredObjectAttributes:expected_response,send_payload", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.tcp_health_check:RequiredObjectAttributes:expected_response,send_payload", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "type": "requires"}], "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/health_checks/health_check/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "List of Health Checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/)
- origin_servers.health_checks.health_check

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Health Checks. List of Health Checks.

Upstream description:

List of Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns_health_check",
    "icmp_health_check"),
  validators.ConflictingListObjectAttributes("dns_health_check",
    "tcp_health_check"),
  validators.ConflictingListObjectAttributes("icmp_health_check",
    "tcp_health_check")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dns_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/): complete subsection reference.

- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/icmp_health_check/): complete subsection reference.

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/): complete subsection reference.

## Next pages

- [origin_servers.health_checks.health_check.dns_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/)
- [origin_servers.health_checks.health_check.icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/icmp_health_check/)
- [origin_servers.health_checks.health_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
