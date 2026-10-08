---
page_title: "advertise_custom.advertise_where.virtual_network"
subcategory: "Load Balancing"
description: "Parameters to advertise on a given virtual network."
xcsh_docs: {"aliases": ["advertise custom advertise where virtual network"], "body_bytes": 3571, "body_sha256": "sha256:13091044acf60555600c1fd17a95fd27b32539ec04d3eb2aa89e40f49e954aab", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_v6_vip", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_vip", "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where", "path": "documentation/data-sources/http_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1332223010333312-0131102221313003-0300220233112013-0122111312302200-2033133230020123-2130313110323123-1023313033002113-0111200011131110", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["advertise custom advertise where virtual network default v6 vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_v6_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "default_v6_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise custom advertise where virtual network default vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "default_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise custom advertise where virtual network specific v6 vip"], "anchor": "schema-advertise_custom--advertise_where--virtual_network--specific_v6_vip", "description": "Exclusive with Use given IPv6 address as VIP on virtual Network.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "specific_v6_vip"], "syntax": "attribute", "type": "string"}, {"aliases": ["advertise custom advertise where virtual network specific vip"], "anchor": "schema-advertise_custom--advertise_where--virtual_network--specific_vip", "description": "Exclusive with Use given IPv4 address as VIP on virtual Network.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "specific_vip"], "syntax": "attribute", "type": "string"}, {"aliases": ["advertise custom advertise where virtual network virtual network"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Parameters to advertise on a given virtual network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.virtual_network

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/advertise_custom/advertise_where/)
- advertise_custom.advertise_where.virtual_network

<a id="section"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

## Direct properties

- [default_v6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/default_v6_vip/): complete subsection reference.

- [default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/default_vip/): complete subsection reference.

<a id="schema-advertise_custom--advertise_where--virtual_network--specific_v6_vip"></a>

### specific_v6_vip property

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-advertise_custom--advertise_where--virtual_network--specific_vip"></a>

### specific_vip property

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/virtual_network/): complete subsection reference.
