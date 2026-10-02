---
page_title: "ingress_egress_gw.outside_network.new_network"
subcategory: "Infrastructure"
description: "Parameters to create a new GCP VPC Network."
xcsh_docs: {"aliases": ["ingress egress gw outside network new network"], "body_bytes": 2950, "body_sha256": "sha256:becab8e2d73bdc6804de103f9b6a58aba60c96db33328dd42c074821259ce694", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_network", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/new_network/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3003031120212300-3203113033003113-2131102210302233-0311203030003333-2033222222112012-3103031101213323-3132213300022302-1132302013200212", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "outside_network", "new_network"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-ingress_egress_gw--outside_network--new_network--name", "description": "Name for your GCP VPC Network.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "outside_network", "new_network", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/new_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters to create a new GCP VPC Network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_network.new_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/)
- ingress_egress_gw.outside_network.new_network

<a id="section"></a>

Type: `"single"`. Computed.

Parameters to create a new GCP VPC Network.

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

<a id="schema-ingress_egress_gw--outside_network--new_network--name"></a>

### name property

Type: `"string"`. Computed.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [ingress_egress_gw.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
