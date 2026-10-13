---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_registration."
xcsh_docs: {"aliases": ["registration"], "body_bytes": 361, "body_sha256": "sha256:ac17a4149edaab17abfa64d0a9ae0f9bf822df198852dbcf50e41c8c667c8fbd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:registration:fundamentals", "path": "documentation/resources/registration/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:1b9ea59a790d92083efb68ff1abebe57ebbb3f7b716b91cdb7b0416134676174", "provider_type": "resources", "registry_anchor": "canonical-0003213230202100-1010100003100000-3213030201322202-1222011212300110-1030000212113132-2232120010122000-2210001300013102-0211313022120221", "registry_path": "docs/guides/resources--registration--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/lifecycle/import/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Import for xcsh_registration.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["registrationCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_registration.example system/example
```
