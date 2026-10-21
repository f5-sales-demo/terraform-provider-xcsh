---
page_title: "waf_type.inherit_waf"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["waf type inherit waf"], "body_bytes": 981, "body_sha256": "sha256:78fb3df8718b8b0fe1cd97cf32fd679407d3870836cafd380b33e5249979860d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:waf_type:inherit_waf", "parent_id": "xcsh-docs:resources:virtual_host:properties:waf_type", "path": "documentation/resources/virtual_host/properties/waf_type/inherit_waf/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1310320202133212-2221033303020302-0011200030000030-2010112132322220-1232023133223203-1113301212012300-1132200001133223-1303021333223200", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_type", "inherit_waf"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/waf_type/inherit_waf/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.inherit_waf

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/)
- waf_type.inherit_waf

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit waf.

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
inherit_waf = {}
```

This is an empty object or choice marker. It has no direct properties.
