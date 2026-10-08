---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_network_firewall."
xcsh_docs: {"aliases": ["network firewall"], "body_bytes": 373, "body_sha256": "sha256:0319dbdab92ad8bfba9e76712861ff869149323098d9983065f93b99eaa1ba58", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:network_firewall:fundamentals", "path": "documentation/resources/network_firewall/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0332032223301202-2101103111010232-0221122011112100-2222113100031120-3212302122222311-2300100022022233-0032223123323302-3012203132031203", "registry_path": "docs/guides/resources--network_firewall--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Import for xcsh_network_firewall.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_firewallCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_network_firewall.example system/example
```
