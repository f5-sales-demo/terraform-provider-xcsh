---
page_title: "active_service_policies"
subcategory: ""
description: "List of service policies."
xcsh_docs: {"aliases": ["active service policies"], "body_bytes": 2072, "body_sha256": "sha256:e596927e512b3bf9c0bfcc804dd3cabd9bb92914f90740a86cb83998af18bffd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies:policies"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:reference", "path": "documentation/data-sources/udp_loadbalancer/properties/active_service_policies/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3012132102203030-3133031323022332-1202102030001302-2101121232121222-1230212330110021-3133230213123232-3113221120113321-3020002310301231", "registry_path": "docs/guides/data-sources--udp_loadbalancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "sections": [{"aliases": ["policies"], "anchor": "section", "description": "Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to GET the intended result. For each request, its characteristics are evaluated", "document_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["active_service_policies", "policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of service policies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_service_policies

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/)
- active_service_policies

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Upstream description:

List of service policies.

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

- [active_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/active_service_policies/#section)
- [no_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/no_service_policies/#section)
- [service_policies_from_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/service_policies_from_namespace/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/active_service_policies/policies/): complete subsection reference.

## Next pages

- [active_service_policies.policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/active_service_policies/policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
