---
page_title: "custom_proxy.enable_re_tunnel"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["custom proxy enable re tunnel"], "body_bytes": 1041, "body_sha256": "sha256:f125a6ae4682cb37162ab13b672ba93a8d6b12b43b175d5c67ef3593679e4bad", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "path": "documentation/resources/securemesh_site_v2/properties/custom_proxy/enable_re_tunnel/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3111102123022002-1031033320333131-2212221101213001-2311200122113001-1132310110311032-0212310303232030-2022223013213220-2022110121020030", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_proxy", "enable_re_tunnel"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/custom_proxy/enable_re_tunnel/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_proxy.enable_re_tunnel

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [custom_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/)
- custom_proxy.enable_re_tunnel

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable re tunnel.

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
enable_re_tunnel = {}
```

This is an empty object or choice marker. It has no direct properties.
