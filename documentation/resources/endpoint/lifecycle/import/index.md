---
page_title: "Import"
subcategory: "Networking"
description: "Import for xcsh_endpoint."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 349, "body_sha256": "sha256:1dd44cd5c910a5a7fc9ed30c664ce4c99f1ae8ab8b9056129419346fbdd7715e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:endpoint:fundamentals", "path": "documentation/resources/endpoint/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2011110030121120-1212131101212003-1032010213223012-0101131313103302-0231112212111322-0212322233113233-0230120303132200-2220122112220322", "registry_path": "docs/guides/resources--endpoint--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/lifecycle/import/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Import for xcsh_endpoint.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["endpointCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_endpoint.example system/example
```
