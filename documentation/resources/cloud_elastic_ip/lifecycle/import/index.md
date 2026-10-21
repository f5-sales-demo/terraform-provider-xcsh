---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": ["cloud elastic ip"], "body_bytes": 373, "body_sha256": "sha256:71861e6d4ebe91a4f5f5eba049c1c7990fb26d3ce92614e9c6f2eaafca366313", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_elastic_ip:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_elastic_ip:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:cloud_elastic_ip:fundamentals", "path": "documentation/resources/cloud_elastic_ip/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1100222313213321-3102211332332302-0121312102002123-2212121031030030-2231021010322103-0032031232020320-3310221223102311-2222322011311333", "registry_path": "docs/guides/resources--cloud_elastic_ip--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_elastic_ip/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_cloud_elastic_ip.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_cloud_elastic_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_cloud_elastic_ip.example system/example
```
