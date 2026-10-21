---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_app_firewall."
xcsh_docs: {"aliases": ["app firewall"], "body_bytes": 361, "body_sha256": "sha256:a59dee5c81a7368e7a272f12f89def382a466618266413a60e827871173b61d1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:app_firewall:fundamentals", "path": "documentation/resources/app_firewall/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3033201022113101-3211011303131013-0032132120121103-0122023302120031-0320321330332110-1122232103021330-2201123000113020-1323003211223223", "registry_path": "docs/guides/resources--app_firewall--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_app_firewall.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_app_firewall.example system/example
```
