---
page_title: "no_s2s_connectivity_slo"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no s2s connectivity slo"], "body_bytes": 915, "body_sha256": "sha256:b0a3b52fae5db42d3f20cd3d94c7785a7719a1cea17422c3924458bd527d26e5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:no_s2s_connectivity_slo", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/no_s2s_connectivity_slo/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1032112221231302-2123032103000023-1212321030120112-1301100133021333-3122002203103112-2022222012000331-0333232203322311-3201133120130202", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_s2s_connectivity_slo"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/no_s2s_connectivity_slo/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_s2s_connectivity_slo

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- no_s2s_connectivity_slo

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity slo.

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
no_s2s_connectivity_slo = {}
```

This is an empty object or choice marker. It has no direct properties.
