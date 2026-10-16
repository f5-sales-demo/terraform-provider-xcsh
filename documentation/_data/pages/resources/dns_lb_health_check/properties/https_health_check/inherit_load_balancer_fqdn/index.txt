---
page_title: "https_health_check.inherit_load_balancer_fqdn"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https health check inherit load balancer fqdn"], "body_bytes": 1109, "body_sha256": "sha256:f8caae3070d72e29ad422308ee83446d48b208cf4bd4551d6be6586af3bcfea8", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:inherit_load_balancer_fqdn", "parent_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "path": "documentation/resources/dns_lb_health_check/properties/https_health_check/inherit_load_balancer_fqdn/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1012102332121320-2033103023300022-1201003031022222-2330312202111103-3232232222321322-3100330300321323-0030330132132203-2333001313300111", "registry_path": "docs/guides/resources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_health_check", "inherit_load_balancer_fqdn"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/https_health_check/inherit_load_balancer_fqdn/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_health_check.inherit_load_balancer_fqdn

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- [https_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/https_health_check/)
- https_health_check.inherit_load_balancer_fqdn

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit load balancer fqdn.

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
inherit_load_balancer_fqdn = {}
```

This is an empty object or choice marker. It has no direct properties.
