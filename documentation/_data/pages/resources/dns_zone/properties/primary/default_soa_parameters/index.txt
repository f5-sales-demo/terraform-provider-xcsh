---
page_title: "primary.default_soa_parameters"
subcategory: "DNS"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["primary default soa parameters"], "body_bytes": 1005, "body_sha256": "sha256:a74007449086b6390d3eccf68bdc5026ffea714ed41a2d68fcc6dd31bb63fda9", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "documentation/resources/dns_zone/properties/primary/default_soa_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0001013031002232-2323220130321002-0110031123323021-1302303333101221-0212302022021323-3331121010231120-3131130311320003-3312131103003301", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_soa_parameters"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_soa_parameters/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_soa_parameters

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- primary.default_soa_parameters

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default soa parameters.

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
default_soa_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.
