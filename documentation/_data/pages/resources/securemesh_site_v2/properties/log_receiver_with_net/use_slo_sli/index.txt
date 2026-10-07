---
page_title: "log_receiver_with_net.use_slo_sli"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["log receiver with net use slo sli"], "body_bytes": 1057, "body_sha256": "sha256:0db8f41b44fce51a17a63e23766823a61e1cef809053ccdaa455e8f8fd281a89", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net", "path": "documentation/resources/securemesh_site_v2/properties/log_receiver_with_net/use_slo_sli/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0031021002303321-3112130012003112-1011131101211211-0033112221120300-3311120313031021-3210203333312121-1200221213301132-1033211302020230", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["log_receiver_with_net", "use_slo_sli"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/log_receiver_with_net/use_slo_sli/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# log_receiver_with_net.use_slo_sli

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [log_receiver_with_net](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/)
- log_receiver_with_net.use_slo_sli

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use slo sli.

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
use_slo_sli = {}
```

This is an empty object or choice marker. It has no direct properties.
