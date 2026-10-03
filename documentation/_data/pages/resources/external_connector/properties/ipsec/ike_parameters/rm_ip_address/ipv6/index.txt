---
page_title: "ipsec.ike_parameters.rm_ip_address.ipv6"
subcategory: ""
description: "IPv6 Address specified as hexadecimal numbers separated by ':'"
xcsh_docs: {"aliases": ["ipsec ike parameters rm ip address ipv6"], "body_bytes": 3128, "body_sha256": "sha256:e8a1877a80619bf8b26a36482bc70b0df882febf9afb8985e69e1f39663045b8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "path": "documentation/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv6/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0113131123101301-2232011211100010-1303113030120332-1113202110330200-3010002111121003-2033222212100031-0300133103323112-2101203111100230", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "ipv6"], "schema_version": 1, "sections": [{"aliases": ["ipsec ike parameters rm ip address ipv6 addr"], "anchor": "schema-ipsec--ike_parameters--rm_ip_address--ipv6--addr", "description": "IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by ':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "ipv6", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv6/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPv6 Address specified as hexadecimal numbers separated by ':'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.rm_ip_address.ipv6

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/)
- [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/)
- ipsec.ike_parameters.rm_ip_address.ipv6

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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
ipv6 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ipsec--ike_parameters--rm_ip_address--ipv6--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
