---
page_title: "site_virtual_sites.advertise_where.use_default_port"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["site virtual sites advertise where use default port"], "body_bytes": 1177, "body_sha256": "sha256:2e8125601942747be0c953cc094b33c182c526e05e9c92fea985570783fecccd", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "parent_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "path": "documentation/resources/proxy/properties/site_virtual_sites/advertise_where/use_default_port/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2012100001011113-3312200301101231-3022131230331303-2320221302012122-0221222000320123-2101221112033311-1111331100113220-1312222100321310", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_virtual_sites", "advertise_where", "use_default_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/site_virtual_sites/advertise_where/use_default_port/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_virtual_sites.advertise_where.use_default_port

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [site_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/)
- [site_virtual_sites.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/advertise_where/)
- site_virtual_sites.advertise_where.use_default_port

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.
