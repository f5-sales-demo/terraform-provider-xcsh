---
page_title: "origin_servers.origin_servers.public_ip"
subcategory: ""
description: "Specify origin server with public IP address."
xcsh_docs: {"aliases": ["origin servers origin servers public ip"], "body_bytes": 2140, "body_sha256": "sha256:6c0291a35b6ae8345828a0fdefbd3d7c3bc497bd79cadd07d0d0408aba49a32c", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/resources/dns_proxy/properties/origin_servers/origin_servers/public_ip/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0332203222013220-0130321101032111-2012303000311332-3112321103103132-1013310013011203-1231231132310313-1133221133303333-1132322120103311", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "public_ip"], "schema_version": 1, "sections": [{"aliases": ["origin servers origin servers public ip ip"], "anchor": "schema-origin_servers--origin_servers--public_ip--ip", "description": "Exclusive with Public IPv4 address.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_ip", "ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/public_ip/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Specify origin server with public IP address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.public_ip

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/)
- origin_servers.origin_servers.public_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--origin_servers--public_ip--ip"></a>

### ip property

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
