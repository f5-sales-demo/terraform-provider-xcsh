---
page_title: "site_subnet_params"
subcategory: ""
description: "Configure subnet parameters per site."
xcsh_docs: {"aliases": ["site subnet params"], "body_bytes": 2931, "body_sha256": "sha256:c007c34098c4d6c47195fd02383bcd49fba991c72a69d326278568d8b4e702f9", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:subnet:properties:site_subnet_params:dhcp", "xcsh-docs:data-sources:subnet:properties:site_subnet_params:site", "xcsh-docs:data-sources:subnet:properties:site_subnet_params:static_ip", "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params", "parent_id": "xcsh-docs:data-sources:subnet:reference", "path": "documentation/data-sources/subnet/properties/site_subnet_params/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003", "registry_path": "docs/guides/data-sources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params"], "schema_version": 1, "sections": [{"aliases": ["dhcp"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:dhcp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_subnet_params", "dhcp"], "syntax": "attribute", "type": "object"}, {"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_subnet_params", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["static ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_subnet_params", "static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnet dhcp server params"], "anchor": "section", "description": "Subnet DHCP parameters will be a subset of network_interface.dhcpserverparameterstype as all features in network_interface.dhcpserverparameterstype may not be supported in a subnet.", "document_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_subnet_params", "subnet_dhcp_server_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/site_subnet_params/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure subnet parameters per site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
