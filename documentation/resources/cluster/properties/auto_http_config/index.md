---
page_title: "auto_http_config"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["auto http config"], "body_bytes": 1379, "body_sha256": "sha256:081c5c4fcf0aee52ff83882316804a8a2d5e58621870c66cb5e2291cf5780712", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:auto_http_config", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/auto_http_config/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2322231020220022-3201131011222101-3111130311210132-1100300313031231-3130321322022012-1103100132333333-0022230220220021-1222033112300230", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["auto_http_config"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/auto_http_config/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["clusterCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# auto_http_config

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- auto_http_config

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: auto\_http\_config, http1\_config, http2\_options\] Enable this option

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

OneOf alternatives in this subsection:

- [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/auto_http_config/#section)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/#section)
- [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http2_options/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
auto_http_config = {}
```

This is an empty object or choice marker. It has no direct properties.
