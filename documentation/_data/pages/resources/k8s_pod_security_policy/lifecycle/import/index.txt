---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": ["k8s pod security policy"], "body_bytes": 394, "body_sha256": "sha256:3634a1362bd5925e3b6601a4e7135bb97f1344b80ec6b7e3b52547d5b36a6f93", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:fundamentals", "path": "documentation/resources/k8s_pod_security_policy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1303131300310130-3302312132233323-0211121212103303-3103110122303033-0031320121012303-2021310330310011-2303123102231302-3030213013323020", "registry_path": "docs/guides/resources--k8s_pod_security_policy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Import for xcsh_k8s_pod_security_policy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_k8s_pod_security_policy.example system/example
```
