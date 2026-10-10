---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_bgp."
xcsh_docs: {"aliases": ["bgp"], "body_bytes": 334, "body_sha256": "sha256:d99fb686f41a80deaeb22e93cffeff966b6fee4f290e70a857d250710dbeb149", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:bgp:fundamentals", "path": "documentation/resources/bgp/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0100132031003210-2120031120132100-3313102103301211-1202001221130302-0312112022100313-1001222310002020-0031122332111031-1210300301223030", "registry_path": "docs/guides/resources--bgp--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Import for xcsh_bgp.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgpCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_bgp.example system/example
```
