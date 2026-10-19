---
page_title: "disable_fast_acl"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable fast acl"], "body_bytes": 953, "body_sha256": "sha256:796730eeed4a9139b01e244fa788ea0c16bff491ff9be479c0348d713d4b1ba2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:disable_fast_acl", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/disable_fast_acl/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1030000322023102-2003301200002102-0123111133031210-3301113222100120-1330210120023320-2333233223032033-0232123000221113-0213001130313013", "registry_path": "docs/guides/resources--network_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_fast_acl"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/disable_fast_acl/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["network_firewallCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_fast_acl

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- disable_fast_acl

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
omitted.

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
disable_fast_acl = {}
```

This is an empty object or choice marker. It has no direct properties.
