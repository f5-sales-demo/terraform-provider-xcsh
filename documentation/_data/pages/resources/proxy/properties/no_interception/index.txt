---
page_title: "no_interception"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no interception"], "body_bytes": 1275, "body_sha256": "sha256:f91b1958b220f76526ad0e547fe7fb9100bddeaf2401fc08fd042a49ec92b593", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:no_interception", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/no_interception/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3001311021210113-1203322120301211-2003112333102203-1131010301220310-1212123211110331-3201330202033133-1001002102301300-0023323331213323", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_interception"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/no_interception/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["proxyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_interception

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- no_interception

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_interception, tls\_intercept; Default: no\_interception\] Configuration parameter for
no interception.

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

- [no_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/no_interception/#section)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_interception = {}
```

This is an empty object or choice marker. It has no direct properties.
