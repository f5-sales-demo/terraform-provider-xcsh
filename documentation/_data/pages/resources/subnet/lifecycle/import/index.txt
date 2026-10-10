---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_subnet."
xcsh_docs: {"aliases": ["subnet"], "body_bytes": 343, "body_sha256": "sha256:6eca1c99357ff547c89e0fdcf64a1157559972448be63cd7e364a187db7e144e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:subnet:fundamentals", "path": "documentation/resources/subnet/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2133302013023312-0313221302123301-1313011330200313-3112332131333123-1023233100311221-1011221302300011-1222312112300210-1212333213022020", "registry_path": "docs/guides/resources--subnet--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Import for xcsh_subnet.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["subnetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_subnet.example system/example
```
