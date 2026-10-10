---
page_title: "caching_policy"
subcategory: "Load Balancing"
description: "Caching Policies for the CDN."
xcsh_docs: {"aliases": ["caching policy"], "body_bytes": 1687, "body_sha256": "sha256:676c66a66da3c3e0873b820943cae6676872117ec2dacfae1d87c0a2c60a99f0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/caching_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy"], "schema_version": 1, "sections": [{"aliases": ["caching policy custom cache rule"], "anchor": "section", "description": "Caching policies for CDN.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["caching_policy", "custom_cache_rule"], "syntax": "block", "type": "object"}, {"aliases": ["caching policy default cache action"], "anchor": "section", "description": "This defines a Default Cache Action.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["caching_policy", "default_cache_action"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/caching_policy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Caching Policies for the CDN.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- caching_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Additional upstream details:

Caching Policies for the CDN.

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

- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/#section)
- [disable_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/disable_caching/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
caching_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/): complete subsection reference.

- [default_cache_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/default_cache_action/): complete subsection reference.
