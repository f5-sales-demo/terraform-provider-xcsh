---
page_title: "response_cache"
subcategory: "DNS"
description: "Response Cache x-required."
xcsh_docs: {"aliases": ["response cache"], "body_bytes": 1481, "body_sha256": "sha256:74d4c8be51f44635adc268f3bdb0fd8d515df753f27d4db9e8f32982403bfcfc", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:disable_spec", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "path": "documentation/data-sources/dns_load_balancer/properties/response_cache/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120", "registry_path": "docs/guides/data-sources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache"], "schema_version": 1, "sections": [{"aliases": ["response cache default response cache parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "default_response_cache_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["response cache disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["response cache response cache parameters"], "anchor": "section", "description": "Configuration parameter for response cache parameters.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["response_cache", "response_cache_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/response_cache/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Response Cache x-required.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- response_cache

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache.

Additional upstream details:

Response Cache x-required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

## Direct properties

- [default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/disable_spec/): complete subsection reference.

- [response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/): complete subsection reference.
