---
page_title: "software_settings.waf_signatures.manual"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["software settings waf signatures manual"], "body_bytes": 1213, "body_sha256": "sha256:84de8f1606cef2d9bb3b9bc8670ae05173dfb92344dc9eab73b829a41cb0a5f1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/waf_signatures/manual/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings", "waf_signatures", "manual"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/waf_signatures/manual/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings.waf_signatures.manual

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [software_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/)
- [software_settings.waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/)
- software_settings.waf_signatures.manual

<a id="section"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
manual = {}
```

This is an empty object or choice marker. It has no direct properties.
