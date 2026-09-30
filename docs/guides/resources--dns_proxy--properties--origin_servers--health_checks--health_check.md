---
page_title: "origin_servers.health_checks.health_check"
subcategory: ""
description: "origin_servers.health_checks.health_check for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3124, "body_sha256": "sha256:dddbb20db9609b0ab789ee8c6ef54e70d67bb18acce1c65e43afd756349134a5", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--health_checks--health_check.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/health_checks/health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.health_checks.health_check for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.health_checks.health_check

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [origin_servers.health_checks](resources--dns_proxy--properties--origin_servers--health_checks.md)
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [dns_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md): complete subsection reference.

- [icmp_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--icmp_health_check.md): complete subsection reference.

- [tcp_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md): complete subsection reference.

## Next pages

- [origin_servers.health_checks.health_check.dns_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md)
- [origin_servers.health_checks.health_check.icmp_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--icmp_health_check.md)
- [origin_servers.health_checks.health_check.tcp_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md)
- [origin_servers.health_checks](resources--dns_proxy--properties--origin_servers--health_checks.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
