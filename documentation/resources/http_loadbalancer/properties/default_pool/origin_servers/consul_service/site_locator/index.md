---
page_title: "default_pool.origin_servers.consul_service.site_locator"
subcategory: "Load Balancing"
description: "This message defines a reference to a site or virtual site object."
xcsh_docs: {"aliases": ["default pool origin servers consul service site locator"], "body_bytes": 1848, "body_sha256": "sha256:4fd4b4bbf545e6321a2c7cb15ed9648c3a7ac7e4eae6e6f57a0900412c488006", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:site_locator:site", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:site_locator:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:site_locator", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/site_locator/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1323111100223303-1312011332301001-3103112212130120-1123302201120002-2201130210332220-3203331113133221-3111311301322111-1101213123011332", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "consul_service", "site_locator"], "schema_version": 1, "sections": [{"aliases": ["default pool origin servers consul service site locator site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:site_locator:site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "consul_service", "site_locator", "site"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers consul service site locator virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:site_locator:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "consul_service", "site_locator", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/site_locator/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This message defines a reference to a site or virtual site object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.consul_service.site_locator

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/)
- default_pool.origin_servers.consul_service.site_locator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/site_locator/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/site_locator/virtual_site/): complete subsection reference.
