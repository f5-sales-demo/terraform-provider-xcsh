---
page_title: "site_subnet_params"
subcategory: ""
description: "site_subnet_params for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 2931, "body_sha256": "sha256:c007c34098c4d6c47195fd02383bcd49fba991c72a69d326278568d8b4e702f9", "child_ids": ["xcsh-docs:data-sources:subnet:properties:site_subnet_params:dhcp", "xcsh-docs:data-sources:subnet:properties:site_subnet_params:site", "xcsh-docs:data-sources:subnet:properties:site_subnet_params:static_ip", "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params"], "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params", "parent_id": "xcsh-docs:data-sources:subnet:reference", "path": "documentation/data-sources/subnet/properties/site_subnet_params/index.md", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["site_subnet_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/site_subnet_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_subnet_params for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/)
- site_subnet_params

<a id="section"></a>

Type: `"list"`. Computed.

Site Subnet Parameters. Configure subnet parameters per site.

Upstream description:

Configure subnet parameters per site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [dhcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/dhcp/): complete subsection reference.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/site/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/static_ip/): complete subsection reference.

- [subnet_dhcp_server_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/): complete subsection reference.

## Next pages

- [site_subnet_params.dhcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/dhcp/)
- [site_subnet_params.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/site/)
- [site_subnet_params.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/static_ip/)
- [site_subnet_params.subnet_dhcp_server_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
