---
page_title: "Import"
subcategory: "Load Balancing"
description: "Import for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["tcp loadbalancer"], "body_bytes": 373, "body_sha256": "sha256:1cd41d0cffa92c36743100ae2ba15bc71410bdbddfc6d79f5c6facbce466dd82", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:fundamentals", "path": "documentation/resources/tcp_loadbalancer/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0010103130330001-3321333010133010-0123010301022122-0013201010103131-2103100023233113-3300122101333230-2102201103021130-2001121313201221", "registry_path": "docs/guides/resources--tcp_loadbalancer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Import for xcsh_tcp_loadbalancer.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_tcp_loadbalancer.example system/example
```
