---
page_title: "isolated_nw"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["isolated nw"], "body_bytes": 831, "body_sha256": "sha256:f2048c7688151db423de50ce485275301030eae03525fa5ba0d267078aa768fc", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:isolated_nw", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "documentation/resources/subnet/properties/isolated_nw/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0103130132011023-3120212213021202-2003313000030100-2223133212222130-0211123130000313-0300010310210232-2313122120312201-1221310100311300", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["isolated_nw"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/isolated_nw/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["subnetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# isolated_nw

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- isolated_nw

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for isolated nw.

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
isolated_nw = {}
```

This is an empty object or choice marker. It has no direct properties.
