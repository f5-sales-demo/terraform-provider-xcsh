---
page_title: "auto_http_config"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["auto http config"], "body_bytes": 1613, "body_sha256": "sha256:ca7aa9f9d1f80210889d12b9bd1c672a7c404291dc2aa370fff462137fb4d44e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:auto_http_config", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/auto_http_config/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2322231020220022-3201131011222101-3111130311210132-1100300313031231-3130321322022012-1103100132333333-0022230220220021-1222033112300230", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["auto_http_config"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/auto_http_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
