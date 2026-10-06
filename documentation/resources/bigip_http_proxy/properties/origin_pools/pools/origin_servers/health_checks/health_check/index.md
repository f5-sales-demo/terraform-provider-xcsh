---
page_title: "origin_pools.pools.origin_servers.health_checks.health_check"
subcategory: ""
description: "List of Health Checks."
xcsh_docs: {"aliases": ["origin pools pools origin servers health checks health check"], "body_bytes": 2953, "body_sha256": "sha256:7d5f347413710b4b51448023809f35ecd96f7654ef7aaf868c14ab7007734875", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:icmp_health_check", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2120333231121323-2003031320012101-2302210211013333-0122323100321303-1310022230223303-2202022123002231-0033332311313302-2323130331102303", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks.health_check:ConflictingListObjectAttributes:icmp_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:icmp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks.health_check:ConflictingListObjectAttributes:icmp_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check"], "schema_version": 1, "sections": [{"aliases": ["origin pools pools origin servers health checks health check icmp health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:icmp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check", "icmp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin pools pools origin servers health checks health check tcp health check", "succeeded", "success", "successful"], "anchor": "section", "description": "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--expected_response", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check:RequiredObjectAttributes:expected_response,send_payload", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "type": "requires"}, {"anchor": "schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--send_payload", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check:RequiredObjectAttributes:expected_response,send_payload", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "type": "requires"}], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check", "tcp_health_check"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of Health Checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.health_checks.health_check

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/)
- origin_pools.pools.origin_servers.health_checks.health_check

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Health Checks. List of Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("icmp_health_check",
    "tcp_health_check")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
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

- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/icmp_health_check/): complete subsection reference.

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/tcp_health_check/): complete subsection reference.
