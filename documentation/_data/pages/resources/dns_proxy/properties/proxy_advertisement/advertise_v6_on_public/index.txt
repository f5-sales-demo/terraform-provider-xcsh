---
page_title: "proxy_advertisement.advertise_v6_on_public"
subcategory: ""
description: "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available."
xcsh_docs: {"aliases": ["proxy advertisement advertise v6 on public"], "body_bytes": 1285, "body_sha256": "sha256:b88cd65b94e792abe54d9fb0b094d03d4480736b3792450b7778e6d562ebd6e1", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public:public_ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "parent_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement", "path": "documentation/resources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123", "registry_path": "docs/guides/resources--dns_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_v6_on_public"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise v6 on public public ip"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_advertisement--advertise_v6_on_public--public_ip--name", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_v6_on_public.public_ip:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public:public_ip", "type": "requires"}], "schema_path": ["proxy_advertisement", "advertise_v6_on_public", "public_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_v6_on_public

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/)
- proxy_advertisement.advertise_v6_on_public

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/public_ip/): complete subsection reference.
