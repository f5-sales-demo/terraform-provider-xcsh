---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_network_policy."
xcsh_docs: {"aliases": ["network policy"], "body_bytes": 367, "body_sha256": "sha256:a8c6ab9d9ee0fc5c8bc9bc0d813d919e5b96b739008154d0e940cbf0ecce074c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:network_policy:fundamentals", "path": "documentation/resources/network_policy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0203203102310102-3230133322230311-0211213233220222-0120312020122001-0231021211322131-3011220331203322-2332111033312010-3103313001010222", "registry_path": "docs/guides/resources--network_policy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Import for xcsh_network_policy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_network_policy.example system/example
```
