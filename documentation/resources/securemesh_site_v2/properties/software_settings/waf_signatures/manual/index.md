---
page_title: "software_settings.waf_signatures.manual"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["software settings waf signatures manual"], "body_bytes": 1213, "body_sha256": "sha256:84de8f1606cef2d9bb3b9bc8670ae05173dfb92344dc9eab73b829a41cb0a5f1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/waf_signatures/manual/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings", "waf_signatures", "manual"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/waf_signatures/manual/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
