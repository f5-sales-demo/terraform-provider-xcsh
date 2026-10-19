---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_tenant_configuration."
xcsh_docs: {"aliases": ["tenant configuration"], "body_bytes": 385, "body_sha256": "sha256:fb86bef73aeabdc656e4acf4e753597dc273f281d204d525786c1cc8fe19de70", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:tenant_configuration:fundamentals", "path": "documentation/resources/tenant_configuration/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2111222301210110-0322113322233301-0220023023110102-0313333133103330-2123200121120022-0013200202301033-2001310232311332-2322212221312132", "registry_path": "docs/guides/resources--tenant_configuration--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Import for xcsh_tenant_configuration.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_tenant_configuration.example system/example
```
