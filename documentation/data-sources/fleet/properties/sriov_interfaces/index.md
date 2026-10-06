---
page_title: "sriov_interfaces"
subcategory: ""
description: "List of all custom SR-IOV interfaces configuration."
xcsh_docs: {"aliases": ["sriov interfaces"], "body_bytes": 830, "body_sha256": "sha256:b0a5db11eaa2f8a3ec2b51160cc020210f2798e7e3551eddf8f36f1239ad01a2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:sriov_interfaces:sriov_interface"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/sriov_interfaces/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1132110120003111-0323103313100123-2113310212200301-0133232320111330-0030220003213233-1113311012021102-1212011223322232-3112201313103010", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sriov_interfaces"], "schema_version": 1, "sections": [{"aliases": ["sriov interfaces sriov interface"], "anchor": "section", "description": "Use custom SR-IOV interfaces Configuration.", "document_id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces:sriov_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/sriov_interfaces/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of all custom SR-IOV interfaces configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- sriov_interfaces

<a id="section"></a>

Type: `"single"`. Computed.

List of all custom SR-IOV interfaces configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/sriov_interfaces/sriov_interface/): complete subsection reference.
