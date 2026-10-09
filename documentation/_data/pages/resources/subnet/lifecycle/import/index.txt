---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_subnet."
xcsh_docs: {"aliases": ["subnet"], "body_bytes": 343, "body_sha256": "sha256:6eca1c99357ff547c89e0fdcf64a1157559972448be63cd7e364a187db7e144e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:subnet:fundamentals", "path": "documentation/resources/subnet/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2133302013023312-0313221302123301-1313011330200313-3112332131333123-1023233100311221-1011221302300011-1222312112300210-1212333213022020", "registry_path": "docs/guides/resources--subnet--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Import for xcsh_subnet.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["subnetCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_subnet.example system/example
```
