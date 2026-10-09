---
page_title: "global"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["global"], "body_bytes": 824, "body_sha256": "sha256:3f41486c90e4a46e433cf25b93f564d8260feaded79058f33288e815efd721cc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:geo_location_set:properties:global", "parent_id": "xcsh-docs:resources:geo_location_set:reference", "path": "documentation/resources/geo_location_set/properties/global/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3321203011103202-0103022333122011-1322313223010213-1110200030333021-1103232001011002-0001332310101033-3132221322222331-2122032032321002", "registry_path": "docs/guides/resources--geo_location_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["global"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/geo_location_set/properties/global/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# global

Breadcrumbs:

- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/)
- global

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
global = {}
```

This is an empty object or choice marker. It has no direct properties.
