---
page_title: "dns_ntp_config.custom_dns"
subcategory: ""
description: "DNS Servers."
xcsh_docs: {"aliases": ["dns ntp config custom dns"], "body_bytes": 2307, "body_sha256": "sha256:e1316afcb63ce2c9b502ec03b17851956fba1bd0f80d97d9875344d6d71f4490", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config", "path": "documentation/resources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3012001213011312-2021230120013323-2020301303130133-2103101312332311-3022013310032001-1213000122300303-1311133100123222-3132220000230031", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_ntp_config", "custom_dns"], "schema_version": 1, "sections": [{"aliases": ["dns servers"], "anchor": "schema-dns_ntp_config--custom_dns--dns_servers", "description": "DNS Servers.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "custom_dns", "dns_servers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "DNS Servers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config.custom_dns

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [dns_ntp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/)
- dns_ntp_config.custom_dns

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNS Servers. DNS Servers.

Upstream description:

DNS Servers.

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
custom_dns {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dns_ntp_config--custom_dns--dns_servers"></a>

### dns_servers property

Type: `["list", "string"]`. Optional.

DNS Servers. DNS Servers.

Upstream description:

DNS Servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [dns_ntp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
