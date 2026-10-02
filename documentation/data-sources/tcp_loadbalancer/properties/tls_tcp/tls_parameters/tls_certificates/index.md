---
page_title: "tls_tcp.tls_parameters.tls_certificates"
subcategory: "Load Balancing"
description: "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls parameters tls certificates"], "body_bytes": 5599, "body_sha256": "sha256:eed3238c3ce6ed2fdec2a2fa68538ac7dabb3e66ed7667c45a758697e0d0df3b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-tls_tcp--tls_parameters--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "custom_hash_algorithms"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "description spec", "existing certificates", "tls certificates"], "anchor": "schema-tls_tcp--tls_parameters--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters.tls_certificates

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- tls_tcp.tls_parameters.tls_certificates

<a id="section"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-tls_tcp--tls_parameters--tls_certificates--certificate_url"></a>

### certificate_url property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-tls_tcp--tls_parameters--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/custom_hash_algorithms/)
- [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/disable_ocsp_stapling/)
- [tls_tcp.tls_parameters.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/)
- [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/use_system_defaults/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
