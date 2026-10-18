---
page_title: "logs_streaming_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["logs streaming disabled"], "body_bytes": 842, "body_sha256": "sha256:69eace51bf8fa7724e7ebf834e543c8d02918b72c62a8a1e78d11b6da42063af", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:logs_streaming_disabled", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/logs_streaming_disabled/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3212213123002023-0202231202121031-0011310312010030-0202200201110132-0022133203021121-1110002111002320-0011010213122003-0312100022301211", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["logs_streaming_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/logs_streaming_disabled/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["fleetCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# logs_streaming_disabled

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- logs_streaming_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
logs_streaming_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
