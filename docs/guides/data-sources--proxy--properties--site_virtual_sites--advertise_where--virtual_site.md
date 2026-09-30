---
page_title: "site_virtual_sites.advertise_where.virtual_site"
subcategory: ""
description: "site_virtual_sites.advertise_where.virtual_site for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3250, "body_sha256": "sha256:42869c5b60e0d830d7c5f60d82b11c383536c0454b9a8da14eb5ca59d9066f95", "canonical_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "child_ids": ["xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site:virtual_site"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "parent_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where", "path": "docs/guides/data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_virtual_sites", "advertise_where", "virtual_site"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_virtual_sites.advertise_where.virtual_site for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_virtual_sites.advertise_where.virtual_site

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [site_virtual_sites](data-sources--proxy--properties--site_virtual_sites.md)
- [site_virtual_sites.advertise_where](data-sources--proxy--properties--site_virtual_sites--advertise_where.md)
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

- [virtual_site](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site--virtual_site.md): complete subsection reference.

## Next pages

- [site_virtual_sites.advertise_where.virtual_site.virtual_site](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site--virtual_site.md)
- [site_virtual_sites.advertise_where](data-sources--proxy--properties--site_virtual_sites--advertise_where.md)
- [xcsh_proxy](../data-sources/proxy.md)
