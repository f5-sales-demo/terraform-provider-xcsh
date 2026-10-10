---
page_title: "where.site.enable_internet_vip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["where site enable internet vip"], "body_bytes": 1084, "body_sha256": "sha256:0f9e26c63e09d0b824e8d974bde3a88ce2ee831e966e88aedc91e61b96115ab8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:where:site:enable_internet_vip", "parent_id": "xcsh-docs:resources:discovery:properties:where:site", "path": "documentation/resources/discovery/properties/where/site/enable_internet_vip/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2110231323102212-0020122101231302-3330133200323223-2201010003300321-2322213312332232-3102021022010321-1320030113212012-2303013132303030", "registry_path": "docs/guides/resources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "site", "enable_internet_vip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/where/site/enable_internet_vip/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.site.enable_internet_vip

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/)
- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/)
- where.site.enable_internet_vip

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
