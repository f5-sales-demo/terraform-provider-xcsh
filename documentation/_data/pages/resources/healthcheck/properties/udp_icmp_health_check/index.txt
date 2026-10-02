---
page_title: "udp_icmp_health_check"
subcategory: "Monitoring"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["udp icmp health check"], "body_bytes": 1132, "body_sha256": "sha256:ed65cd17e0da879ea604fd91284b0607bf88465c0c8d74a3a027770eacd526fb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:udp_icmp_health_check", "parent_id": "xcsh-docs:resources:healthcheck:reference", "path": "documentation/resources/healthcheck/properties/udp_icmp_health_check/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2113132323233313-0012332121130131-0221300000220210-3133200021330003-1022130321212233-1211101310120213-3323202003303021-0302001212112303", "registry_path": "docs/guides/resources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["udp_icmp_health_check"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/udp_icmp_health_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["healthcheckCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# udp_icmp_health_check

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- udp_icmp_health_check

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for udp icmp health check.

Upstream description:

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
udp_icmp_health_check = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
