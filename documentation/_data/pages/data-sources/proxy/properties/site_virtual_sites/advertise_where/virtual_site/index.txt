---
page_title: "site_virtual_sites.advertise_where.virtual_site"
subcategory: ""
description: "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised."
xcsh_docs: {"aliases": ["site virtual sites advertise where virtual site"], "body_bytes": 3751, "body_sha256": "sha256:a5bdbf3bdf005198b53b36f5db798d43a46c81ce9b7b298da8f74f62917790a6", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site:virtual_site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "parent_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where", "path": "documentation/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0011000332332130-3130012231211212-3202112322222113-3111230132210223-2100313233112001-3112212110232222-2031021001200031-3310310312213001", "registry_path": "docs/guides/data-sources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_virtual_sites", "advertise_where", "virtual_site"], "schema_version": 1, "sections": [{"aliases": ["site virtual sites advertise where virtual site network"], "anchor": "schema-site_virtual_sites--advertise_where--virtual_site--network", "description": "This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. VK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This", "document_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "virtual_site", "network"], "syntax": "attribute", "type": "string"}, {"aliases": ["site virtual sites advertise where virtual site virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "virtual_site", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_virtual_sites.advertise_where.virtual_site

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [site_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/)
- [site_virtual_sites.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/)
- site_virtual_sites.advertise_where.virtual_site

<a id="section"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
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

<a id="schema-site_virtual_sites--advertise_where--virtual_site--network"></a>

### network property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/virtual_site/): complete subsection reference.

## Next pages

- [site_virtual_sites.advertise_where.virtual_site.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/virtual_site/)
- [site_virtual_sites.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
