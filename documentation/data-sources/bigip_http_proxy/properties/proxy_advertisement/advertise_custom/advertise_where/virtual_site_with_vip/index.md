---
page_title: "proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip"
subcategory: ""
description: "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where virtual site with vip"], "body_bytes": 3439, "body_sha256": "sha256:20e0405c6c9dad7c5bbb42cc41f5595e95203f04397ec4d2a48a462a75939ea4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0301313112323223-1130020222220223-2203231102032211-2233233211022332-1020011330022232-0201311310121232-0123123020333210-0022212123000333", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom advertise where virtual site with vip ip"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip", "description": "Use given IP address as VIP on the site.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy advertisement advertise custom advertise where virtual site with vip network"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network", "description": "This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip", "network"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy advertisement advertise custom advertise where virtual site with vip virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="section"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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

## Direct properties

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip"></a>

### ip property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network"></a>

### network property

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/): complete subsection reference.
