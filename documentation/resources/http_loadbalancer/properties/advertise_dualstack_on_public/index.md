---
page_title: "advertise_dualstack_on_public"
subcategory: "Load Balancing"
description: "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available."
xcsh_docs: {"aliases": ["advertise dualstack on public"], "body_bytes": 1152, "body_sha256": "sha256:20e5002a6fee3944fb73ab3e2840d66a9a8882c8e6aadea06622e7da8352b421", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:advertise_dualstack_on_public:public_ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_dualstack_on_public", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/advertise_dualstack_on_public/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0223122302110001-1233213021312213-1303312231122322-3210130310233332-1323202101300301-0012222323223300-1111202232330131-0221301031130130", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_dualstack_on_public"], "schema_version": 1, "sections": [{"aliases": ["advertise dualstack on public public ip"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_dualstack_on_public:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-advertise_dualstack_on_public--public_ip--name", "enforcement": "provider-schema", "group": "advertise_dualstack_on_public.public_ip:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_dualstack_on_public:public_ip", "type": "requires"}], "schema_path": ["advertise_dualstack_on_public", "public_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_dualstack_on_public/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_dualstack_on_public

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- advertise_dualstack_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_dualstack_on_public/public_ip/): complete subsection reference.
