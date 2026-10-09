---
page_title: "discovery_k8s.publish_info.dns_delegation"
subcategory: ""
description: "Configuration parameter for dns delegation."
xcsh_docs: {"aliases": ["discovery k8s publish info dns delegation"], "body_bytes": 3387, "body_sha256": "sha256:e33b7fe80fd41a4bc81e2da389ae4b175faff2975bb7afd15d7a02e01f68dcf3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "path": "documentation/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320", "registry_path": "docs/guides/resources--discovery--reference--group-002.md", "relationships": [{"anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info.dns_delegation:RequiredObjectAttributes:subdomain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s publish info dns delegation dns mode"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--dns_mode", "description": "Two modes are possible CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running kube-DNS.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["CORE_DNS", "KUBE_DNS"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "dns_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["discovery k8s publish info dns delegation subdomain"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "description": "The DNS subdomain for which F5XC will respond to DNS queries.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "subdomain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration parameter for dns delegation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["discoveryCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info.dns_delegation

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/)
- discovery_k8s.publish_info.dns_delegation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dns delegation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("subdomain")}
```

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
dns_delegation {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-discovery_k8s--publish_info--dns_delegation--dns_mode"></a>

### dns_mode property

Type: `"string"`. Optional.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CORE_DNS","KUBE_DNS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CORE_DNS",
    "KUBE_DNS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-discovery_k8s--publish_info--dns_delegation--subdomain"></a>

### subdomain property

Type: `"string"`. Optional.

The DNS subdomain for which F5XC will respond to DNS queries.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
