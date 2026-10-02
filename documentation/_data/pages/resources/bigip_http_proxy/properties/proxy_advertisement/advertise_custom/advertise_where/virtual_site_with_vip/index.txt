---
page_title: "proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip"
subcategory: ""
description: "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where virtual site with vip"], "body_bytes": 4698, "body_sha256": "sha256:a033b43463a12cd8f8faec32d25c285c3ca98600d9bce335d3b2f74a8c0018fb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "documentation/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2210222003023302-3322232233330320-1133232011210322-2302301032121031-3332022221001312-0223333301221201-1212020213132130-1210122103301303", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip"], "schema_version": 1, "sections": [{"aliases": ["ip"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip", "description": "Use given IP address as VIP on the site.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["network"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network", "description": "This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip", "network"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site", "type": "requires"}], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

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

Terraform syntax:

```terraform
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip"></a>

### ip property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

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

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
