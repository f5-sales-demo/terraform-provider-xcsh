---
page_title: "primary.rr_set_group.rr_set.tlsa_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary rr set group rr set tlsa record values"], "body_bytes": 6341, "body_sha256": "sha256:33383eb578e0f6f2980431d40c4ca34b32f50741019df8c5746a35e9b90b7eb2", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record:values", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record", "path": "documentation/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2300322000010100-3330131011030130-0002231310231000-0231233002101022-1131210322121011-2330020130222313-3123203000313102-2213110333130232", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "tlsa_record", "values"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate association data", "existing certificates", "tls certificates"], "anchor": "schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_association_data", "description": "The actual data to be matched given the settings of the other fields.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "tlsa_record", "values", "certificate_association_data"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "certificate usage", "existing certificates", "tls certificates"], "anchor": "schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_usage", "description": "- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint: Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion - DomainIssuedCertificate: Domain Issued Certificate.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "tlsa_record", "values", "certificate_usage"], "syntax": "attribute", "type": "string"}, {"aliases": ["matching type"], "anchor": "schema-primary--rr_set_group--rr_set--tlsa_record--values--matching_type", "description": "- NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "tlsa_record", "values", "matching_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "existing certificates", "selector", "tls certificates"], "anchor": "schema-primary--rr_set_group--rr_set--tlsa_record--values--selector", "description": "- FullCertificate: Full Certificate - UseSubjectPublicKey: Use Subject Public Key.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "tlsa_record", "values", "selector"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.tlsa_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary.rr_set_group.rr_set.tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="section"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_association_data"></a>

### certificate_association_data property

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_usage"></a>

### certificate_usage property

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--tlsa_record--values--matching_type"></a>

### matching_type property

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--tlsa_record--values--selector"></a>

### selector property

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [primary.rr_set_group.rr_set.tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
