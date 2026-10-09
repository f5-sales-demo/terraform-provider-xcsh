---
page_title: "http1_config.header_transformation.default_header_transformation"
subcategory: ""
description: "Use the platform's current default HTTP header transformation behavior."
xcsh_docs: {"aliases": ["http1 config header transformation default header transformation"], "body_bytes": 1178, "body_sha256": "sha256:4b3b2f3382a90ee9d94980caa5b7a1e6a6a8dfabda4006dc779d9097eb00c668", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "parent_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "path": "documentation/resources/cluster/properties/http1_config/header_transformation/default_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3210111201011333-3303233101013001-1220302301313223-2002332233223022-0331321212130333-2311311023333103-3300021323021322-0323322311021201", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http1_config", "header_transformation", "default_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/header_transformation/default_header_transformation/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Use the platform's current default HTTP header transformation behavior.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["clusterCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation.default_header_transformation

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/)
- [http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/)
- http1_config.header_transformation.default_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.
