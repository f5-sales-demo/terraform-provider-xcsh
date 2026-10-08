---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_dns_proxy."
xcsh_docs: {"aliases": ["dns proxy"], "body_bytes": 352, "body_sha256": "sha256:470d0c3d6d25110416c8c578e7bb68df4375b5c6f36b9e561b58f86b8e6aa6e9", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:dns_proxy:fundamentals", "path": "documentation/resources/dns_proxy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3303233331131231-1203112221022131-1013330320213321-3003331120010332-2132111030023201-0001023321302302-0313200222313111-0320223111222131", "registry_path": "docs/guides/resources--dns_proxy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Import for xcsh_dns_proxy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_dns_proxy.example system/example
```
