---
page_title: "icmp_health_check"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["icmp health check"], "body_bytes": 894, "body_sha256": "sha256:ddd1ed20f26847a7ded06416a136ccac11464c3781a06445026a9e5589cca2dc", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:icmp_health_check", "parent_id": "xcsh-docs:resources:dns_lb_health_check:reference", "path": "documentation/resources/dns_lb_health_check/properties/icmp_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0022113030000323-2330222301031023-0120100012200301-2333110032132213-0003230321003203-2023312222130303-1322131221322010-0133303003303103", "registry_path": "docs/guides/resources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["icmp_health_check"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/icmp_health_check/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# icmp_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- icmp_health_check

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
icmp_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.
