---
page_title: "http1_config.header_transformation.preserve_case_header_transformation"
subcategory: ""
description: "Preserve HTTP header-name case when upstream case must remain unchanged."
xcsh_docs: {"aliases": ["http1 config header transformation preserve case header transformation"], "body_bytes": 1489, "body_sha256": "sha256:3522d3977d3cfe70035af494baac6d29ce9d467a83b25adfba2512f59e06b526", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "parent_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "path": "documentation/resources/cluster/properties/http1_config/header_transformation/preserve_case_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2233030001120322-0200313113322200-1313032121203330-2322320303213222-1213122030323010-1013033223110220-0331000121023311-1202330301032131", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http1_config", "header_transformation", "preserve_case_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/header_transformation/preserve_case_header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Preserve HTTP header-name case when upstream case must remain unchanged.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation.preserve_case_header_transformation

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/)
- [http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/)
- http1_config.header_transformation.preserve_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
