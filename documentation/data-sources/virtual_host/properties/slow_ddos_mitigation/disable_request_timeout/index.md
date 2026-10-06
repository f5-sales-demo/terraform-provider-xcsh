---
page_title: "slow_ddos_mitigation.disable_request_timeout"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["duration", "slow ddos mitigation disable request timeout"], "body_bytes": 1020, "body_sha256": "sha256:c481421a041250b249c00d8700bd1b2ed4a5916f1a2ccef261f90320c779104f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:slow_ddos_mitigation:disable_request_timeout", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:slow_ddos_mitigation", "path": "documentation/data-sources/virtual_host/properties/slow_ddos_mitigation/disable_request_timeout/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2332311122002120-0102020211112123-3101022133022100-2132210101231300-2123013103112213-2123200322222131-3202202312030120-1203310031330002", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/slow_ddos_mitigation/disable_request_timeout/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation.disable_request_timeout

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/slow_ddos_mitigation/)
- slow_ddos_mitigation.disable_request_timeout

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

This is an empty object or choice marker. It has no direct properties.
