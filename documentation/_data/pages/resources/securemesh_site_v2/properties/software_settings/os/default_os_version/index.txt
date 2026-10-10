---
page_title: "software_settings.os.default_os_version"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["software settings os default os version"], "body_bytes": 1201, "body_sha256": "sha256:98fb2a75ffad154cc853ffba634a04379b28a4abf63be63b230f34091342daa4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/os/default_os_version/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1030220122022302-1232023121110020-0231001010323331-1333120012030332-3301310103110120-1230101212120030-2013303130333212-2131233130011232", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings", "os", "default_os_version"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/os/default_os_version/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings.os.default_os_version

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [software_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/)
- [software_settings.os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/os/)
- software_settings.os.default_os_version

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
default_os_version = {}
```

This is an empty object or choice marker. It has no direct properties.
