---
page_title: "ipsec.ike_parameters.rm_ip_address.ipv4"
subcategory: ""
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["ipsec ike parameters rm ip address ipv4"], "body_bytes": 2357, "body_sha256": "sha256:75c4fb5e79f5113e2bd7845a3a3b5d97aa6f54a78d0dca1c737a642c228ccacf", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "path": "documentation/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3200202022202200-2220212130011323-3231233323231031-2031123102033032-1202030023031202-0020320001100303-2201202010113012-3330113011002020", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["ipsec ike parameters rm ip address ipv4 addr"], "anchor": "schema-ipsec--ike_parameters--rm_ip_address--ipv4--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "ipv4", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.rm_ip_address.ipv4

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/)
- [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/)
- ipsec.ike_parameters.rm_ip_address.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="schema-ipsec--ike_parameters--rm_ip_address--ipv4--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

## Next pages

- [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
