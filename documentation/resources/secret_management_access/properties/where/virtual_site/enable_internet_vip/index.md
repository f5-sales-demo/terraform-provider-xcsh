---
page_title: "where.virtual_site.enable_internet_vip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["where virtual site enable internet vip"], "body_bytes": 1191, "body_sha256": "sha256:5c338f69da44c8066607940421fafc50c40fbfe9560214964aca8c53fde95b26", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:where:virtual_site:enable_internet_vip", "parent_id": "xcsh-docs:resources:secret_management_access:properties:where:virtual_site", "path": "documentation/resources/secret_management_access/properties/where/virtual_site/enable_internet_vip/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1020013133130332-3230320032032223-0023121331121033-0330301002031212-3133331122203202-0331110223220330-1012212320031023-2221032000223332", "registry_path": "docs/guides/resources--secret_management_access--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_site", "enable_internet_vip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/where/virtual_site/enable_internet_vip/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_site.enable_internet_vip

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/where/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/where/virtual_site/)
- where.virtual_site.enable_internet_vip

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.
