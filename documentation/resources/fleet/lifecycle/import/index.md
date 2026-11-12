---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_fleet."
xcsh_docs: {"aliases": ["fleet"], "body_bytes": 340, "body_sha256": "sha256:ab139d2e1c01db425c4bb7a213bfc3a3513291acc4cf9da72090ff982cb47ff2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:fleet:fundamentals", "path": "documentation/resources/fleet/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3221301112320030-3323110332312301-0001311302000302-1300222133101221-2120323300022212-1210133323011120-1330033031131301-3222311203330132", "registry_path": "docs/guides/resources--fleet--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_fleet.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["fleetCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_fleet.example system/example
```
