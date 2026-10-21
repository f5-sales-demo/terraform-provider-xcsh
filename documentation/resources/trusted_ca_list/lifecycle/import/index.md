---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["trusted ca list"], "body_bytes": 370, "body_sha256": "sha256:c523ffed6263fd46b6da6d24b4d6098d51be5a41c362a0277450c0f7549f75dd", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "id": "xcsh-docs:resources:trusted_ca_list:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:trusted_ca_list:fundamentals", "path": "documentation/resources/trusted_ca_list/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2100031021303333-2312102102232223-1130221210320233-0033003021231131-1111033121121320-1122131302001000-3121332113222202-3222212211031113", "registry_path": "docs/guides/resources--trusted_ca_list--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_trusted_ca_list.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_trusted_ca_list.example system/example
```
